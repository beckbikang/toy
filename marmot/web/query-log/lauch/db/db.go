package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"sync"
	"toy/marmot/web/query-log/lauch/config"

	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
)

var (
	globalDb *sql.DB
	dbOnce   sync.Once

	err error
)

//init db
func InitDb() {
	dbOnce.Do(func() {
		host := config.Gcfg.GetString("mysql.host")
		port := config.Gcfg.GetInt("mysql.port")
		user := config.Gcfg.GetString("mysql.user")
		password := config.Gcfg.GetString("mysql.pswd")
		db := config.Gcfg.GetString("mysql.db")
		charset := config.Gcfg.GetString("mysql.charset")
		locale := config.Gcfg.GetString("mysql.locale")

		globalDb, err = manager.New(db, user, password, host).Set(
			manager.SetCharset(charset),
			manager.SetParseTime(true),
			manager.SetLoc(url.QueryEscape(locale))).Port(port).Open(true)

		if err != nil {
			fmt.Printf("connect db error happen error:%v", err)
			panic(-1)
		}
	})
}

//get db
func GetDb() *sql.DB {
	return globalDb
}
