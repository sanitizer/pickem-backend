package model

import (
	"fmt"
	"strings"
)

const (
	TEAM_INSERT_COLS = "(legacyId, name, shortName, logo)"
	TEAM_INSERT_PARTIAL_QUERY = "INSERT INTO team " + TEAM_INSERT_COLS + " VALUES "
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

func (this Team) StringForInsert() string {
	var legacy = "NULL"
	var name = "NULL"
	var shortName = "NULL"
	var logo = "NULL"

	if this.LegacyId != "" {
		legacy = "\"" + this.LegacyId + "\""
	}

	if this.Name != "" {
		name = "\"" + this.Name + "\""
	}

	if this.ShortName != "" {
		shortName = "\"" + this.ShortName + "\""
	}

	if this.Logo != "" {
		logo = "\"" + this.Logo + "\""
	}

	return fmt.Sprintf("(%s, %s, %s, %s)",
		legacy,
		name,
		shortName,
		logo)
}

func (this Team) BuildFromFirestoreData(data map[string]interface{}) DaoModel {
	var legacy = ""
	var name = ""
	var shortName = ""
	var logo = ""

	if data["id"] != nil {
		legacy = data["id"].(string)
	}

	if data["name"] != nil {
		name = data["name"].(string)
	}

	if data["shortName"] != nil {
		shortName = data["shortName"].(string)
	}

	if data["logo"] != nil {
		logo = data["logo"].(string)
	}

	this.LegacyId = legacy
	this.Name = strings.Replace(name, "\"", "'", -1)
	this.ShortName = strings.Replace(shortName, "\"", "'", -1)
	this.Logo = strings.Replace(logo, "\"", "'", -1)

	return this
}

func (this Team) GetPartialInsertQuery () string {
	return TEAM_INSERT_PARTIAL_QUERY
}