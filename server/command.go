package main

import (
	"fmt"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
	"github.com/pkg/errors"
)

const CommandTrigger = "test"

func (p *Plugin) registerCommand(teamId string) error {
	if err := p.API.RegisterCommand(&model.Command{
		TeamId:           teamId,
		Trigger:          CommandTrigger,
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
	return p.API.UnregisterCommand(teamId, CommandTrigger)
}

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {

	post := model.Post{
		ChannelId: args.ChannelId,
		UserId:    p.botId,
		Message:   "bot post",
	}

	if _, pErr := p.API.CreatePost(&post); pErr != nil {
		p.API.LogError(fmt.Sprintf("%v", pErr))
	}

	return &model.CommandResponse{}, nil

}
