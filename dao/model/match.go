package model

import (
	"fmt"
	"time"
)

type Match struct {
	Id            int
	LegacyId      string
	CompetitionId int
	AwayTeamId    int
	HomeTeamId    int
	IsLocked      bool
	Start         time.Time
	Stage         int
	Week          int
}

func (this Match) String() string {
	return fmt.Sprintf("{ id: %d, legacyId: %s, competitionId: %d, awayTeamId: %d, homeTeamId: %d, isLocked: %t, start: %s, stage: %d, week: %d }",
		this.Id,
		this.LegacyId,
		this.CompetitionId,
		this.AwayTeamId,
		this.HomeTeamId,
		this.IsLocked,
		this.Start.String(),
		this.Stage,
		this.Week)
}
