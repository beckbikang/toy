package entity

/**
CREATE TABLE `log_1` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `uid` bigint(20) unsigned NOT NULL COMMENT 'uid',
  `log_type` tinyint(5) unsigned NOT NULL DEFAULT '0' COMMENT '操作类型',
  `log_target_id` int(11) DEFAULT NULL COMMENT '操作对象id',
  `log_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '日志ID',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_uid_tid` (`uid`,`log_target_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='日志表';
**/
type LogEntity struct {
	Id          int    `ddb:"id"`
	Uid         int64  `ddb:"uid"`
	LogType     int    `ddb:"log_type"`
	LogTargetId int    `ddb:"log_target_id"`
	LogId       int64  `ddb:"log_id"`
	Mtime       string `ddb:"mtime"`
}
