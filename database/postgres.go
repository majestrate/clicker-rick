package database

import (
	//"fmt"
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
	/*
		tables := map[string]string{
			"actors": "",
		}
		table_order := []string{
			"actors",
		}

		for _, table := range table_order {
			p.conn.MustExec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ( %s )", table, tables[table]))
		}
	*/
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
	logrus.Infof("find local user %s", username)
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
	db.conn, err = sqlx.Connect("postgres", url)
	if err == nil {
		db.conn.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	} else {
		db = nil
	}
	return
}
