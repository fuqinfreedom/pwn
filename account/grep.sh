#!/bin/bash
## 保证参数个数正确
if [ "$#" -ne 2 ];then  ##"$#" 参数个数 ##[Int1 -ne Int2] 比较Int1和Int2值，不等为真
	echo "参数错误:请使用 ./grep.sh text text"
	echo "如 text text "
	exit
fi
echo "aaaaaaaaaaaa"
interface=$1 ##$1 运行脚本时输入的第一个参数
interface=$2
##
cat $1 | awk -F: '$3>=0' | cut -f 1 -d :  > $2