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
	Uid         int64 `validate:"required"`
	LogType     int   `validate:"gte=0"`
	LogTargetId int   `validate:"gte=0"`
	StartTime   string
	EndTime     string
	Page        int `validate:"required,number"`
	PageSize    int `validate:"required,number"`
}
