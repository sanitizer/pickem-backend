package model

import (
	"fmt"
)

type LeaderBoard struct {
	Id       int
	Type     string
	LeagueId int
	Stage    int
	IsActive bool
}

func (this LeaderBoard) String() string {
	return fmt.Sprintf("{ id: %d, type: %s, leagueId: %d, stage: %d, isActive: %t }",
		this.Id,
		this.Type,
		this.LeagueId,
		this.Stage,
		this.IsActive)
}
