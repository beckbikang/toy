package service

import (
	"log"
	"fmt"
	"toy/marmot/web/query-log-echo/controller/response"
	"toy/marmot/web/query-log-echo/model/Dao"
	"toy/marmot/web/query-log-echo/model/entity"
	"toy/marmot/web/query-log-echo/model/mysql"
	"sort"
	"toy/marmot/web/query-log-echo/util"
	"toy/marmot/web/query-log-echo/launch/cache"
	klog "toy/marmot/web/query-log-echo/launch/log"
	"encoding/json"
)

func GetLogData(query *entity.LogQuery) ([]response.LogResultIn, error) {

	var logDao dao.LogDao
	var err error
	var logData []entity.LogEntity

	logDao = new(mysql.LogDaoMysql)
	logData, err = logDao.GetLogData(query)
	if err != nil || len(logData) <= 0 {
		log.Printf("fetch error :%v", err)
		return nil, fmt.Errorf("fetch empty")
	}

	resultList := make([]response.LogResultIn, len(logData))
	//uidList := make([]int64, 0)
	uidList := util.Int64List{}
	for i, info := range logData {
		uidList = append(uidList, info.LogId)
		resultList[i] = response.LogResultIn{LogEntity:info}
	}
	sort.Sort(uidList)

	var redisKey string

	var logContentData []entity.LogContentEntity

	//add cache
	redisOk := false
	redisKey,err = util.GetHashSliceInt64(uidList)
	//klog.LOGGER.Info("redis key",zap.String("key",redisKey))
	if err == nil {
		var cacheVal []byte
		cacheCmd := cache.CommonRedis.Get(redisKey)
		if cacheCmd.Err() == nil {
			cacheVal,err = cacheCmd.Bytes()
			if err == nil {
				err = json.Unmarshal(cacheVal, logContentData)
				if err == nil {
					//klog.LOGGER.Info("get redis success")
					redisOk = true
				}
			}
		}
	}


	if !redisOk{
		var lcd dao.LogContentDao
		lcd = new(mysql.LogContentDaoMysql)
		logContentData, err = lcd.GetLogContentByIds(uidList, logData[0].Uid)
		if err != nil {
			var cacheData []byte
			cacheData,err = json.Marshal(logContentData)
			if err != nil {
				cache.CommonRedis.Set(redisKey, cacheData, 3600)
			}
		}


		klog.LOGGER.Info("get db success")
	}

	if err == nil {
		for k, resultInfo := range resultList {
			logId := resultInfo.LogId
			for _, logContentInfo := range logContentData {
				if logId == logContentInfo.Id {
					resultList[k].From = logContentInfo.Cfrom
					resultList[k].To = logContentInfo.Cto
				}
			}
		}
	}
	return resultList, nil
}
