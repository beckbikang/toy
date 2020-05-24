package mysql

import (
	"toy/marmot/web/query-log/util"
	"toy/marmot/web/query-log/model/entity"
	log "log"
	"fmt"
	"github.com/didi/gendry/builder"
	"toy/marmot/web/query-log/launch/db"
	"github.com/didi/gendry/scanner"
)

const (
	LogTablePre = "log_%d"
	LogTpl             = 2
)

type LogDaoMysql struct{}

func(ldm *LogDaoMysql)getTableName(uid int64)string{
	return fmt.Sprintf(LogTablePre, uid%LogTpl)
}


//get data from mysql database
func(ldm *LogDaoMysql) GetLogData(query *entity.LogQuery) ([]entity.LogEntity,error){

	//uid
	uid := query.Uid
	if uid <= 0 {
		return nil, fmt.Errorf("uid is empty")
	}

	//log type
	logType := query.LogType


	targetId := query.LogTargetId

	//start time end time
	startTime := query.StartTime
	endTime := query.EndTime

	if startTime == ""{
		startTime = util.GetBeforeDate(7)
	}

	if endTime == ""{
		endTime = util.GetNextDate()
	}

	//where
	where := map[string]interface{}{
		"uid": uid,
		"mtime >=": startTime,
		"mtime <": endTime,
	}
	if targetId > 0{
		where["log_target_id"] = targetId
	}
	if logType > 0{
		where["log_type"] = logType
	}

	selectFields := []string{
		"id",
		"uid",
		"log_type",
		"log_target_id",
		"log_id",
		"mtime",
	}

	tableName := ldm.getTableName(uid)


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
	var logEntitys []entity.LogEntity
	scanner.Scan(rows, &logEntitys)
	return logEntitys, nil







}