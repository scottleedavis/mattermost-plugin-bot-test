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

	return nil
}

func (p *Plugin) unregisterCommand(teamId string) error {
	p.API.UnregisterCommand(teamId, "test")
	p.API.UnregisterCommand(teamId, "test-button")
	p.API.UnregisterCommand(teamId, "test-time")
	p.API.UnregisterCommand(teamId, "test-bot-create-bot")
	return p.API.UnregisterCommand(teamId, "test-ephemeral-post-override")
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

		if _, err := p.API.CreatePost(&model.Post{
			UserId:    p.botId,
			ChannelId: args.ChannelId,
			Message:   "Bot ephemeral link",
			Type:      "custom_test_plugin",
		}); err != nil {
			return &model.CommandResponse{}, err
		}

	}

	return &model.CommandResponse{}, nil

}
