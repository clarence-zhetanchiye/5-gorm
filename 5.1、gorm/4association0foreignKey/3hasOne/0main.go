package main

import material "guanlian/0material"
//todo：-------------------------------------------采用gorm默认外键的方式---------------------------------------------------

//has one 关系是指，“主”有且仅有一个“从”，是一一对应关系。即一个购物车有且仅有一个购物项

//“主” 【相当于购物车】
type User2 struct {
	ID int			//todo:gorm将ID默认为primary_key和Auto_increment和references
	Name string
	Card CreditCard2

}
//“从”  【相当于购物项】
type CreditCard2 struct{
	ID int      //todo:gorm将ID默认为primary_key和Auto_increment
	Code string
	Holder string
	User2ID int  //todo:如不自定义外键，这被gorm默认为外键
}

func main() {
	material.GetDb()
	material.MyDb.AutoMigrate(&User2{},&CreditCard2{})

	one2 := User2{
		Name:"jack",
		Card:CreditCard2{Code: "111"},
	}
	one22 := User2{
		Name:"jack",
		Card:CreditCard2{Code: "111"},
	}
	material.MyDb.Create(&one2)
	material.MyDb.Create(&one22)

}
/*
user2表的情况
+----+------+
| id | name |
+----+------+
|  1 | jack |
|  2 | jack |
+----+------+
credit_card2表的情况
+----+------+--------+----------+
| id | code | holder | user2_id |
+----+------+--------+----------+
|  1 | 111  |        |        1 |
|  2 | 111  |        |        2 |
+----+------+--------+----------+
*/
