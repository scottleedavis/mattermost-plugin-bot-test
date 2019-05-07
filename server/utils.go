package main

import (
	"time"

	"github.com/mattermost/mattermost-server/model"
	"github.com/tkuchiki/go-timezone"
)

func (p *Plugin) location(user *model.User) *time.Location {
	tz := user.GetPreferredTimezone()
	if tz == "" {
		tzCode, _ := time.Now().Zone()

		if tzLoc, err := timezone.GetTimezones(tzCode); err != nil {
			return time.Now().Location()
		} else {
			if len(tzLoc) > 0 {
				if l, lErr := time.LoadLocation(tzLoc[0]); lErr != nil {
					return time.Now().Location()
				} else {
					return l
				}
			} else {
				return time.Now().Location()
			}
		}
	} else {
		location, _ := time.LoadLocation(tz)
		return location
	}

}
