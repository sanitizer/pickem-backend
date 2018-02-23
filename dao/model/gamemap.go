package model

import (
	"fmt"
)

type GameMap struct {
	Id       int
	LegacyId string
	Name     string
	Type     string
}

func (this GameMap) String() string {
	return fmt.Sprintf("{ id: %d, legacyId: %s, name: %s, type: %s }",
		this.Id,
		this.LegacyId,
		this.Name,
		this.Type)
}
