package model

import (
	"fmt"
	"time"
	"strings"
	"strconv"
)

const (
	COMPETITION_INSERT_COLS = "(legacyId, winstonsId, name, type, organizer, description, start, end, imageRef, isActive)"
	COMPETITION_INSERT_PARTIAL_QUERY = "INSERT INTO competition " + COMPETITION_INSERT_COLS + " VALUES "
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
	ImageRef    string
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

func (this Competition) StringForInsert() string {
	var legacy = "NULL"
	var winstons = "NULL"
	var name = "NULL"
	var cType = "NULL"
	var organizer = "NULL"
	var description = "NULL"
	var start = "NULL"
	var end = "NULL"
	var imageRef = "NULL"
	var isActive = "NULL"

	if this.LegacyId != "" {
		legacy = "\"" + this.LegacyId + "\""
	}

	winstons = strconv.Itoa(this.WinstonsId)

	if this.Name != "" {
		name = "\"" + this.Name + "\""
	}

	if this.Type != "" {
		cType = "\"" + this.Type + "\""
	}

	if this.Organizer != "" {
		organizer = "\"" + this.Organizer + "\""
	}

	if this.ImageRef != "" {
		imageRef = "\"" + this.ImageRef + "\""
	}

	isActive = strconv.FormatBool(this.IsActive)

	if this.Description != "" {
		description = "\"" + this.Description + "\""
	}

	if this.Start.Unix() != int64(-62135596800) {
		start = "\"" + this.Start.Format("2006-01-02") + "\""
	}

	if this.End.Unix() != int64(-62135596800) {
		end = "\"" + this.End.Format("2006-01-02") + "\""
	}

	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)",
		legacy,
		winstons,
		name,
		cType,
		organizer,
		description,
		start,
		end,
		imageRef,
		isActive)
}

func (this Competition) Build(data map[string]interface{}) DaoModel {
	var legacyId = ""
	var winstons = 0
	var name = ""
	var cType = ""
	var organizer = ""
	var description = ""
	var start = time.Time{}
	var end = time.Time{}
	var imageRef = ""
	var isActive = false

	if data["id"] != nil {
		legacyId = data["id"].(string)
	}

	if data["winstonsId"] != nil {
		winstons = int(data["winstonsId"].(int64))
	}

	if data["name"] != nil {
		name = data["name"].(string)
	}

	if data["type"] != nil {
		cType = data["type"].(string)
	}

	if data["organizer"] != nil {
		organizer = data["organizer"].(string)
	}

	if data["description"] != nil {
		description = data["description"].(string)
	}

	if data["startDate"] != nil {
		start, _ = time.Parse("20060102", strconv.Itoa(int(data["startDate"].(int64))))
	}

	if data["endDate"] != nil {
		end, _ = time.Parse("20060102", strconv.Itoa(int(data["endDate"].(int64))))
	}

	if data["image"] != nil {
		imageRef = data["image"].(string)
	}

	if data["isActive"] != nil {
		isActive = data["isActive"].(bool)
	}

	this.LegacyId = legacyId
	this.WinstonsId = winstons
	this.Name = strings.Replace(name, "\"", "'", -1)
	this.Type = strings.Replace(cType, "\"", "'", -1)
	this.Organizer = strings.Replace(organizer, "\"", "'", -1)
	this.Description = strings.Replace(description, "\"", "'", -1)
	this.Start = start
	this.End = end
	this.ImageRef = imageRef
	this.IsActive = isActive

	return this
}

func (this Competition) GetPartialInsertQuery () string {
	return COMPETITION_INSERT_PARTIAL_QUERY
}