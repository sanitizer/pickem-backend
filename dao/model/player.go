package model

import (
	"fmt"
	"strings"
)

const (
	PLAYER_INSERT_COLS = "(name, fullName)"
	PLAYER_INSERT_PARTIAL_QUERY = "INSERT INTO player " + PLAYER_INSERT_COLS + " VALUES "
)

type Player struct {
	Id        int
	Name      string
	FullName string
	Role      string
}

func (this Player) String() string {
	return fmt.Sprintf("{ id: %d, name: %s, fullName: %s, role: %s }",
		this.Id,
		this.Name,
		this.FullName,
		this.Role)
}

func (this Player) StringForInsert() string {
	var name = "NULL"
	var fullName = "NULL"

	if this.Name != "" {
		name = "\"" + this.Name + "\""
	}

	if this.FullName != "" {
		fullName = "\"" + this.FullName + "\""
	}

	return fmt.Sprintf("(%s, %s)",
		name,
		fullName)
}

func (this Player) BuildFromFirestoreData(data map[string]interface{}) DaoModel {
	var name = ""
	var fullName = ""

	if data["player"] != nil {
		name = data["player"].(string)
	}

	if data["realName"] != nil {
		fullName = data["realName"].(string)
	}

	this.Name = strings.Replace(name, "\"", "'", -1)
	this.FullName = strings.Replace(fullName, "\"", "'", -1)

	return this
}

func (this Player) GetPartialInsertQuery () string {
	return PLAYER_INSERT_PARTIAL_QUERY
}