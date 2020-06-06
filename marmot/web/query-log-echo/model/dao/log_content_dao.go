package dao

import (
	"toy/marmot/web/query-log-echo/model/entity"
)

//deal with log content
type LogContentDao interface {
	GetLogContent(id ,uid int64) ( *entity.LogContentEntity,error)
	GetLogContentByIds(ids []int64, uid int64) ( []entity.LogContentEntity,error)
	Save(id uint64, from, to string, uid int64) bool
	SaveEntity(lce *entity.LogContentEntity, uid int64) bool
}
