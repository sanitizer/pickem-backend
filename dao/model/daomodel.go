package model

type DaoModel interface {
	Build(data map[string]interface{}) DaoModel
	StringForInsert() string
	GetPartialInsertQuery() string
}