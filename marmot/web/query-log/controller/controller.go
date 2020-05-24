package controller

import (
	"github.com/gin-gonic/gin"

	"toy/marmot/web/query-log/controller/response"
	"log"
	"WEIBO/go/src/net/http"
)

func GetResultFail(c *gin.Context,err error){
	resultObj := new(response.Result)
	log.Printf("error :%v", err)
	resultObj.Code = response.Failed
	resultObj.Msg = err.Error()
	c.JSON(
		http.StatusOK,
		resultObj,
	)
}

func GetResultSuccess(c *gin.Context,ret interface{}){
	resultObj := new(response.Result)
	resultObj.Code = response.Success
	resultObj.Result = ret
	c.JSON(
		http.StatusOK,
		resultObj,
	)
}