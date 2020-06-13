#include<stdio.h>

#include "log.h"

int main(){
    printf("start testing log\n");
    printf("test it\n");
    int ret;
    
    //打开日志
    ret = logInit("/data0/collect/logs/klog/t.log",PUT_STD_OUT | PUT_SFILE);//| PUT_SYSLOG);
    if (ret < 0) {
        puts("\nopen log error\n");
    }else{
        puts("open log sucess\n");
    }
    //打印当前时间
    char *ptime;
    ptime = getCurrentTime();
    printf("%s;len=%ld\n",ptime,strlen(ptime));
    
    //打印级别
    printf("kl.flag=%x\n",(int)getLogerFlag());
    
    //写日志
    klogTrace("%s","atest");
    klogDebug("%s","atest");
    klogInfo("%s","atest");
    klogWarning("%s","atest");
    klogError("%s","atest");
    klogFatal("%s","atest");
    
    //写日志
    setShowFile(1);
    setSpliteData("||");
    klogTrace("%s","atest");
    
    
    ret = logClose();
    if (ret < 0) {
        puts("close log error\n");
    }else{
        puts("close log sucess\n");
    }
}