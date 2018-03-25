package database

type Object struct {
	ID   int64      `json:"id"`
	Data JSONObject `json:"data"`
}

func (o Object) TableName() string {
	return "objects"
}

func (o Object) TableDef() string {
	return "id BIG SERIAL PRIMARY KEY, data JSONB NOT NULL"
}
