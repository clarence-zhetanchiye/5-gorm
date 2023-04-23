package main

import (
	"crud/0material"
	"crud/2cha2/cha"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func main() {
	material.GetDB()
	//插入数据
	insertFactory()

	//同理添加users表的数据
	insertUser()

	//同理添加teachers、users、sex表的数据
	initTeacherStudentsSex()

//------------------------------------------------------------Where条件-------------------------------------------------
	//Where()中搭配 ？ 查询。
	cha.FindByWhere()	//todo:生产中就用这种！！！下面三种不建议用。
	//Where(struct) 查询。
	cha.FindByWhereStruct()
	//根据主键id查询。该方式只用于主键id，主键必须是整型。
	cha.FindById()
	//根据其他条件查询。
	cha.FindByCondition()

//------------------------------------------------------------Or条件----------------------------------------------------
	//Or条件的查询。Or里面的参数填写方式和Where一样
	cha.FindWhereOr()	//todo:建议直接使用cha.FindByWhere()

//------------------------------------------------------------Not条件---------------------------------------------------
	//Not条件查询。Not的性质和Where是一样的
	cha.FindWhereNot()	//todo:建议直接使用cha.FindByWhere()

//-------------------------------------------------------Select查询部分字段-----------------------------------------------
	//Select()查询部分字段。
	cha.SelectFind()
	//智能地查询部分字段！！！！！！
	cha.AutoSelectFind()
//-------------------------------------------------------Count计数------------------------------------------------------
	cha.Count()

//-------------------------------------------------------Distinct------------------------------------------------------
	cha.Distinct()

//-------------------------------------------------------Join----------------------------------------------------------
	cha.Join()

//-------------------------------------------------------GroupBy&Having------------------------------------------------
	cha.GroupBy()


//--------------------------------------------------------Order对查到的数据排序-----------------------------------------
	cha.Order()


//--------------------------------------------------------Limit&Offset对排序结果分页--------------------------------------------------
	cha.Limit()


//--------------------------------------------------------Pluck一或若干列------------------------------------------------
	cha.Pluck()

}


//-----------------------------
func insertFactory() {
	//迁移建表
	material.MyDb.Exec("DROP TABLE IF EXISTS factories")
	material.MyDb.AutoMigrate(&factory{}) //如果数据库中已经有该表，则此处不会重复创建。
	types := []string{"one", "two", "three"}
	belong := ""
	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			belong = "country"
		} else {
			belong = "people"
		}
		k := i % 3
		material.MyDb.Create(&factory{Name: "sanliu" + strconv.Itoa(i), Output: i, Type: types[k], Belong: belong})
		//Product字段在这里没赋值，数据库表里会是空白也即是空字符串。
	}
	material.MyDb.Create(&factory{Name: "sanliu" + strconv.Itoa(100), Product: "broom", Output: 100})
	material.MyDb.Create(&factory{Name: "wuqi" + strconv.Itoa(999), Product: "brush", Output: 999})
	material.MyDb.Create(&factory{Name: "wuqi" + strconv.Itoa(88), Product: "tank", Output: 88})
}
/* 数据表中的数据为
id	name	    product 	output		type	belong
1	sanliu1						1		two		country
2	sanliu2						2		three	people
3	sanliu3						3		one		country
4	sanliu4						4		two		people
5	sanliu5						5		three	country
6	sanliu6						6		one		people
7	sanliu7						7		two		country
8	sanliu8						8		three	people
9	sanliu9						9		one		country
10	sanliu100	broom			100
11	wuqi999		brush			999
12	wuqi88		tank			88
*/
type factory struct {
	Id      int
	Name    string
	Product string
	Output  int
	Type    string //第一产业农业、第二产业工业、第三产业服务业
	Belong  string //国企、私企
}

//--------------------------------
func insertUser() {
	material.MyDb.Exec("DROP TABLE IF EXITS users")
	if err := material.MyDb.AutoMigrate(&User{}); err != nil {
		fmt.Println("迁移建表出错:", err)
		return
	}
	for _, v := range strings.Split(inertUserSql, ";") {
		material.MyDb.Exec(v)
	}
}
/*数据表的数据为
"id"	"created_at"		"updated_at"		"deleted_at"	"name"		"age"	"role"	"count"
1	"21/11/2022 21:45:43"	"21/11/2022 21:45:43"	null		"jack"		12		"root"		1
2	"23/11/2022 16:54:04"	"23/11/2022 16:54:09"	null		"tom"		11		"user"		6
3	"23/11/2022 16:54:31"	"23/11/2022 16:54:35"	null		"bob"		30		"user"		3
4	"23/11/2022 16:54:50"	"23/11/2022 16:54:53"	null		"jimy"		11		"user"		7
5	"23/11/2022 16:55:40"	"23/11/2022 16:55:43"	null		"teacher"	30		"root"		5
6	"23/11/2022 17:06:10"	"23/11/2022 17:06:12"	null		"monitor"	12		"root"		8
7	"23/11/2022 17:09:39"	"23/11/2022 17:09:42"	null		"alice"		11		"user"		9
8	"23/11/2022 19:00:27"	"23/11/2022 19:00:30"	null		"ck"		11		"user"		10
9	"24/11/2022 17:36:25"	"24/11/2022 17:36:27"	null		"plus"		11		"user"		11
*/
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

//-----------------------------
func initTeacherStudentsSex() {
	material.MyDb.Exec("DROP TABLE IF EXISTS teachers")
	if err := material.MyDb.AutoMigrate(&Teacher{}); err != nil {
		fmt.Println("迁移建老师表出错:", err)
		return
	}
	for _, v := range strings.Split(insertTeacher, ";") {
		material.MyDb.Exec(v)
	}

	material.MyDb.Exec("DROP TABLE IF EXISTS students")
	if err := material.MyDb.AutoMigrate(&Student{}); err != nil {
		fmt.Println("迁移建学生表出错:", err)
		return
	}
	for _, v := range strings.Split(insertStudent, ";") {
		material.MyDb.Exec(v)
	}

	material.MyDb.Exec("DROP TABLE IF EXISTS sexes")
	if err := material.MyDb.AutoMigrate(&Sex{}); err != nil {
		fmt.Println("迁移建性别表出错:", err)
		return
	}
	for _, v := range strings.Split(insertSex, ";") {
		material.MyDb.Exec(v)
	}
}
var insertTeacher = `
INSERT INTO teachers VALUES (1, 'smith', 'math');
INSERT INTO teachers VALUES (2, 'green', 'english');
INSERT INTO teachers VALUES (3, 'trump', 'sport')
`
type Teacher struct {
	Id int
	Name string
	Subject string
}
type Student struct {
	Id int
	Name string
	TeacherId int
	Score int
	GenderId int
}
var insertStudent = `
INSERT INTO students VALUES (1, 'jack', 1, 90, 1);
INSERT INTO students VALUES (2, 'tom', 1, 80, 1);
INSERT INTO students VALUES (3, 'bob', 1, 70, 1);
INSERT INTO students VALUES (4, 'alice', 2, 30, 2);
INSERT INTO students VALUES (5, 'lily', 2, 50, 2);
INSERT INTO students VALUES (6, 'robot', 3, 40, 1)
`
type Sex struct {
	Id int
	Gender string
}
var insertSex = `
INSERT INTO sexes VALUES (1, 'male');
INSERT INTO sexes VALUES (2, 'female')
`