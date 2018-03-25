package database

import "time"
import "github.com/majestrate/apub"

type User struct {
	Server          string     `json:"-"`
	ID              int64      `json:"id"`
	PrivateKey      string     `json:"privkey"`
	Email           string     `json:"email"`
	LoginCred       string     `json:"login"`
	Name            string     `json:"name"`
	Nick            string     `json:"nick"`
	Profile         string     `json:"profile"`
	Inserted        time.Time  `json:"inserted_at"`
	Updated         time.Time  `json:"updated_at"`
	APID            string     `json:"apid"`
	Avatar          JSONObject `json:"avatar"`
	Info            JSONObject `json:"info"`
	FollowerAddress string     `json:"follower_address"`
	Following       []string   `json:"following"`
}

func (u User) UserInfo() *apub.UserInfo {
	i := &apub.UserInfo{
		UserName: u.Name,
	}
	i.LoadPrivateKey(u.PrivateKey)
	return i
}

func (u User) TableName() string {
	return "local_users"
}

func (u User) TableDef() string {
	return "id SERIAL PRIMARY KEY, privkey TEXT NOT NULL, email VARCHAR(255) NOT NULL, login VARCHAR(255) NOT NULL, name VARCHAR(255) NOT NULL, nick VARCHAR(255) NOT NULL, profile TEXT NOT NULL, inserted_at TIMESTAMP DEFAULT current_timestamp, updated_at TIMESTAMP NOT NULL, apid VARCHAR(255), avatar JSONB NOT NULL, info JSONB NOT NULL, follower_address VARCHAR(255) NOT NULL, following VARCHAR(255)[]"
}
