package mysql

import (
	"testing"
	"toy/marmot/web/query-log/launch/config"
	"toy/marmot/web/query-log/launch/db"
	"toy/marmot/web/query-log/model/entity"
	"toy/marmot/web/query-log/model/dao"
	"github.com/stretchr/testify/assert"
)

func TestGetLogData(t *testing.T) {
	config.LoadGlobalConfig("../../conf", "dev")
	db.InitDb()

	var logDao dao.LogDao
	logDao = new(LogDaoMysql)

	var queryLog *entity.LogQuery
	queryLog = new(entity.LogQuery)
	queryLog.Uid = 123

	infos, err := logDao.GetLogData(queryLog)
	assert.Nil(t, err)
	assert.NotNil(t, infos)

}