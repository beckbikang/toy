package response

import "toy/marmot/web/query-log-echo/model/entity"

type LogResult []LogResultIn

type LogResultIn struct {
	Log entity.LogEntity
	LogContent entity.LogContentEntity
}