package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
	"github.com/pkg/errors"
)

func (p *Plugin) registerCommand(teamId string) error {
	if err := p.API.RegisterCommand(&model.Command{
		TeamId:           teamId,
		Trigger:          "test",
		Username:         botName,
		AutoComplete:     true,
		AutoCompleteHint: "",
		AutoCompleteDesc: "",
		DisplayName:      "",
		Description:      "",
	}); err != nil {
		return errors.Wrap(err, "failed to register command")
	}

	if err := p.API.RegisterCommand(&model.Command{
		TeamId:           teamId,
		Trigger:          "test-button",
		Username:         botName,
		AutoComplete:     true,
		AutoCompleteHint: "",
		AutoCompleteDesc: "",
		DisplayName:      "",
		Description:      "",
	}); err != nil {
		return errors.Wrap(err, "failed to register command")
	}

	if err := p.API.RegisterCommand(&model.Command{
		TeamId:           teamId,
		Trigger:          "test-button-ephemeral",
		Username:         botName,
		AutoComplete:     true,
		AutoCompleteHint: "",
		AutoCompleteDesc: "",
		DisplayName:      "",
		Description:      "",
	}); err != nil {
		return errors.Wrap(err, "failed to register command")
	}

	if err := p.API.RegisterCommand(&model.Command{
		TeamId:           teamId,
		Trigger:          "test-time",
		Username:         botName,
		AutoComplete:     true,
		AutoCompleteHint: "",
		AutoCompleteDesc: "",
		DisplayName:      "",
		Description:      "",
	}); err != nil {
		return errors.Wrap(err, "failed to register command")
	}

	if err := p.API.RegisterCommand(&model.Command{
		TeamId:           teamId,
		Trigger:          "test-bot-create-bot",
		Username:         botName,
		AutoComplete:     true,
		AutoCompleteHint: "",
		AutoCompleteDesc: "",
		DisplayName:      "",
		Description:      "",
	}); err != nil {
		return errors.Wrap(err, "failed to register command")
	}

	if err := p.API.RegisterCommand(&model.Command{
		TeamId:           teamId,
		Trigger:          "test-ephemeral-post-override",
		Username:         botName,
		AutoComplete:     true,
		AutoCompleteHint: "",
		AutoCompleteDesc: "",
		DisplayName:      "",
		Description:      "",
	}); err != nil {
		return errors.Wrap(err, "failed to register command")
	}

	if err := p.API.RegisterCommand(&model.Command{
		TeamId:           teamId,
		Trigger:          "test-public-remind-api",
		Username:         botName,
		AutoComplete:     true,
		AutoCompleteHint: "",
		AutoCompleteDesc: "",
		DisplayName:      "",
		Description:      "",
	}); err != nil {
		return errors.Wrap(err, "failed to register command")
	}

	return nil
}

func (p *Plugin) unregisterCommand(teamId string) error {
	p.API.UnregisterCommand(teamId, "test")
	p.API.UnregisterCommand(teamId, "test-button")
	p.API.UnregisterCommand(teamId, "test-button-ephemeral")
	p.API.UnregisterCommand(teamId, "test-time")
	p.API.UnregisterCommand(teamId, "test-bot-create-bot")
	p.API.UnregisterCommand(teamId, "test-ephemeral-post-override")
	return p.API.UnregisterCommand(teamId, "test-public-remind-api")
}

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {

	p.API.LogInfo("ExecuteCommand " + args.Command)

	command := strings.Trim(args.Command, " ")
	if command == "/test" {
		p.API.LogInfo("in test command")
		post := model.Post{
			ChannelId: args.ChannelId,
			UserId:    p.botId,
			Message:   "bot post",
		}

		if _, pErr := p.API.CreatePost(&post); pErr != nil {
			p.API.LogError(fmt.Sprintf("%v", pErr))
		}
	} else if command == "/test-button" {

		URL := fmt.Sprintf("%s", *p.ServerConfig.ServiceSettings.SiteURL)

		post := model.Post{
			ChannelId: args.ChannelId,
			UserId:    p.botId,
			Message:   "test button 2",
			Props: model.StringInterface{
				"attachments": []*model.SlackAttachment{
					{
						Actions: []*model.PostAction{
							{
								Id: model.NewId(),
								Integration: &model.PostActionIntegration{
									Context: model.StringInterface{},
									URL:     fmt.Sprintf("%s/plugins/%s/button2", URL, manifest.Id),
								},
								Type: model.POST_ACTION_TYPE_BUTTON,
								Name: "click",
							},
						},
					},
				},
			},
		}
		if _, pErr := p.API.CreatePost(&post); pErr != nil {
			p.API.LogError(fmt.Sprintf("%v", pErr))
		}
	} else if command == "/test-button-ephemeral" {

		URL := fmt.Sprintf("%s", *p.ServerConfig.ServiceSettings.SiteURL)

		post := model.Post{
			ChannelId: args.ChannelId,
			UserId:    p.botId,
			Message:   "test button 2",
			Props: model.StringInterface{
				"attachments": []*model.SlackAttachment{
					{
						Actions: []*model.PostAction{
							{
								Id: model.NewId(),
								Integration: &model.PostActionIntegration{
									Context: model.StringInterface{},
									URL:     fmt.Sprintf("%s/plugins/%s/button-ephemeral", URL, manifest.Id),
								},
								Type: model.POST_ACTION_TYPE_BUTTON,
								Name: "click",
							},
						},
					},
				},
			},
		}
		p.API.SendEphemeralPost(args.UserId, &post)

	} else if command == "/test-time" {

		user, _ := p.API.GetUser(args.UserId)
		location := p.location(user)

		post := model.Post{
			ChannelId: args.ChannelId,
			UserId:    p.botId,
			Message:   time.Now().In(location).Format(time.RFC822),
			Props:     model.StringInterface{},
		}
		if _, pErr := p.API.CreatePost(&post); pErr != nil {
			p.API.LogError(fmt.Sprintf("%v", pErr))
		}
	} else if command == "/test-bot-create-bot" {

		botName := model.NewId()
		if bot, err := p.API.CreateBot(&model.Bot{
			Username:    botName,
			DisplayName: botName,
			OwnerId:     p.botId,
			Description: "Bot created by TestBot",
		}); err != nil {
			post := model.Post{
				ChannelId: args.ChannelId,
				UserId:    p.botId,
				Message:   "Failed to create bot " + err.Message,
				Props:     model.StringInterface{},
			}
			if _, pErr := p.API.CreatePost(&post); pErr != nil {
				p.API.LogError(fmt.Sprintf("%v", pErr))
			}
		} else {
			post := model.Post{
				ChannelId: args.ChannelId,
				UserId:    p.botId,
				Message:   "Created bot " + bot.Username,
				Props:     model.StringInterface{},
			}
			if _, pErr := p.API.CreatePost(&post); pErr != nil {
				p.API.LogError(fmt.Sprintf("%v", pErr))
			}
		}

	} else if command == "/test-ephemeral-post-override" {

		p.API.SendEphemeralPost(args.UserId, &model.Post{
			UserId:    p.botId,
			ChannelId: args.ChannelId,
			Message:   "Bot ephemeral link",
			Props: model.StringInterface{
				"type": "system_ephemeral_test_plugin",
			},
		})

	} /*else if command == "/test-public-remind-api" {

		URL := fmt.Sprintf("%s", *p.ServerConfig.ServiceSettings.SiteURL)
		remindPluginId := "com.github.scottleedavis.mattermost-plugin-remind"
		target := "@skawtus"
		message := "a message"
		when := "in 2 seconds"
		//http://localhost:8065/plugins/com.github.scottleedavis.mattermost-plugin-remind/public/reminder
		var jsonStr = []byte(`{"target":"@skawtus"}`)
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/plugins/%s/public/reminder", URL, remindPluginId), bytes.NewBuffer(jsonStr))
		//req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			p.API.LogError(err.Error())
		} else {
			defer resp.Body.Close()
			p.API.LogInfo("response Status:" + resp.Status)
			//p.API.LogInfo("response Headers:" + resp.Header)
			body, _ := ioutil.ReadAll(resp.Body)
			p.API.LogInfo("response Body:" + string(body))

		}

		post := model.Post{
			ChannelId: args.ChannelId,
			UserId:    p.botId,
			Message:   target + " " + message + " " + when,
			Props: model.StringInterface{
				"attachments": []*model.SlackAttachment{
					{
						Actions: []*model.PostAction{
							{
								Id: model.NewId(),
								Integration: &model.PostActionIntegration{
									Context: model.StringInterface{
										"target":  target,
										"message": message,
										"when":    when,
									},
									URL: fmt.Sprintf("%s/plugins/%s/public/reminder", URL, remindPluginId),
								},
								Type: model.POST_ACTION_TYPE_BUTTON,
								Name: "click",
							},
						},
					},
				},
			},
		}
		if _, pErr := p.API.CreatePost(&post); pErr != nil {
			p.API.LogError(fmt.Sprintf("%v", pErr))
		}

	}*/

	return &model.CommandResponse{}, nil

}
