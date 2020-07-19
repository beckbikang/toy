package controller

import (
	"strconv"
	"toy/marmot/web/query-log-echo/biz/service"
	"toy/marmot/web/query-log-echo/model/entity"
	"github.com/labstack/echo/v4"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zht "github.com/go-playground/validator/v10/translations/zh"
	"fmt"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans ut.Translator
)

func init(){
	z := zh.New()
	uni = ut.New(z,z)
	trans, _ = uni.GetTranslator("zh")
	validate = validator.New()
	zht.RegisterDefaultTranslations(validate, trans)
}


func getQueryStr(c echo.Context) (*entity.LogQuery, error) {
	//uid
	uidStr := c.QueryParam("uid")
	var err error
	var uid int64 = 0
	if uidStr != ""{
		uid, err = strconv.ParseInt(uidStr, 10, 64)
		if err != nil {
			GetResultFail(c, err)
			return nil, err
		}
	}

	//log type
	logTypeStr := c.QueryParam("log_type")
	var logType = 0
	if logTypeStr != ""{
		logType, err = strconv.Atoi(logTypeStr)
		if err != nil {
			GetResultFail(c, err)
			return nil, err
		}
	}


	//log target id
	logTargetIdStr := c.QueryParam("log_target_id")
	var logTargetId = 0
	if logTargetIdStr != ""{
		logTargetId, err = strconv.Atoi(logTargetIdStr)
		if err != nil {
			GetResultFail(c, err)
			return nil, err
		}
	}


	startTime := c.QueryParam("start_time")
	endTime := c.QueryParam("end_time")

	//page
	pageStr := c.QueryParam("page")
	var page = 1
	if  pageStr != ""{
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			GetResultFail(c, err)
			return nil, err
		}
	}


	//page size
	pageSizeStr := c.QueryParam("page_size")
	var pageSize = 20
	if pageSizeStr != ""{
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil {
			GetResultFail(c, err)
			return nil, err
		}
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


	//tans
	//validator
	err = validate.Struct(query)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		GetResultFail(c, fmt.Errorf("%s", errs.Translate(trans)))
		return nil, err
	}

	return query, nil
}

//获取日志列表
func OpLogList(c echo.Context) error{

	query, err := getQueryStr(c)
	if err != nil {
		return err
	}

	ret, err := service.GetLogData(query)

	if err != nil {
		GetResultFail(c, err)
	} else {
		GetResultSuccess(c, ret)
	}
	return nil
}
