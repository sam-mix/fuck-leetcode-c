CC = gcc
CPPFLAGS = -Wall

FILE_SUFFIX = c

# 所有源文件
SRCS := $(wildcard *.$(FILE_SUFFIX))
# 所有可执行文件
EXES := $(patsubst %.$(FILE_SUFFIX),%,$(SRCS))

# 以输入文件名称生成所有输出文件并执行后清除
all:$(EXES)
$(EXES): % : %.$(FILE_SUFFIX)
	@$(CC) $(CPPFLAGS) $< -o $@	&& echo $< && ./$@ && rm -rf $@ && echo "\n"

.PHONY: all

clean:
	@rm -rf $(EXES)
	
.PHONY: clean