#!/bin/bash
## 保证参数个数正确
if [ "$#" -ne 3 ];then  ##"$#" 参数个数 ##[Int1 -ne Int2] 比较Int1和Int2值，不等为真
	echo "参数错误:请使用 ./read.sh [interface]"
	echo "如./read text  test"
	exit
fi
echo "aaaaaaaaaaaa"
## 根据输入的网卡名查询IP
## 先筛选出含"inet"的行;然后以‘t’隔开，选取隔开后第二个域;再以'.'隔开，选取第1-3个域，也就是IP地址前24位;最后选取第一行(可能会有IPv6地址也被筛出)
interface=$1 ##$1 运行脚本时输入的第一个参数
interface=$2 ##$2 运行脚本时输入的第一个参数
interface=$3
##将passwd中拥有shell的用户过滤出来
cat $1 |grep $3 > $2