package model

import (
	"fmt"
)

type LeagueUser struct {
	Id       int
	LegacyId string
	UserId   int
	LeagueId int
	IsActive bool
}

func (this LeagueUser) String() string {
	return fmt.Sprintf("{ id: %d, legacyId: %s, userId: %d, leagueId: %d, isActive: %t }",
		this.Id,
		this.LegacyId,
		this.UserId,
		this.LeagueId,
		this.IsActive)
}
