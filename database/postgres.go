package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	_ "github.com/lib/pq"
	"github.com/majestrate/apub"
	"github.com/sirupsen/logrus"
	"strings"
)

type PostgresDB struct {
	conn     *sqlx.DB
	Hostname string
}

func (p *PostgresDB) LocalHost() string {
	return p.Hostname
}

func (p *PostgresDB) Close() error {
	return p.conn.Close()
}

func (p *PostgresDB) Init() error {
	logrus.Info("Initialize database tables....")
	entities := []Model{User{}, Object{}}

	for idx := range entities {
		q := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ( %s )", entities[idx].TableName(), entities[idx].TableDef())
		logrus.Infof("query: %s", q)
		p.conn.MustExec(q)
	}

	return nil
}

func (p *PostgresDB) ListFollowers(userid string) (followers []apub.User, err error) {
	return
}

func (p *PostgresDB) ListFollowing(userid string) (following []apub.User, err error) {
	return
}

func (p *PostgresDB) LocalPost(postid string) (post *apub.Post, err error) {
	post = nil
	return
}

func (p *PostgresDB) LocalUser(username string) (user apub.User, err error) {
	logrus.Infof("find local user %s", username)
	var users []User
	err = p.conn.Select(&users, "SELECT * FROM local_users WHERE name=$1 LIMIT 1", username)
	if err == nil && len(users) > 0 {
		user = users[0].ToInfo()
	}
	return
}

func (p *PostgresDB) LocalUserPosts(username string, offset int64, limit int) (posts []*apub.Post, err error) {
	return
}

func newPostgresDB(url, localhostname string) (db *PostgresDB, err error) {
	db = &PostgresDB{
		Hostname: localhostname,
	}
	db.conn, err = sqlx.Connect("postgres", url)
	if err == nil {
		db.conn.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	} else {
		db = nil
	}
	return
}
