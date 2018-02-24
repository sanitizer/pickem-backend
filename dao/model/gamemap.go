package model

import (
	"fmt"
)

type GameMap struct {
	Id       int
	Name     string
	Type     string
}

func (this GameMap) String() string {
	return fmt.Sprintf("{ id: %d, name: %s, type: %s }",
		this.Id,
		this.Name,
		this.Type)
}
