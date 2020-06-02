package controller

import (
	"strconv"
	"toy/marmot/web/query-log/biz/service"
	"toy/marmot/web/query-log/model/entity"
	"github.com/gin-gonic/gin"
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
