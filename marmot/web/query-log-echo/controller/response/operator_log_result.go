package response

import "toy/marmot/web/query-log-echo/model/entity"

type LogResult []LogResultIn

type LogResultIn struct {
	entity.LogEntity
	From string `json:"from"`
	To string `json:"to"`
}