package database

import (
	"github.com/jmoiron/sqlx/types"
)

type JSONObject = types.NullJSONText

type Actor struct {
	ID              int64      `json:"id"`
	PrivateKey      string     `json:"privkey"`
	Email           string     `json:"email"`
	LoginCred       string     `json:"login"`
	Name            string     `json:"name"`
	Nickname        string     `json:"nickname"`
	Profile         string     `json:"profile"`
	Local           bool       `json:"local"`
	Inserted        int64      `json:"inserted_at"`
	Updated         int64      `json:"updated_at"`
	APID            string     `json:"apid"`
	Avatar          JSONObject `json:"avatar"`
	Info            JSONObject `json:"info"`
	FollowerAddress string     `json:"follower_address"`
	Following       []string   `json:"following"`
}

type Object struct {
	ID   int64      `json:"id"`
	Data JSONObject `json:"data"`
}
