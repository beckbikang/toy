package entity

type LogEntity struct {
	Id          int    `ddb:"id"`
	Uid         int64  `ddb:"uid"`
	LogType     int    `ddb:"log_type"`
	LogTargetId int    `ddb:"log_target_id"`
	LogId       int64  `ddb:"log_id"`
	Mtime       string `ddb:"mtime"`
}

//log query
type LogQuery struct {
	Uid int64
	LogType int
	LogTargetId int
	StartTime string
	EndTime string
	Page int
	PageSize int
}