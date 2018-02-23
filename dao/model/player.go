package model

import (
	"fmt"
)

type Player struct {
	Id        int
	LegacyId  string
	Name      string
	FirstName string
	LastName  string
	Role      string
}

func (this Player) String() string {
	return fmt.Sprintf("{ id: %d, legacyId: %s, name: %s, firstName: %s, lastName: %s, role: %s }",
		this.Id,
		this.LegacyId,
		this.Name,
		this.FirstName,
		this.LastName,
		this.Role)
}
