package service

import (
	"toy/marmot/web/query-log/controller/response"
	"toy/marmot/web/query-log/model/entity"
	"toy/marmot/web/query-log/model/mysql"
	"toy/marmot/web/query-log/model/Dao"
	"log"
)

func GetLogData(query *entity.LogQuery)([]response.LogResultIn, error){

	var logDao dao.LogDao
	logDao = new(mysql.LogDaoMysql)
	logData,err := logDao.GetLogData(query)
	if err != nil {
		log.Printf("fetch error :%v", err)
		return nil,err
	}

	resultList := make([]response.LogResultIn, len(logData))
	uidList := make([]int64,0)
	for i,info := range logData{
		uidList = append(uidList, info.LogId)
		resultList[i] = response.LogResultIn{Log:info}
	}

	var lcd dao.LogContentDao
	lcd = new(mysql.LogContentDaoMysql)
	logContentData,err := lcd.GetLogContentByIds(uidList, logData[0].Uid)
	if err == nil {
		for k,resultInfo := range resultList{
			logId := resultInfo.Log.LogId
			for _,logContentInfo := range logContentData{
				if logId == logContentInfo.Id{
					resultList[k].LogContent = logContentInfo
				}
			}
		}
	}
	return resultList,nil
}