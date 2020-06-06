package db

import (
	"testing"
	"toy/marmot/web/query-log-echo/launch/config"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"

	"toy/marmot/web/query-log-echo/model/entity"

	"github.com/stretchr/testify/assert"

)

//test get db
func TestGetDb(t *testing.T) {
	config.LoadGlobalConfig("../../conf", "dev")
	InitDb()

	db := GetDb()
	err := db.Ping()
	assert.Nil(t, err, "err is not nil")

	id := 1
	tableName := "log_content_1"

	where := map[string]interface{}{
		"id": id,
	}
	selectFields := []string{
		"id",
		"cfrom",
		"cto",
		"mtime",
	}
	cond, values, err := builder.BuildSelect(tableName,
		where, selectFields)
	if err != nil {
		t.Logf("build error %v", err)
	}

	//fetch data

	t.Logf("%s <= %v ", cond, values)
	rows, err := db.Query(cond, values...)
	if err != nil {
		t.Logf("fetch result: %v", err)
	}
	defer rows.Close()

	//get data
	var logContentEntity entity.LogContentEntity
	scanner.Scan(rows, &logContentEntity)


}
