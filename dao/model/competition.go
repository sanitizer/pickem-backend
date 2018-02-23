package model

import (
	"fmt"
	"time"
)

type Competition struct {
	Id          int
	LegacyId    string
	WinstonsId  int
	Name        string
	Type        string
	Organizer   string
	Description string
	Start       time.Time
	End         time.Time
	ImageRef    time.Time
	IsActive    bool
}

func (this Competition) String() string {
	return fmt.Sprintf("{ id: %d, legacyId: %s, winstonsId: %d, name: %s, type: %s, organizer: %s, description: %s, start: %s, end: %s, imageRef: %s, isActive: %t }",
		this.Id,
		this.LegacyId,
		this.WinstonsId,
		this.Name,
		this.Type,
		this.Organizer,
		this.Description,
		this.Start.String(),
		this.End.String(),
		this.ImageRef,
		this.IsActive)
}
