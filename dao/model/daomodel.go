package model

type DaoModel interface {
	BuildFromFirestoreData(data map[string]interface{}) DaoModel
	StringForInsert() string
	GetPartialInsertQuery() string
}