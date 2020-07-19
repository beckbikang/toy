package mysql


import (
	"fmt"
	log "log"

	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"

	"toy/marmot/web/query-log-echo/model/entity"
	"toy/marmot/web/query-log-echo/launch/db"

)


const (
	LogContentTablePre = "log_content_%d"
	LogContentTpl             = 2
)

type LogContentDaoMysql struct{}

//get table name
func (lcdm *LogContentDaoMysql) getTableName(id int64) string {
	return fmt.Sprintf(LogContentTablePre, id%LogContentTpl)
}

//get content
func (lcdm *LogContentDaoMysql) GetLogContent(id int64, uid int64) (*entity.LogContentEntity,error ) {
	tableName := lcdm.getTableName(uid)

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
		log.Printf("build error %v", err)
		return nil, err
	}

	//fetch data
	logdb := db.GetDb()

	log.Printf("%s <= %v ", cond, values)
	rows, err := logdb.Query(cond, values...)
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


//get
func(lcdm *LogContentDaoMysql) GetLogContentByIds(ids []int64, uid int64) ( []entity.LogContentEntity,error){
	tableName := lcdm.getTableName(uid)
	where := map[string]interface{}{
		"id in":ids,
	}
	selectFields := []string{
		"id", "cfrom", "cto", "mtime",
	}
	cond,values,err := builder.BuildSelect(tableName, where, selectFields)
	if err != nil{
		log.Printf("build error %v", err)
		return nil,err
	}
	//fetch data
	logdb := db.GetDb()
	log.Printf("%s => %v ",cond, values)
	rows, err := logdb.Query(cond, values...)
	if err != nil {
		log.Printf("fetch result: %v",err)
		return nil,err
	}
	defer rows.Close()

	//get data
	var logContents []entity.LogContentEntity
	err = scanner.Scan(rows, &logContents)
	if err != nil {
		log.Printf("scan result: %v",err)
		return nil, err
	}

	return logContents, nil
}



//add content
func (lcdm *LogContentDaoMysql) Save(id uint64, from, to string,  uid int64) bool {
	return false
}

//add entity
func (lcdm *LogContentDaoMysql) SaveEntity(lce *entity.LogContentEntity, uid int64) bool{
	return false
}
