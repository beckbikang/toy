#include "log.h"


//定义基本的结构
typedef struct Klog{
    const char* logpath;//存储地址
    char *date;
    FILE * fp;
    int fd;
    ulong flag;//日志发送的地方
    char *splite_char;//分隔符
    klevel l;
    pthread_mutex_t mutex;//锁
}KLogger;


//初始化一个全局日志
static KLogger KL = {"","", NULL, -1,STD_OUT,"",KLOG_DEBUG};

//根据条件初始化日志
KLogger* logInit(const char *str_path,ulong flag){
    
    KL.flag = flag;
    KL.logpath = str_path;
    if((flag & PUT_SFILE) == PUT_SFILE){
        KL.fd = open(str_path, O_WRONLY|O_CREAT|O_APPEND,S_IRWXU|S_IRWXO);
        if(KL.fd < 0){
            return KLOG_FAILD;
        }
    }
    if ((KL.flag & PUT_SYSLOG) == PUT_SYSLOG) {
        //记录pid 立即打开连接， 记录用户日志
        openlog(KL.logpath, LOG_PID | LOG_NDELAY, LOG_USER);
    }
    int len = LOG_DEFAULT_SPLITE_LEN;
    memset(KL.splite_char, '\0', len);
    KL.splite_char[0] = LOG_DEFAULT_SPLITE;
    KL.is_show_line = 0;
    return KLOG_SUCESS;
}


//重新设置路径
int setLogerPath(const char *str_path){
    logClose();
    if (str_path && strlen(str_path) > 0){
        return logInit(str_path,KL.flag);
    }
    return KLOG_FAILD;
}


//设置切分符号
void setSpliteData(const char *split){
    int len = LOG_DEFAULT_SPLITE_LEN;
    if(strlen(split)> 0 && strlen(split) < len){
        strcpy(KL.splite_char, split);
    }
}


int logClose(){
    
    if(KL.fd != -1 && close(KL.fd) < 0){
        //todo  add something
        return KLOG_FAILD;
    }
    
    if ((KL.flag & PUT_SYSLOG) == PUT_SYSLOG) {
        closelog();
    }
    KL.fd = -1;
    
    return KLOG_SUCESS;
}

//获取当前时间
char *getCurrentTime(){
    struct tm *ptr;
    int len = SHOW_TIME_LEN;
    time_t tl;
    char *buf = (char*)malloc(sizeof(char) * len);
    
    tl = time(NULL);
    ptr = localtime(&tl);
    buf[strftime(buf, len , "%Y-%m-%d %H:%M:%S", ptr)]='\0';
    return buf;
}

/**
处理日志:
    1 输出时间，文件名，行数，你需要的信息
    2 注意各个数据的切分
    3 注意数据的长度
 
 **/
void kshow(int level, char *file,int line,const char *fmt, ...){
    char str_format[MAX_LOG_LEN];
    char buffer[MAX_LOG_LEN];
    va_list args;
    
    memset(str_format, '\0', MAX_LOG_LEN);
    memset(buffer, '\0', MAX_LOG_LEN);
    char *atime = getCurrentTime();
    
    //获取当前的需要格式化的字符串
    if (KL.is_show_line == 1){
        snprintf(str_format, MAX_LOG_LEN,"%s%s%s:%d%s%s%s%s",
        KLONG_NAME[level],KL.splite_char,file,line,KL.splite_char,
                 atime,KL.splite_char,fmt);
    }else{
        snprintf(str_format, MAX_LOG_LEN,"%s%s%s%s%s\n",
                 KLONG_NAME[level],KL.splite_char,
                 atime,KL.splite_char,fmt);
    }
    //获取需要拼接的字符串
    
    va_start(args,fmt);
    size_t buf_len = vsnprintf(buffer,MAX_LOG_LEN, str_format,args);
    va_end(args);
    //拼接出需要的字符串，然后根据flag输出到对应的fd里面
    
    if ((KL.flag & PUT_STD_OUT) == PUT_STD_OUT) {
        printf("%s",buffer);
        fflush(stdout);
    }
    if ((KL.flag & PUT_STD_ERR) == PUT_STD_ERR) {
        printf("%s",buffer);
        fflush(stderr);
    }
    
    
    if ((KL.flag & PUT_SYSLOG) == PUT_SYSLOG) {
        syslog(level, buffer,"");
    }
    
    if ((KL.flag & PUT_SFILE) == PUT_SFILE) {
        write(KL.fd, buffer, buf_len);
    }
    free(atime);
}

