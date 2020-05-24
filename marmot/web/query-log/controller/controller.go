package controller

import (
	"github.com/gin-gonic/gin"

	"toy/marmot/web/query-log/controller/response"
	"log"
)

func GetResultFail(c *gin.Context,err error){
	resultObj := new(response.Result)
	log.Printf("error :%v", err)
	resultObj.Code = response.Failed
	resultObj.Msg = err.Error()
	c.JSON(
		200,
		resultObj,
	)
}

func GetResultSuccess(c *gin.Context,ret interface{}){
	resultObj := new(response.Result)
	resultObj.Code = response.Success
	resultObj.Result = ret
	c.JSON(
		200,
		resultObj,
	)
}