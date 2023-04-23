package main

import (
	material "crud/0material"
	"fmt"
)

//todo:gorm语句指明接收查询结果的结构体变量时，该变量的结构体类型名就已经对应了从哪个数据库表查询，区别是蛇形复数规则；该结构体变量的字段对应
// 着相应数据表的列，区别也是蛇形复数规则。若接收查询结果的该结构体变量的类型名与数据库表名蛇形不一致,则可通过.Table()指定数据库表；若接收查
// 询结果的该结构体变量的字段与相应数据表的列不对应，则可以通过在该字段后面写上gorm的column标签来让字段和数据库表相应的column对应。若接收
// 查询结果的该结构体变量的字段数量与相应数据表的列的数量不一样多，也没有关系，只会自动不接收相应-------
// 一般地，gorm语句中都建议用上.Table()来显示地指明是针对数据库的哪个表。
// Find()的第一个入参只建议是&struct或&[]struct; First()、Last()、Take()由于只会查到一条记录，故第一个入参只建议是&struct。
// --------------------------------------------------------------------------------------------------------------------
var createTable = `CREATE TABLE IF NOT EXISTS teachers (
name varchar(64),
age bigint(20) DEFAULT NULL,
frequent bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1`
var createData = "INSERT INTO `teachers` VALUES ('laoWu', 0, 2),('laowang', 0, 2),('laoliu', 35, 3),('laoLi', 66, 7);"

func read0() {
	material.MyDb.Exec("DROP TABLE IF EXISTS `teachers`;")
	material.MyDb.Exec(createTable)
	material.MyDb.Exec(createData)


	//todo:若接收查询结果的自定义结构体，它的类型名与数据库表名蛇形不一致,则可通过.Table()指定数据库表; 同时该自定义结构体内的字段
	// 名需与数据库表中相应的column蛇形一致，不一致时可借助gorm的column标签来让字段和数据库表相应的column对应。
	type Teach struct { //类型名与数据表名称不是蛇形对应的时候
		Name string
		Age  int
		Freq string `gorm:"column:frequent"` //通过gorm的column标签来和数据表的字段对应。
	}
	var t1 Teach
	//下一行的material.MyDb已指明到哪个数据库查，First(&t1)里t1的类型则指明了去数据库的哪个表查，但当.Table()指明表时则听.Table()的。
	material.MyDb.Table("teachers").First(&t1) //SELECT * FROM `teachers` ORDER BY `teachers`.`name` LIMIT 1
	fmt.Println("t1=", t1) //t1= {laoLi 66 7}

	type teacher struct { //todo:少了一个字段 Age 也没关系，查仍会SELECT *，但只是不会接收到相应的字段的值。
		Name     string
		Frequent string
	}
	var t2 teacher
	//下一行的material.MyDb通过(&t2)中t2的类型表明了到对应数据库下的teachers表查询
	material.MyDb.First(&t2)//SELECT * FROM `teachers` ORDER BY `teachers`.`name` LIMIT 1
	fmt.Println("t2=", t2) //t2= {laoLi 7}


	//todo:Find()的入参只建议是&struct或&[]struct
	var t3 teacher
	material.MyDb.Table("teachers").Where("name=?", "laoWu").Find(&t3)
	//SELECT * FROM `teachers` WHERE name='laoWu'
	fmt.Println("t3=", t3)//t3= {laoWu 2}

	var t4 []teacher
	material.MyDb.Table("teachers").Where("age=?", 0).Find(&t4)//SELECT * FROM `teachers` WHERE age=0
	fmt.Println("t4=", t4)//t5= [{laoWu 2} {laowang 2}]

	var t5 teacher //发现虽然查询结果有多个，但这样也不会报错，不够只会接收到若干条中的一条。
	material.MyDb.Table("teachers").Where("age=?", 0).Find(&t5)//SELECT * FROM `teachers` WHERE age=0
	fmt.Println("t5=", t5)//t4= {laoWu 2}
}

func main() {
	material.GetDB() //初始化。同时也就顺带指定了是针对Mysql的哪个数据库
	//gorm语句的入门使用和关键的两点。
	read0()
}
