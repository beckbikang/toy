package dao

import (
	"testing"
	"toy/marmot/web/query-log/lauch/config"
	"toy/marmot/web/query-log/lauch/db"

	"toy/marmot/web/query-log/lauch/config"

	"github.com/stretchr/testify/assert"
)

func TestGetLogContent(t *testing.T) {
	config.LoadGlobalConfig("../conf", "dev")
	db.InitDb()

	var lcd LogContentDao

	lcd = new(LogContentDaoMysql)

	info, _ := lcd.GetLogContent(1)
	assert.NotNil(t, info)
}
