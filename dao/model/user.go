package model

import (
	"fmt"
	"time"
	"strings"
)

const (
	USER_INSERT_COLS = "(legacyId, battleNetId, discordId, displayName, email, gravatar, teamLogo, lastActive)"
	USER_INSERT_PARTIAL_QUERY = "INSERT INTO app_user " + USER_INSERT_COLS + " "
)

type User struct {
	Id          int
	LegacyId    string
	BattleNetId string
	DiscordId   string
	DisplayName string
	Email       string
	GrAvatar    string
	TeamLogo    string
	LastActive  time.Time
}

func (this User) String() string {
	return fmt.Sprintf("{ id: %d, legacyId: %s, battleNetId: %s, discordId: %s, displayName: %s, email: %s, gravatar: %s, teamLogo: %s, lastActive: %d }",
		this.Id,
		this.LegacyId,
		this.BattleNetId,
		this.DiscordId,
		this.DisplayName,
		this.Email,
		this.GrAvatar,
		this.TeamLogo,
		this.LastActive.Unix())
}

func (this User) StringForInsert() string {
	var legacy = "NULL"
	var battleNet = "NULL"
	var discord = "NULL"
	var display = "NULL"
	var email = "NULL"
	var gravatar = "NULL"
	var teamLogo = "NULL"
	var lastActive = "NULL"

	if this.LegacyId != "" {
		legacy = "\"" + this.LegacyId + "\""
	}

	if this.BattleNetId != "" {
		battleNet = "\"" + this.BattleNetId + "\""
	}

	if this.DisplayName != "" {
		display = "\"" + this.DisplayName + "\""
	}

	if this.Email != "" {
		email = "\"" + this.Email + "\""
	}

	if this.GrAvatar != "" {
		gravatar = "\"" + this.GrAvatar + "\""
	}

	if this.TeamLogo != "" {
		teamLogo = "\"" + this.TeamLogo + "\""
	}

	if this.LastActive.Unix() != int64(-62135596800) {
		lastActive = "\"" + this.LastActive.Format("2006-01-02 15:04:05") + "\""
	}

	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s)",
		legacy,
		battleNet,
		discord,
		display,
		email,
		gravatar,
		teamLogo,
		lastActive)
}

func (this User) BuildFromFirestoreData(data map[string]interface{}) DaoModel {
	var battleNet = ""
	var discord = ""
	var display = ""
	var email = ""
	var gravatar = ""
	var teamLogo = ""
	var lastActive time.Time
	var legacyId = ""

	if data["battleNet"] != nil {
		battleNet = data["battleNet"].(string)
	}

	if data["discord"] != nil {
		discord = data["discord"].(string)
	}

	if data["displayName"] != nil {
		display = data["displayName"].(string)
	}

	if data["email"] != nil {
		email = data["email"].(string)
	}

	if data["gravatar"] != nil {
		gravatar = data["gravatar"].(string)
	}

	if data["teamLogo"] != nil {
		teamLogo = data["teamLogo"].(string)
	}

	if data["lastActive"] != nil {
		lastActive = time.Unix(data["lastActive"].(int64), int64(0))
	}

	if data["id"] != nil {
		legacyId = data["id"].(string)
	}

	this.BattleNetId = strings.Replace(battleNet, "\"", "'", -1)
	this.DiscordId = strings.Replace(discord, "\"", "'", -1)
	this.DisplayName = strings.Replace(display, "\"", "'", -1)
	this.Email = strings.Replace(email, "\"", "'", -1)
	this.GrAvatar = strings.Replace(gravatar, "\"", "'", -1)
	this.TeamLogo = strings.Replace(teamLogo, "\"", "'", -1)
	this.LastActive = lastActive
	this.LegacyId = legacyId

	return this
}

func (this User) GetPartialInsertQuery () string {
	return USER_INSERT_PARTIAL_QUERY
}