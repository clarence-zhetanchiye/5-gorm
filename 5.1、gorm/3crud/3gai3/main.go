//todo:不用看官网上其他的方式，那些方式都有不少要注意的地方，以下提供的方式扯着用，不会有啥问题和要注意的地方。
// 另外
// 1、GORM 也允许使用 SQL 表达式、自定义数据类型的 Context Valuer 来更新，可见GORM官网/CRUD接口/更新。
// 2、MySQL不支持GORM官网/CRUD接口/更新/返回修改行的数据
// 3、GORM官网/CRUD接口/检查字段是否有变更？、在Update时修改值，他的意思并不是说gorm语句是否成功地更改了
//   数据库表中的值，只是判定gorm语句中Update、Updates中的和gorm语句里Model的相应字段的值是否相同，见其官方举例即可。所以用途不大。

package main

import (
	material "crud/0material"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func main() {
	material.GetDB()
	insertGoods()
	insertUser()

	//todo:如果想看每个gorm语句的效果，可以断点然后Debug逐步地看。

	//更改数据表中的一个字段
	updateOne() //todo:想gorm语句中不写Where()条件来更新所有行会报错。除非在初始获取gormDB时gorm.Config{}的AllowGlobalUpdate设为true

	//更改数据表中的多个字段
	updatesMany() //todo:想gorm语句中不写Where()条件来更新所有行会报错。除非在初始获取gormDB时gorm.Config{}的AllowGlobalUpdate设为true

	//更新所有行的一到多column的值，只能用本方法内介绍的三类方式！！！！！！
	updateAllRows()

	save()
}
//todo:改数据库表中一到多行里某一个column的值--------------------------------------------------------------------------------
func updateOne() {
	//更新一行的某个column的值------------------
	material.MyDb.Model(&good{}).Where("id=? AND name=?",1,"apple").Update("price",9999)
	//UPDATE `goods` SET `price`=9999 WHERE id=1 AND name='apple'

	//更新多行的同一个column--------------------
	material.MyDb.Table("goods").Where("name=?","banana").Update("amount",0)
	//UPDATE `goods` SET `amount`=0 WHERE name='banana'

	material.MyDb.Model(&good{}).Where("id in ?",[]int{5, 6}).Update("amount",1)
	//UPDATE `goods` SET `amount`=1 WHERE id in (5,6)

	//在3crud/2cha2/cha中创建过users表且插入过数据，这里就直接申明一个该类型结构体供Model()使用。
	//字段不完整甚至无字段都没关系，结构体类型名叫user就行，就能找到material.MyDb指向的数据库下的users表。
	type user struct {}
	material.MyDb.Model(&user{}).Where("id=?",9).Update("name", "stranger")
	//UPDATE `users` SET `name`='stranger' WHERE id=9

	//todo：进行查询的gorm语句中的各种Where()条件都可以用在上面。


	//todo:使用gorm.Expr表达式，------------------------------
	// 在数据表某值的原数据基础上修改！！！
	material.MyDb.Table("goods").Where("id=?", 1).
		Update("price", gorm.Expr("price * ? + ?", 100, 9))//切记这里Update的第一个参数不得有等号问号。
	//UPDATE `goods` SET `price`=price * 100 + 9 WHERE id=1


	//todo:根据子查询进行更新----------------------------------
	material.MyDb.Table("goods").Where("id=?", 1).
		Update("price", material.MyDb.Table("users AS u").
			Select("age").Where("u.id=goods.id AND u.name=?", "jack"))
	//UPDATE `goods` SET `price`=(SELECT age FROM users AS u WHERE u.id=goods.id AND u.name='jack') WHERE id=1
}

//todo:更改数据库表中一到多行里多个column的值--------------------------------------------------------------------------------
func updatesMany(){
	//更改一行的多个column的值----------------------------------
	material.MyDb.Model(&good{}).Where("id=?",1).Updates(map[string]interface{}{"name":"juice","price":9})
	//UPDATE `goods` SET `name`='juice',`price`=9 WHERE id=1

	//更改多行的多个column的值---------------------------------
	material.MyDb.Table("goods").Where("name=?", "banana").
		Updates(map[string]interface{}{"name":"b", "price":6})
	//UPDATE `goods` SET `name`='b',`price`=6 WHERE name='banana'

	//todo:使用gorm.Expr表达式，------------------------------
	// 在数据表某值的原数据基础上修改！！！
	material.MyDb.Table("goods").Where("name=?", "b").
		Updates(map[string]interface{}{
			"name": gorm.Expr("`name` + ?", "xx"), //todo:MySQL中不存在字符串相加，故这样写得出了name列变成"0"的现象。
			"price":gorm.Expr("price * ?", 10),
		})
	//UPDATE `goods` SET `name`=`name` + 'xx',`price`=price * 10 WHERE name='banana'

	//todo:根据子查询进行更新----------------------------------
	material.MyDb.Table("goods").Where("id=?", 1).
		Updates(map[string]interface{}{
			"price": material.MyDb.Table("users").Select("age").Where("users.id=goods.id"),
			"amount": material.MyDb.Table("users").Select("age").Where("users.id=goods.id"),
	})
	//UPDATE `goods` SET `amount`=(SELECT age FROM `users` WHERE users.id=goods.id),
	//                    `price`=(SELECT age FROM `users` WHERE users.id=goods.id) WHERE id=1
}

//todo:一次更改全部行的一或多column字段的值----------------------------------------------------------------------------------
func updateAllRows(){
	//todo:方式一，采用原始SQL
	material.MyDb.Exec("UPDATE goods SET price=?",2022)//将数据库表里所有行的price更改为2022
	//UPDATE goods SET price=2022

	material.MyDb.Exec("UPDATE goods SET price=? , amount=?",0, 0)	//todo:注意SET的两个字段之间是逗号而不是AND
	//UPDATE goods SET price=0 , amount=0


	//todo:方式二，采用带有Where("1=1")的gorm语句
	material.MyDb.Table("goods").Where("1=1").Update("amount",1000)
	//UPDATE `goods` SET `amount`=1000 WHERE 1=1

	material.MyDb.Table("goods").Where("1=1").Updates(map[string]interface{}{"price":0, "amount":0})
	//UPDATE `goods` SET `amount`=0,`price`=0 WHERE 1=1


	//todo:方式三，在Session中设定gorm.Session{AllowGlobalUpdate: true}
	material.MyDb.Session(&gorm.Session{AllowGlobalUpdate: true}).Table("goods").Update("amount", 0)
	//UPDATE `goods` SET `amount`=0
}

type good struct{
	Id int
	Name string
	Price int
	Amount int
	Code int `gorm:"unique"`
}
var insertGoodsSql = `
INSERT INTO goods VALUES (1, 'apple', 5, 100, 111);
INSERT INTO goods VALUES (2, 'apple', 8, 200, 222);
INSERT INTO goods VALUES (3, 'banana', 3, 300, 333);
INSERT INTO goods VALUES (4, 'banana', 6, 400, 444);
INSERT INTO goods VALUES (5, 'orange', 9, 500, 555);
INSERT INTO goods VALUES (6, 'peach', 12, 600, 666)
`
func insertGoods(){
	material.MyDb.Exec("DROP TABLE IF EXISTS `goods`;")
	material.MyDb.AutoMigrate(&good{})
	for _, v := range strings.Split(insertGoodsSql, ";") {
		material.MyDb.Exec(v)
	}
}

func insertUser() {
	material.MyDb.Exec("DROP TABLE IF EXISTS users")
	if err := material.MyDb.AutoMigrate(&User{}); err != nil {
		fmt.Println("迁移建表出错:", err)
		return
	}
	for _, v := range strings.Split(inertUserSql, ";") {
		material.MyDb.Exec(v)
	}
}
type User struct {
	gorm.Model
	Name string
	Age int
	Role string
	Count int
}
var inertUserSql = `
INSERT INTO users VALUES (1, '2022-11-21 21:45:43', '2022-11-21 21:45:43', NULL, 'jack', 12, 'root', 1);
INSERT INTO users VALUES (2, '2022-11-23 16:54:04', '2022-11-23 16:54:09', NULL, 'tom', 11, 'user', 6);
INSERT INTO users VALUES (3, '2022-11-23 16:54:31', '2022-11-23 16:54:35', NULL, 'bob', 30, 'user', 3);
INSERT INTO users VALUES (4, '2022-11-23 16:54:50', '2022-11-23 16:54:53', NULL, 'jimy', 11, 'user', 7);
INSERT INTO users VALUES (5, '2022-11-23 16:55:40', '2022-11-23 16:55:43', NULL, 'teacher', 30, 'root', 5);
INSERT INTO users VALUES (6, '2022-11-23 17:06:10', '2022-11-23 17:06:12', NULL, 'monitor', 12, 'root', 8);
INSERT INTO users VALUES (7, '2022-11-23 17:09:39', '2022-11-23 17:09:42', NULL, 'alice', 11, 'user', 9);
INSERT INTO users VALUES (8, '2022-11-23 19:00:27', '2022-11-23 19:00:30', NULL, 'ck', 11, 'user', 10);
INSERT INTO users VALUES (9, '2022-11-24 17:36:25', '2022-11-24 17:36:27', NULL, 'plus', 11, 'user', 11)
`