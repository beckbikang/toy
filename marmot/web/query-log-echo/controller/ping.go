package controller

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func Ping(ctx echo.Context) error{


	return ctx.JSON(http.StatusOK, "{pong}")

}