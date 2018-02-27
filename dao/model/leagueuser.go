package model

import (
	"fmt"
	"strconv"
)

const (
	LEAGUEUSER_INSERT_COLS = "(legacyId, userId, leagueId, isActive)"
	LEAGUEUSER_INSERT_PARTIAL_QUERY = "INSERT INTO league_user " + LEAGUEUSER_INSERT_COLS + " VALUES "
)

type LeagueUser struct {
	Id             int
	LegacyId       string
	UserId         int
	LegacyUserId   string
	LeagueId       int
	LegacyLeagueId string
	IsActive       bool
}

func (this LeagueUser) String() string {
	return fmt.Sprintf("{ id: %d, legacyId: %s, userId: %d, leagueId: %d, isActive: %t }",
		this.Id,
		this.LegacyId,
		this.UserId,
		this.LeagueId,
		this.IsActive)
}

func (this LeagueUser) StringForInsert() string {
	legacyId, legacyUserId,
	legacyLeagueId, isActive := "NULL", "NULL", "NULL", "NULL"

	if this.LegacyId != "" {
		legacyId = "\"" + this.LegacyId + "\""
	}

	if this.LegacyLeagueId != "" {
		legacyLeagueId = "(SELECT IFNULL((select id from league where legacyId = \"" + this.LegacyLeagueId + "\"), -1))"
	}

	if this.LegacyUserId != "" {
		legacyUserId = "(SELECT IFNULL((select id from app_user where legacyId = \"" + this.LegacyUserId + "\"), -1))"
	}

	isActive = strconv.FormatBool(this.IsActive)

	return fmt.Sprintf("(%s, %s, %s, %s)",
		legacyId,
		legacyUserId,
		legacyLeagueId,
		isActive)
}

func (this LeagueUser) BuildFromFirestoreData(data map[string]interface{}) DaoModel {
	var legacyId, legacyUserId, legacyLeagueId string
	var isActive bool

	if data["id"] != nil {
		legacyId = data["id"].(string)
	}

	if data["leagueId"] != nil {
		legacyLeagueId = data["leagueId"].(string)
	}

	if data["userId"] != nil {
		legacyUserId = data["userId"].(string)
	}

	if data["active"] != nil {
		isActive = data["active"].(bool)
	}

	this.LegacyId = legacyId
	this.LegacyUserId = legacyUserId
	this.LegacyLeagueId = legacyLeagueId
	this.IsActive = isActive

	return this
}

func (this LeagueUser) GetPartialInsertQuery() string {
	return LEAGUEUSER_INSERT_PARTIAL_QUERY
}