#ifndef __KLOG_H__
#define __KLOG_H__

#ifdef __cplusplus
extern "C" {
#endif


#include <string.h>
#include <stdio.h>
#include <stdarg.h>
#include <stdbool.h>
#include <syslog.h>
#include <unistd.h>
#include <time.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <stdlib.h>
#include <pthread.h>

//定义unsigned long长度
typedef unsigned long ulong;

//默认的切分字符串长度
#define LOG_DEFAULT_SPLITE ','
#define LOG_DEFAULT_SPLITE_LEN 5;
    
//时间的长度
#define SHOW_TIME_LEN  64;
    
//最长的路径
#define MAX_LOG_LEN 2048

#define MAX_BUF 4096
    
//正确和错误
#define KLOG_FAILD -1
#define KLOG_SUCESS 0

//日志的级别
typedef enum {
    KLOG_TRACE = 0,
    KLOG_DEBUG,
    KLOG_INFO,
    KLOG_WARING,
    KLOG_ERROR,
    KLOG_FATAL,
}klevel;

static const char *kname[] = {
    "[trace]", "[debug]", "[info]", "[warning]", "[error]", "[fatal]"
};
    
//日志存储的地方, 1 stdout  2 stderr 3 file 4 syslog
typedef enum {
    STD_OUT = 1<< 0,
    STD_ERR = 1 << 1,
    SFILE = 1 << 2,
    SSYSLOG = 1 << 3,
}show_type;

//默认按天输出日志

//初始化和关闭
int logInit(const char *str_path,ulong type);
int logClose();
    
//设置各种数据
int setLogerPath(const char *str_path);
//设置切分符号
void setSpliteData(const char *split);
    
//获取路径
const char* getLogPath();

void kshow(int level, char *file,int line,const char *fmt, ...);
//可以显示
#define klogTrace(...) kshow(KLOG_TRACE,__FILE__,__LINE__,__VA_ARGS__);
#define klogDebug(...) kshow(KLOG_DEBUG,__FILE__,__LINE__,__VA_ARGS__);
#define klogInfo(...) kshow(KLOG_INFO,__FILE__,__LINE__,__VA_ARGS__);
#define klogWarning(...) kshow(KLOG_WARING,__FILE__,__LINE__,__VA_ARGS__);
#define klogError(...) kshow(KLOG_ERROR,__FILE__,__LINE__,__VA_ARGS__);
#define klogFatal(...) kshow(KLOG_FATAL,__FILE__,__LINE__,__VA_ARGS__);



#ifdef __cplusplus
}
#endif


#ifdef __cplusplus
}
#endif

#endif