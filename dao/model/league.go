package model

import (
	"fmt"
	"strconv"
	"strings"
	"log"
)

const (
	LEAGUE_INSERT_COLS = "(legacyId, name, description, ownerId, competitionId, isLocked, isPublic, maxUsers, simpleMode)"
	LEAGUE_INSERT_PARTIAL_QUERY = "INSERT INTO league " + LEAGUE_INSERT_COLS + " VALUES "
)

type League struct {
	Id                  int
	LegacyId            string
	Name                string
	Description         string
	OwnerId             int
	LegacyOwnerId       string
	CompetitionId       int
	LegacyCompetitionId string
	SimpleMode          bool
	IsLocked            bool
	IsPublic            bool
	MaxUsers            int64
}

func (this League) String() string {
	return fmt.Sprintf("{ id: %d, legacyId: %s, name: %s, description: %s, ownerId: %d, competitionId: %d, simpleMode: %t, isLocked: %t, isPublic: %t, maxUsers: %d }",
		this.Id,
		this.LegacyId,
		this.Name,
		this.Description,
		this.OwnerId,
		this.CompetitionId,
		this.SimpleMode,
		this.IsLocked,
		this.IsPublic,
		this.MaxUsers)
}

func (this League) StringForInsert() string {
	legacyId, name,
	description, legacyOwnerId,
	legacyCompeteId, isLocked,
	isPublic, maxUsers,
	simpleMode := "NULL", "NULL", "NULL", "NULL", "NULL", "NULL", "NULL", "NULL", "NULL"

	if this.LegacyId != "" {
		legacyId = "\"" + this.LegacyId + "\""
	}

	if this.Name != "" {
		name = "\"" + this.Name + "\""
	}

	if this.Description != "" {
		description = "\"" + this.Description + "\""
	}

	if this.LegacyOwnerId != "" {
		legacyOwnerId = "(SELECT IFNULL((select id from app_user where legacyId = \"" + this.LegacyOwnerId + "\"), -1))"
	}

	if this.LegacyCompetitionId != "" {
		legacyCompeteId = "(SELECT IFNULL((select id from competition where legacyId = \"" + this.LegacyCompetitionId + "\"), -1))"
	}

	isLocked = strconv.FormatBool(this.IsLocked)
	isPublic = strconv.FormatBool(this.IsPublic)
	simpleMode = strconv.FormatBool(this.SimpleMode)

	maxUsers = strconv.Itoa(int(this.MaxUsers))

	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s)",
		legacyId,
		name,
		description,
		legacyOwnerId,
		legacyCompeteId,
		isLocked,
		isPublic,
		maxUsers,
		simpleMode)
}

func (this League) BuildFromFirestoreData(data map[string]interface{}) DaoModel {
	var legacyId, name, description, legacyOwnerId, legacyCompeteId string
	var isLocked, isPublic, simpleMode bool
	var maxUsers int64

	if data["id"] != nil {
		legacyId = data["id"].(string)
	}

	if data["competitionId"] != nil {
		legacyCompeteId = data["competitionId"].(string)
	}

	if data["name"] != nil {
		name = data["name"].(string)
	}

	if data["motd"] != nil {
		description = data["motd"].(string)
	}

	if data["ownerId"] != nil {
		legacyOwnerId = data["ownerId"].(string)
	}

	if data["isLocked"] != nil {
		isLocked = data["isLocked"].(bool)
	}

	if data["isPublic"] != nil {
		isPublic = data["isPublic"].(bool)
	}

	if data["simpleMode"] != nil {
		simpleMode = data["simpleMode"].(bool)
	}

	if data["maxUsers"] != nil {
		users, err := strconv.Atoi(fmt.Sprintf("%v", data["maxUsers"]))
		if err != nil {
			log.Fatal(err.Error())
		}
		maxUsers = int64(users)
	}

	this.LegacyId = legacyId
	this.Name = strings.Replace(name, "\"", "'", -1)
	this.Description = strings.Replace(description, "\"", "'", -1)
	this.LegacyOwnerId = legacyOwnerId
	this.LegacyCompetitionId = legacyCompeteId
	this.IsLocked = isLocked
	this.IsPublic = isPublic
	this.SimpleMode = simpleMode
	this.MaxUsers = maxUsers

	return this
}

func (this League) GetPartialInsertQuery() string {
	return LEAGUE_INSERT_PARTIAL_QUERY
}