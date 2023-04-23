package main

import (
	"fmt"
	material "go0ji1chu3/22gorm/2foreignKey0guan1lian2/0material"
)

type Menber struct{
	ID int
	Name string
	Belong int
}
type Class struct {
	ID int
	Name string
	Menbers []Menber `gorm:"foreignKey:Belong"`//不用写references:ID就会默认Class的ID是被参考的
}
func main() {
	material.GetDb()
	material.MyDb.AutoMigrate(&Class{},&Menber{})

	c := Class{Name:"class1",Menbers: []Menber{{Name:"jack"},Menber{Name:"tom"}}}
	material.MyDb.Create(&c)

	var m []Menber
	//todo:Association意为选取变量c下Menbers字段对应数据库表，Find依然是将查询结果放入m的意思
	// Model中的&c不能是&Class{}，否则m=[]
	material.MyDb.Model(&c).Association("Menbers").Find(&m)
	fmt.Println("m=",m)//m= [{1 jack 1} {2 tom 1}]

	//todo:where必须在Association前面，where里的内容和之前查询方式一样可以有很多
	material.MyDb.Model(&c).Where("name=?","jack").Association("Menbers").Find(&m)
	fmt.Println("m2=",m)//m2= [{1 jack 1}]

	//todo:material.MyDb.Model(&c).Association()  的返回值是一个结构体，其下有
	//   Find,Append,Replace,Delete,Clear,Count六个用于增查改删和计数的方法。
}
/*
mysql> select * from classes;
+----+--------+
| id | name   |
+----+--------+
|  1 | class1 |
+----+--------+
mysql> select * from menbers;
+----+------+--------+
| id | name | belong |
+----+------+--------+
|  1 | jack |      1 |
|  2 | tom  |      1 |
+----+------+--------+
*/
