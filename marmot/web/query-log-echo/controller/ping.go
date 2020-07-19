package controller

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func Ping(ctx echo.Context) error{

	type Ret struct {
		ret string
	}
	map1 := make(map[string]string)
	map1["pong"] = "ok"

	return ctx.JSON(http.StatusOK, map1)

}