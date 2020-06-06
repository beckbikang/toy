package controller

import (
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {

	GetResultSuccess(ctx,
		gin.H{
			"msg": "pong",
		},
	)
}