package model

import (
	"fmt"
	"strconv"
)

const (
	LEADERBOARD_INSERT_COLS = "(type, leagueId, isActive, stage)"
	LEADERBOARD_INSERT_PARTIAL_QUERY = "INSERT INTO leaderboard " + LEADERBOARD_INSERT_COLS + " VALUES "
)

type LeaderBoard struct {
	Id             int
	Type           string
	LeagueId       int
	LegacyLeagueId string
	Stage          int
	IsActive       bool
}

func (this LeaderBoard) String() string {
	return fmt.Sprintf("{ id: %d, type: %s, leagueId: %d, stage: %d, isActive: %t }",
		this.Id,
		this.Type,
		this.LeagueId,
		this.Stage,
		this.IsActive)
}

func (this LeaderBoard) StringForInsert() string {
	typeString, leagueId,
	isActive, stage := "NULL", "NULL", "NULL", "NULL"

	if this.Type != "" {
		typeString = "\"" + this.Type + "\""
	}

	if this.LegacyLeagueId != "" {
		leagueId = "(SELECT IFNULL((select id from league where legacyId = \"" + this.LegacyLeagueId + "\"), -1))"
	}

	isActive = strconv.FormatBool(this.IsActive)
	stage = strconv.Itoa(this.Stage)

	return fmt.Sprintf("(%s, %s, %s, %s)",
		typeString,
		leagueId,
		isActive,
		stage)
}

func (this LeaderBoard) BuildFromFirestoreData(data map[string]interface{}) DaoModel {
	var typeString, legacyLeagueId string
	var isActive bool
	var stage int

	if data["leagueId"] != nil {
		legacyLeagueId = data["leagueId"].(string)
	}

	if data["type"] != nil {
		typeString = data["type"].(string)
	}

	if data["active"] != nil {
		isActive = data["active"].(bool)
	}

	if data["stage"] != nil {
		stage = data["stage"].(int)
	}

	this.LegacyLeagueId = legacyLeagueId
	this.Type = typeString
	this.IsActive = isActive
	this.Stage = stage

	return this
}

func (this LeaderBoard) GetPartialInsertQuery() string {
	return LEADERBOARD_INSERT_PARTIAL_QUERY
}