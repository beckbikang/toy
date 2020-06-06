package mysql

import (
	"testing"
	"toy/marmot/web/query-log-echo/launch/config"
	"toy/marmot/web/query-log-echo/launch/db"
	"toy/marmot/web/query-log-echo/model/dao"
	"github.com/stretchr/testify/assert"

)

func TestGetLogContent(t *testing.T) {
	config.LoadGlobalConfig("../../conf", "dev")
	db.InitDb()


	var lcd dao.LogContentDao

	lcd = new(LogContentDaoMysql)

	info,_  := lcd.GetLogContent(1, 123)
	assert.NotNil(t, info)
}

func TestGetLogContentByIds(t *testing.T) {
	config.LoadGlobalConfig("../../conf", "dev")
	db.InitDb()


	var lcd dao.LogContentDao

	lcd = new(LogContentDaoMysql)

	ids := []int64{1,2,3}

	info,_  := lcd.GetLogContentByIds(ids, 123)
	assert.NotNil(t, info)
}