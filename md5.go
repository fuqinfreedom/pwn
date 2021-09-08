package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
//insert1()
	//insert2()
	//var mdfs [] string
	//mdfs = arrangement("aa","bb","cc","ddd","ee","kkkk","aaa","aaaaaaa")
	//fig := [] string{"61a60170273e74a5be90355ffe8e86ad","1dcfde5bee66097b9a735712ccac9dce", "6559326e464c2affecc7cdc0c125eb14" ,"db759cd3109d62c8b675f80edabe8235","63a1a1e27f890fec877497886e93367a"}
	//fmt.Println(fingerprint(mdfs))
	//var b string =Fingerprintmatching("1dcfde5bee66097b9a735712ccac9dce",fig)
	//fmt.Println(b)
	//var  md5 [] string
	//md5 = Fingerprintinformation("")
	//fmt.Println(md5)
	//risks("123456")

	//fmt.Println(risks("12345678910"))
	fmt.Println("======")
	//fmt.Println(risksset(1))
	//fmt.Println(delect(1))
	fmt.Println(reinforce(1))
}
//电力智能终端数据插入
func insert1() {
db, err := sql.Open("sqlite3", "./test11.db")
checkErr1(err)

//创建表 userinfo

//插入
stmt, err := db.Prepare("INSERT INTO terminalInfo(id,fingerprint,class,vendor,type,banner,ports,accounts,device,os,hardware,risk) values(?,?,?,?,?,?,?,?,?,?,?,?)")
//checkErr(err)
//createdNow :=time.Now()
stmt.Exec(nil,"12345678910", "aa", "bbb","cccc","dddd","eeee","fff","gggg","hhhh","kkk","1,2,3")
}
func checkErr1(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
//2021.8.30
//songchao
//MD5加密
//返回string
func getMD5Encode(data [] string) string {
	var i int = 0
	var a string
	for i < len(data){
		a += string(data[i])
		i++
	}
	h := md5.New()
	h.Write([]byte(a))
	return hex.EncodeToString(h.Sum(nil))
}
//
//2021.8.30
//songchao
//将字符串转换为切片
//return string

func arrangement(stringset ...string) []string {
	var i [] string
	for _,data := range stringset{
		i = append(i,data)
	}
	return i
}
//2021.8.30
//songchao
//指纹信息生成返回切片
func fingerprint(data [] string) [] string {
	var ftngerprinint [] string
	var  len = len(data)
	var i int = 0
	var termination int = 3
	for i < len {
		var finger = getMD5Encode(data[0:termination])
		if len == termination {
			break
		}
		termination += 1
		ftngerprinint = append(ftngerprinint,finger)
	}

	return ftngerprinint

}
//2021.8.30
//songchao
//扫出来的信息与数据库的指纹进行对比
//参数一扫出来的MD5  参数二数据库的MD5 返回与数据库相匹配的MD5的其它信息
func Fingerprintmatching(finger string , data [] string) string {
	var a string = finger
	var b[] string = data
	var c int = 0
	var d string
	for 0 < len(b)  {
		if a == b[c] {
			d = b[c]
			break
		}
		c ++
	}
	return d
}
//2021.9.1
//songchao
//从数据库匹配到的MD5获取其它字段的信息以切片的形式返回
//参数 匹配到MD5
func Fingerprintinformation(md5 string) [] string {
	db, err := sql.Open("sqlite3", "./test11.db")
	checkErr1(err)
	//var sqlit string = "1dcfde5bee66097b9a735712ccac9dced"
	rows, err := db.Query("SELECT * FROM terminalInfo WHERE fingerprint =?",md5)
	fmt.Println(rows)
	checkErr1(err)
	var id string
	var fing string
	var class string
	var vendor string
	var type1 string
	var banner string
	var ports string
	var accounts string
	var device string
	var os string
	var hardware string
	var risk string
	var row []string
	//	i:=0
	//var a string	//Next() 迭代查询数据
	for rows.Next() {
		//numbers[i] = append(numbers[i])
		//Scan() 读取每一行的值
		err = rows.Scan(&id, &fing, &class, &vendor, &type1, &banner, &ports, &accounts, &device, &os, &hardware, &risk)
		checkErr1(err)
		//fmt.Println("Report 数据表中所有数据信息如下：\n", id, date1, user1, ip, terminal, risks)
		row = append(row, id, fing, class, vendor, vendor, type1, banner, ports, accounts, device, os, hardware, risk)
		if (row == nil) {
			row = append(row, "null")
		}

	}
	return row
}
//2021.9.2
//songchao
//创建电力智能终端数据库
//参数匹配到的MD5值  返回风险集合
func risks(md5 string) [] string {
	db, err := sql.Open("sqlite3", "./test11.db")
	checkErr1(err)
	//var sqlit string = "1dcfde5bee66097b9a735712ccac9dced"
	rows, err := db.Query("SELECT risk FROM terminalInfo WHERE fingerprint =?",md5)
	fmt.Println(rows)
	checkErr1(err)

	var risk string
	var row []string
	//	i:=0
	//var a string	//Next() 迭代查询数据
	for rows.Next() {
		//numbers[i] = append(numbers[i])
		//Scan() 读取每一行的值
		err = rows.Scan(&risk)
		checkErr1(err)
		//fmt.Println("Report 数据表中所有数据信息如下：\n", id, date1, user1, ip, terminal, risks)
		row = append(row,risk)
		if (row == nil) {
			row = append(row, "null")
		}

	}
	return row
}
//2021.9.2
//songchao
//风险集合信息输入

func insert2() {
	db, err := sql.Open("sqlite3", "./test11.db")
	checkErr1(err)

	//创建表 userinfo

	//插入
	stmt, err := db.Prepare("INSERT INTO risks(id,name,type,vn,lever,divisor,descript,suggest,delect,reinforce) values(?,?,?,?,?,?,?,?,?,?)")
	//checkErr(err)
	//createdNow :=time.Now()
	stmt.Exec(nil,"aaa","aa", "bbb","cccc","dddd","eeee","hhhh",1,2)
}
//2021.9.2
//songchao
//根据风险集合返回风险集合的信息
//参数 risks  返回 切片
func risksset(risks int) []int {
//	var i = 0
	var risk int
	var row []int
	db, err := sql.Open("sqlite3", "./test11.db")
	checkErr1(err)

		//var sqlit string = "1dcfde5bee66097b9a735712ccac9dced"
	rows, err := db.Query("SELECT delect FROM risks WHERE  id =?",risks)
	fmt.Println(rows)
	checkErr1(err)

	//	i:=0
	//var a string	//Next() 迭代查询数据
	for rows.Next() {
		//numbers[i] = append(numbers[i])
		//Scan() 读取每一行的值
		err = rows.Scan(&risk)
		checkErr1(err)
		//fmt.Println("Report 数据表中所有数据信息如下：\n", id, date1, user1, ip, terminal, risks)
		row = append(row,risk)
		if (row == nil) {
			row = append(row,0)
		}

	}

	return row
}
//2021.9.6
//songchao
//参数 从风险集合中获取检测插件id
//返回插件地址

func delect(delect int) string {
	var risk string
	//var row string
	db, err := sql.Open("sqlite3", "./test11.db")
	checkErr1(err)

	//var sqlit string = "1dcfde5bee66097b9a735712ccac9dced"
	rows, err := db.Query("SELECT path FROM detect_plugin WHERE id =?",delect)
	fmt.Println(rows)
	checkErr1(err)

	//	i:=0
	//var a string	//Next() 迭代查询数据
	for rows.Next() {
		//numbers[i] = append(numbers[i])
		//Scan() 读取每一行的值
		err = rows.Scan(&risk)
		checkErr1(err)
		//fmt.Println("Report 数据表中所有数据信息如下：\n", id, date1, user1, ip, terminal, risks)
		}



	return risk
}
//2021.9.6
//songchao
//参数 从风险集合中获取加固插件id
//返回插件地址
func reinforce(reinforce int) string {
	var riska string
	//var row string
	db, err := sql.Open("sqlite3", "./test11.db")
	checkErr1(err)

	//var sqlit string = "1dcfde5bee66097b9a735712ccac9dced"
	rowsa, err := db.Query("SELECT path FROM modify_plugin WHERE id =?",reinforce)
	fmt.Println(rowsa)
	checkErr1(err)

	//	i:=0
	//var a string	//Next() 迭代查询数据
	for rowsa.Next() {
		//numbers[i] = append(numbers[i])
		//Scan() 读取每一行的值
		err = rowsa.Scan(&riska)
		checkErr1(err)
		//fmt.Println("Report 数据表中所有数据信息如下：\n", id, date1, user1, ip, terminal, risks)
	}
	return riska
}
