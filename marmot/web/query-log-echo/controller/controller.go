package controller

import (
	"github.com/labstack/echo/v4"

	"log"
	"net/http"
	"toy/marmot/web/query-log-echo/controller/response"
)

func GetResultFail(c echo.Context, err error) {
	resultObj := new(response.Result)
	log.Printf("error :%v", err)
	resultObj.Code = response.Failed
	resultObj.Msg = err.Error()

	c.JSON(
		http.StatusOK,
		resultObj,
	)
}

func GetResultSuccess(c echo.Context, ret interface{}) {
	resultObj := new(response.Result)
	resultObj.Code = response.Success
	resultObj.Result = ret
	c.JSON(
		http.StatusOK,
		resultObj,
	)
}
