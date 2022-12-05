CC = gcc
CFLAGS = -Wall
FILE_SUFFIX = c

CPP = g++
CPPFLAGS = -Wall --std=c++20
CPP_FILE_SUFFIX = cpp


# 所有源文件
SRCS := $(wildcard *.$(FILE_SUFFIX))
# 所有可执行文件
EXES := $(patsubst %.$(FILE_SUFFIX),%,$(SRCS))

# 所有源文件
CPPSRCS := $(wildcard *.$(CPP_FILE_SUFFIX))
# 所有可执行文件
CPPEXES := $(patsubst %.$(CPP_FILE_SUFFIX),%,$(CPPSRCS))

# 以输入文件名称生成所有输出文件并执行后清除
all:$(EXES)
$(EXES): % : %.$(FILE_SUFFIX)
	@$(CC) $(CFLAGS) $< -o $@	&& echo $< && ./$@ && rm -rf $@ && echo "\n"

.PHONY: all

cpp:$(CPPEXES)
$(CPPEXES): % : %.$(CPP_FILE_SUFFIX)
	@$(CPP) $(CPPFLAGS) $< -o $@	&& echo $< && ./$@ && rm -rf $@ && echo "\n"

.PHONY: cpp

clean:
	@rm -rf $(EXES)
	@rm -rf $(CPPEXES)

.PHONY: clean