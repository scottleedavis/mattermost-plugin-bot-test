package main

import (
	"github.com/mattermost/mattermost-server/plugin"
	"io/ioutil"
)

type Plugin struct {
	plugin.MattermostPlugin

	botId string

	readFile func(path string) ([]byte, error)
}

func NewPlugin() *Plugin {
	return &Plugin{
		readFile: ioutil.ReadFile,
	}
}
