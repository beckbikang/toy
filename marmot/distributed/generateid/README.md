# id生成器常用的模式

1 数据库自增ID

2 数据库多主模式：多个master生成id

号段模式:从数据库里面取出数据，放置到内存里面
```
CREATE TABLE id_generator (
  id int(10) NOT NULL,
  max_id bigint(20) NOT NULL COMMENT '当前最大id',
  step int(20) NOT NULL COMMENT '号段的布长',
  biz_type	int(20) NOT NULL COMMENT '业务类型',
  version int(20) NOT NULL COMMENT '版本号',
  PRIMARY KEY (`id`)
) 
```

Redis

雪花算法SnowFlake