package controller

import (
	"strconv"
	"toy/marmot/web/query-log/biz/service"
	"toy/marmot/web/query-log/model/entity"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func getQueryStr(c *gin.Context) (*entity.LogQuery, error) {
	//uid
	uidStr := c.DefaultQuery("uid", "0")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		GetResultFail(c, err)
		return nil, err
	}

	//log type
	logTypeStr := c.DefaultQuery("log_type", "0")
	logType, err := strconv.Atoi(logTypeStr)
	if err != nil {
		GetResultFail(c, err)
		return nil, err
	}

	//log target id
	logTargetIdStr := c.DefaultQuery("log_target_id", "0")
	logTargetId, err := strconv.Atoi(logTargetIdStr)
	if err != nil {
		GetResultFail(c, err)
		return nil, err
	}

	startTime := c.DefaultQuery("start_time", "")
	endTime := c.DefaultQuery("end_time", "")

	//page
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		GetResultFail(c, err)
		return nil, err
	}

	//page size
	pageSizeStr := c.DefaultQuery("page_size", "20")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		GetResultFail(c, err)
		return nil, err
	}

	//build query object
	query := new(entity.LogQuery)
	query.Uid = uid
	query.LogType = logType
	query.LogTargetId = logTargetId
	query.StartTime = startTime
	query.EndTime = endTime
	query.Page = page
	query.PageSize = pageSize

	//validator
	vald := validator.New()
	err = vald.Struct(query)
	if err != nil {
		GetResultFail(c, err)
		return nil, err
	}

	return query, nil
}

//获取日志列表
func OpLogList(c *gin.Context) {

	query, err := getQueryStr(c)
	if err != nil {
		return
	}

	ret, err := service.GetLogData(query)

	if err != nil {
		GetResultFail(c, err)
	} else {
		GetResultSuccess(c, ret)
	}
}
