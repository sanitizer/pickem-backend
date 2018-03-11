package model

import (
	"fmt"
	"time"
	"strings"
	"strconv"
)

const (
	MATCH_INSERT_COLS = "(legacyId, competitionId, awayTeamId, homeTeamId, isLocked, start, stage, week)"
	MATCH_INSERT_PARTIAL_QUERY = "INSERT INTO game_match " + MATCH_INSERT_COLS + " VALUES "
)

type Match struct {
	Id                  int
	LegacyId            string
	CompetitionId       int
	LegacyCompetitionId string
	AwayTeamId          int
	LegacyAwayTeamId    string
	HomeTeamId          int
	LegacyHomeTeamId    string
	IsLocked            bool
	Start               time.Time
	Stage               int
	Week                int
}

func (this Match) String() string {
	return fmt.Sprintf("{ id: %d, legacyId: %s, competitionId: %d, awayTeamId: %d, homeTeamId: %d, isLocked: %t, start: %s, stage: %d, week: %d }",
		this.Id,
		this.LegacyId,
		this.CompetitionId,
		this.AwayTeamId,
		this.HomeTeamId,
		this.IsLocked,
		this.Start.String(),
		this.Stage,
		this.Week)
}
func (this Match) StringForInsert() string {
	var legacyId, competitionId, awayTeamId, homeTeamId, isLocked,
	start, stage, week = "NULL", "NULL", "NULL", "NULL", "NULL", "NULL", "NULL", "NULL"

	if this.LegacyId != "" {
		legacyId = "\"" + this.LegacyId + "\""
	}

	if this.LegacyCompetitionId != "" {
		competitionId = "(select id from competition where legacyId = \"" + this.LegacyCompetitionId + "\")"
	}

	if this.LegacyAwayTeamId != "" {
		awayTeamId = "(select id from team where legacyId = \"" + this.LegacyAwayTeamId + "\")"
	}

	if this.LegacyHomeTeamId != "" {
		homeTeamId = "(select id from team where legacyId = \"" + this.LegacyHomeTeamId + "\")"
	}

	isLocked = strconv.FormatBool(this.IsLocked)
	stage = strconv.Itoa(this.Stage)
	week = strconv.Itoa(this.Week)

	if this.Start.Unix() != int64(-62135596800) {
		start = "\"" + this.Start.Format("2006-01-02 15:04:05") + "\""
	}

	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s)",
		legacyId,
		competitionId,
		awayTeamId,
		homeTeamId,
		isLocked,
		start,
		stage,
		week)
}

func (this Match) BuildFromFirestoreData(data map[string]interface{}) DaoModel {
	var legacyId, competitionId, awayTeamId, homeTeamId string
	var isLocked bool
	var start time.Time
	var stage, week int

	if data["id"] != nil {
		legacyId = data["id"].(string)
	}

	if data["compId"] != nil {
		competitionId = data["compId"].(string)
	}

	if data["awayTeamId"] != nil {
		awayTeamId = data["awayTeamId"].(string)
	}

	if data["homeTeamId"] != nil {
		homeTeamId = data["homeTeamId"].(string)
	}

	if data["isLocked"] != nil {
		isLocked = data["isLocked"].(bool)
	}

	if data["startDate"] != nil {
		start = time.Unix(data["startDate"].(int64), int64(0))
	}

	if data["stage"] != nil {
		stage, _ = strconv.Atoi(strings.Replace(data["stage"].(string), "Stage ", "", -1))
	}

	if data["week"] != nil {
		week, _ = strconv.Atoi(strings.Replace(data["week"].(string), "Week ", "", -1))
	}

	this.LegacyId = legacyId
	this.LegacyCompetitionId = competitionId
	this.LegacyAwayTeamId = awayTeamId
	this.LegacyHomeTeamId = homeTeamId
	this.IsLocked = isLocked
	this.Start = start
	this.Stage = stage
	this.Week = week

	return this
}

func (this Match) GetPartialInsertQuery() string {
	return MATCH_INSERT_PARTIAL_QUERY
}