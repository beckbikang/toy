package dao

import 	"toy/marmot/web/query-log/model/entity"

type LogDao interface {
	GetLogData(query *entity.LogQuery) ([]entity.LogEntity,error)
}
