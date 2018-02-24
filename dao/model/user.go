package model

import (
	"fmt"
	"time"
)

type User struct {
	Id          int
	LegacyId    string
	BattleNetId string
	DiscordId   string
	DisplayName string
	Email       string
	GrAvatar    string
	TeamLogo    string
	LastActive  time.Time
}

func (this User) String() string {
	return fmt.Sprintf("{ id: %d, legacyId: %s, battleNetId: %s, discordId: %s, displayName: %s, email: %s, gravatar: %s, teamLogo: %s, lastActive: %d }",
		this.Id,
		this.LegacyId,
		this.BattleNetId,
		this.DiscordId,
		this.DisplayName,
		this.Email,
		this.GrAvatar,
		this.TeamLogo,
		this.LastActive.Unix())
}
