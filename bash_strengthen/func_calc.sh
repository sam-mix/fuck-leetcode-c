#!/usr/bin/env bash

# 简单的命令行计算器
function calc {
    # 仅适用于整数！ --> echo The answer is: $(( $* ))
    # 浮点数
    awk "BEGIN {print \"The answer is: \" $* }";
}
