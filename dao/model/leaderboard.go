package model

import (
	"fmt"
	"time"
)

type LeaderBoard struct {
	Id       int
	LegacyId string
	Type     string
	LeagueId int
	Stage    int
	IsActive bool
}

func (this LeaderBoard) String() string {
	return fmt.Sprintf("{ id: %d, legacyId: %s, type: %s, leagueId: %d, stage: %d, isActive: %t }",
		this.Id,
		this.LegacyId,
		this.Type,
		this.LeagueId,
		this.Stage,
		this.IsActive)
}
