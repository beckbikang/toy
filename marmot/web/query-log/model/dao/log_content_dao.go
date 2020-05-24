package dao

import (
	"fmt"
	log "log"
	"toy/marmot/web/query-log/dao/entity"

	"toy/marmot/web/query-log/lauch/db"

	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
)

type LogContentDao interface {
	GetLogContent(id uint64) (error, *entity.LogContentEntity)
	Save(id uint64, from, to string) bool
}

const (
	LogContentTablePre = "log_content_%d"
	LogTpl             = 2
)

type LogContentDaoMysql struct{}

//get table name
func (lcdm *LogContentDaoMysql) getTableName(id uint64) string {
	return fmt.Sprintf(LogContentTablePre, id/LogTpl)
}

//get content
func (lcdm *LogContentDaoMysql) GetLogContent(id uint64) (error, *entity.LogContentEntity) {
	tableName := lcdm.getTableName(id)

	where := map[string]interface{}{
		"id": id,
	}
	selectFields := []string{
		"id",
		"from",
		"to",
		"mtime",
	}
	cond, values, err := builder.BuildSelect(tableName,
		where, selectFields)
	if err != nil {
		log.Printf("build error %v", err)
		return nil, err
	}

	//fetch data
	db := db.GetDb()

	log.Printf("%s <= %v ", cond, values)
	rows, err := db.Query(cond, values...)
	if err != nil {
		log.Printf("fetch result: %v", err)
		return nil, err
	}
	defer rows.Close()

	//get data
	var logContentEntity entity.LogContentEntity
	scanner.Scan(rows, &logContentEntity)
	return &logContentEntity, nil
}

//add content
func (lcdm *LogContentDaoMysql) Save(
	id uint64, from, to string,
) bool {
	return false
}
