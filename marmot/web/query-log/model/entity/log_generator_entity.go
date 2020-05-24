package entity

/**
CREATE TABLE `log_id_generator` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `gtype` int(11) NOT NULL COMMENT '生成器类型 1 日志',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
**/
type LogIdGenEntity struct {
	Id    int    `ddb:"id"`
	GType int    `ddb:"gtype"`
	Mtime string `ddb:"mtime"`
}
