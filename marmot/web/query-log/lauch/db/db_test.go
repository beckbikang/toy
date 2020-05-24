package db

import (
	"testing"
	"toy/marmot/web/query-log/lauch/config"

	"github.com/stretchr/testify/assert"
)

//test get db
func TestGetDb(t *testing.T) {
	config.LoadGlobalConfig("../../conf", "dev")
	InitDb()

	db := GetDb()
	err := db.Ping()
	assert.Nil(t, err, "err is not nil")
}
