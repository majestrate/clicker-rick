package database

import (
	"github.com/jmoiron/sqlx/types"
)

type JSONObject = types.NullJSONText

type Model interface {
	TableName() string
	TableDef() string
}
