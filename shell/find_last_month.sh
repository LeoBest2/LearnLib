#!/bin/bash


# 查找上个月月初至月末文件，如当前日期: 2022-02-10 则查找 2022-01-01 ~ 2022-01-31 之间的文件
find . -type f -newermt "$(date -d "$(date +%Y-%m-1) - 1 month" +"%Y-%m-%d")" -not -newermt "$(date +%Y-%m-1)"

# 查找距离今天一个月内文件，如当前日期: 2022-02-10 则查找 2022-01-10 ~ 2022-02-10 之间的文件
find . -type f -newermt $(date -d "$(date +%Y-%m-%d) - 1 month" +%Y-%m-%d)

# 查找本月月初至今天内文件, 如当前日期: 2022-02-10 则查找 2022-02-01 ~ 2022-02-10 之间的文件
find . -type f -newermt $(date +%Y-%m-1)