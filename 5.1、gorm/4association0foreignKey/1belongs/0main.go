package main
import (
	"fmt"
	material "guanlian/0material"
)
//todo:---------------------------------------采用gorm默认外键的方式-------------------------------------------------------

//belongs关系就是： 产生一个“从”就会伴随产生“主”。
//详细解释如下：
//向“主”表新增一条“主”并不会导致同时在“从”表里新增一条“从”，相反，向“从”表里新增一条“从”却肯定导致同时
//在“主”表里新增一条“主” ,这就是belongs关系的本质。因此——————有m个“从”就有m个“主”， 这m个“从”里可能
//有若干个有共同的“主”，但这些共同的“主”在“主”的数据库表里仍然是不同的一行行记录[因为虽然其他字段相同，
//但自增的ID不同]；有n个“主”却不一定有n个“从”，因为可能是直接新增的“主”，而这并不会导致新增“从”。


//主  【相当于购物车】
type Company struct {
	ID int  //todo:gorm在迁移建表时默认ID是primary_key和Auto_increment，"从表"不指定references时gorm会默认"父表"该字段为references
	Name string
	Code int
}
//从 【相当于购物项】
type Employee struct {
	ID int   //todo:gorm在迁移建表时会默认ID是PRIMARY KEY和AUTO_INCREMENT
	Name string
	CompanyID int //todo:gorm在迁移建表时会默认该字段为外键,因为该字段是“主”的结构体名加上其primaryKey字段组成的。因是外键，故不能再自主赋值
	ComCode int
	Company Company
}

func main() {
	material.GetDb()
	material.MyDb.Exec("DROP TABLES IF EXISTS companies")
	material.MyDb.Exec("DROP TABLES IF EXISTS employees")
	material.MyDb.AutoMigrate(&Company{},&Employee{})

	e := Employee{Name:"erlang",ComCode:666,Company: Company{Name:"cpGroup",Code: 234}}
	ee := Employee{Name:"sanlang",ComCode:777,Company: Company{Name:"otherGroup",Code: 567}}
	f := Employee{Name:"silang",ComCode:888,Company: Company{Name:"otherGroup",Code: 567}}
	//todo:新增公司不会导致新增一条员工记录，新增员工一定会导致新增一条公司记录
	material.MyDb.Create(&e)
	material.MyDb.Create(&ee)
	material.MyDb.Create(&f)
	eee := Company{Name: "wanda",Code: 888}
	material.MyDb.Create(&eee)//todo:新增公司不会导致新增一条员工记录，新增员工一定会导致新增一条公司记录
	fmt.Println("运行结束")
}
/*
employees 的表如下
+----+---------+------------+----------+
| id | name    | company_id | com_code |
+----+---------+------------+----------+
|  1 | erlang  |          1 |      666 |
|  2 | sanlang |          2 |      777 |
|  3 | silang  |          3 |      888 |
+----+---------+------------+----------+
companies 表的如下
+----+------------+------+
| id | name       | code |
+----+------------+------+
|  1 | cpGroup    |  234 |
|  2 | otherGroup |  567 |
|  3 | otherGroup |  567 |
|  4 | wanda      |  888 |
+----+------------+------+
*/


