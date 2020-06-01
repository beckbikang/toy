# 写一个简单的基于gin的日志查询服务

## 项目大致的流程

通过一个很小的功能实现整个业务体系

step1 基础功能

    配置更新fsnotify/fsnotify
    配置管理spf13/viper
    restfull接口
    使用gin框架
    数据库查询 didi/gendry， 实现简单的查询，不期望通过orm，获取结果
    使用日志库记录日志:zapper和lumberjack写入日志

step2 中间件

    支持参数检测
    使用redis的cache
    完善代码结构
    swagger
    基于net/http/httptest 的单元测试
    
step3 
    
    接入用户认证、跨域、访问日志、请求频率限制、追踪 ID
    依赖注入
    JWT认证
    基于go模板支持html页面
    增加gorm的支持

step4

    支持grpc，rpcx，各种协议的服务调用
    go的内存分配，cpu使用检测，垃圾回收

step5

    抽象接入服务发现和注册层
    接入全链路追踪
    

step6 

    接入监控体系
    vue或者reactor完善前面的界面


    

## 项目详情

### sql 

```
### 日志的查询

提供一个简单的日志服务，根据uid和log_type,log_target_id提供服务


数据表的格式

```sql
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

CREATE TABLE `log_2` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `uid` bigint(20) unsigned NOT NULL COMMENT 'uid',
  `log_type` tinyint(5) unsigned NOT NULL DEFAULT '0' COMMENT '操作类型',
  `log_target_id` int(11) DEFAULT NULL COMMENT '操作对象id',
  `log_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '日志ID',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_uid_tid` (`uid`,`log_target_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='日志表';

CREATE TABLE `log_content_1` (
  `id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '日志ID',
  `cfrom` text NOT NULL COMMENT '修改前',
  `cto` text NOT NULL COMMENT '修改后',
  `mtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='日志修改内容表';


CREATE TABLE `log_content_2` (
  `id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '日志ID',
  `cfrom` text NOT NULL COMMENT '修改前',
  `cto` text NOT NULL COMMENT '修改后',
  `mtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='日志修改内容表' ;

```
