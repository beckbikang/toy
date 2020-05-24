package entity

/**
CREATE TABLE `log_content_1` (
  `id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '日志ID',
  `from` text NOT NULL COMMENT '修改前',
  `to` text NOT NULL COMMENT '修改后',
  `mtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='日志修改内容表';
**/
type LogContentEntity struct {
	Id    int    `ddb:"id"`
	From  uint64 `ddb:"from"`
	To    int    `ddb:"to"`
	Mtime string `ddb:"mtime"`
}
