package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/majestrate/apub"
)

type PostgresDB struct {
	conn     *sql.DB
	Hostname string
}

func (p *PostgresDB) LocalHost() string {
	return p.Hostname
}

func (p *PostgresDB) Close() error {
	return p.conn.Close()
}

func (p *PostgresDB) Init() error {
	return nil
}

func (p *PostgresDB) ListFollowers(userid string) (followers []*apub.UserInfo, err error) {
	return
}

func (p *PostgresDB) ListFollowing(userid string) (following []*apub.UserInfo, err error) {
	return
}

func (p *PostgresDB) LocalPost(postid string) (post *apub.Post, err error) {
	post = nil
	return
}

func (p *PostgresDB) LocalUser(username string) (user *apub.UserInfo, err error) {
	user = nil
	return
}

func (p *PostgresDB) LocalUserPosts(username string, offset int64, limit int) (posts []*apub.Post, err error) {
	return
}

func newPostgresDB(url, localhostname string) (db *PostgresDB, err error) {
	db = &PostgresDB{
		Hostname: localhostname,
	}
	db.conn, err = sql.Open("postgres", url)
	if err != nil {
		db = nil
	}
	return
}
