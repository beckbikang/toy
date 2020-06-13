CC=gcc
BASE_DIR= .
BASE_INC= ${BASE_DIR}/include

INC= -I${BASE_INC}
LIBDIR= 
LIB= 
COMPILE_DIR=compile
#FLAG= ${LIBDIR} ${LIB} ${INC} -O3 -Wall -W -ansi
#FLAG= ${LIBDIR} ${LIB} ${INC} -O3 -ansi -pedantic
FLAG= ${LIBDIR} ${LIB} ${INC} -O3 -std=gnu89 -Wall -W -DNDEBUG

#makefile里的函数使用，和取变量的值类似，是以一个‘$’开始，
#然后是一个括号里面是函数名和需要的参数列表，多个变量用逗号隔开，像这样
#return = $(functionname  arg1,arg2,arg3...)。
#可能这里的'$'更像是从某个地址取值类似的操作

#notdir
#$(wildcard PATTERN...)
#使用“$(wildcard *.c)”来获取工作目录下的所有的.c文件列表
#SRC = $(wildcard *.c ./foo/*.c) 
#使用：SRC = $(notdir wildcard)
#去除所有的目录信息，SRC里的文件名列表将只有文件名。

#使用：OBJ = $(patsubst %.c %.o $(SRC)) 
#patsubst是patten substitude的缩写，匹配替代的意思。
#这句是在SRC中找到所有.c 结尾的文件，然后把所有的.c换成.o。

TARGET=main
SRCS=$(wildcard src/*.c)
SRCSNOTDIR=$(notdir $(SRCS))
OBJS=$(patsubst %.c, compile/%.o, $(SRCSNOTDIR)) 

.PHONY: clean 

#$(shell if [ -f $(COMPILE_DIR) ]; then echo "Existed!"; else echo "NO" ;  fi;)

$(TARGET):$(OBJS)
	${CC} -o $@ $^ ${FLAG} 

compile/%.o:src/%.c
	${CC} -o $@ -c $< ${FLAG} 

clean :
	rm -rf $(TARGET) $(OBJS) 


