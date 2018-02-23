package model

import (
	"fmt"
)

type League struct {
	Id            int
	LegacyId      string
	Name          string
	OwnerId       int
	CompetitionId int
	SimpleMode    bool
	IsLocked      bool
	IsPublic      bool
	MaxUsers      int
}

func (this League) String() string {
	return fmt.Sprintf("{ id: %d, legacyId: %s, name: %s, ownerId: %d, competitionId: %d, simpleMode: %t, isLocked: %t, isPublic: %t, maxUsers: %d }",
		this.Id,
		this.LegacyId,
		this.Name,
		this.OwnerId,
		this.CompetitionId,
		this.SimpleMode,
		this.IsLocked,
		this.IsPublic,
		this.MaxUsers)
}
