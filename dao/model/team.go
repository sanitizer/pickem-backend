package model

import (
	"fmt"
)

type Team struct {
	Id            int
	LegacyId      string
	Name          string
	ShortName     string
	CompetitionId int
	Logo          string
}

func (this Team) String() string {
	return fmt.Sprintf("{ id: %d, legacyId: %s, name: %s, shortName: %s, competitionId: %d, logo: %s }",
		this.Id,
		this.LegacyId,
		this.Name,
		this.ShortName,
		this.CompetitionId,
		this.Logo)
}
