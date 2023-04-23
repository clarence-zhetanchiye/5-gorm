package cha

import (
	material "crud/0material"
	"fmt"
	"gorm.io/gorm"
)

//只查询哪些字段，决定于gorm语句中.Select()中的入参，见下面示例。
//或决定于设置的承接查询结果的结构体里的那些字段（此时gorm语句需有Model(..),见下面第二部分示例。

//todo:----------------------------------------------Select()查部分字段--------------------------------------------------
func SelectFind() {
	//todo:虽然只查一个字段，但仍必须是结构体，不能是一个string等一般类型变量
	var fac factory
	material.MyDb.Model(&factory{}).Select("name").Where("output=?", 100).Find(&fac)
	// SELECT `name` FROM `factories` WHERE output=100
	fmt.Println("fac1=", fac) //fac1=查出一条数据，且只有name字段有值，其他字段是零值。 fac1= {0 sanliu100  0  }
	fmt.Println("fac1.name=", fac.Name) //fac1.name=sanliu100

	var facs []factory
	material.MyDb.Select("name").Find(&facs) //SELECT `name` FROM `factories` //todo:相当于Pluck()一整列。
	fmt.Println("facs2=", facs)//facs2=查出了表中name这一整列的数据，即facs切片中每个结构体都只有Name字段有值，其他字段都是零值。

	//todo:Select()的可变入参填多个字段（最好这样来填写）。
	facs = []factory{}
	material.MyDb.Select("name, id").Find(&facs)//SELECT name, id FROM `factories`
	fmt.Println("facs3=", facs)//facs3=查出了表中id、name这两整列的数据，即facs切片中每个结构体只有Name和Id字段有值，其他字段都是零值。

	facs = []factory{}
	material.MyDb.Select("name, id").Where("output=?", 1).Find(&facs)
	//SELECT name, id FROM `factories` WHERE output=1
	fmt.Println("facs4=", facs)//facs4= [{1 sanliu1  0  }]

	//todo:Select()的可变入参填多个字段，也可以这样来填写（最好这样来填写）。
	facs = []factory{}
	material.MyDb.Select([]string{"name", "id"}). Where("output=?", 1).Find(&facs)
	//SELECT `name`,`id` FROM `factories` WHERE output=1
	fmt.Println("facs5=", facs) //facs5= [{1 sanliu1  0  }]

	facs = []factory{}
	res, err := material.MyDb.Table("factories").Select("COALESCE(name,?)", "my_name").
		Where("id>?", 9).Rows() //todo:COALESCE是MySQL的一个函数。想用该函数只能这样最后.Rows()。参见下面的例子。
		//SELECT COALESCE(name,'my_name') FROM `factories` WHERE id>9
	if err == nil {
		for res.Next() {
			var nm string
			res.Scan(&nm)
			fmt.Println("facs6=", nm) //facs6= sanliu100    facs6= wuqi999	facs6= wuqi88
		}
	}else {
		fmt.Println("出错=", err)
		return
	}

	type facto struct {
		Name string `gorm:"column:name"`
		Id int `gorm:"column:id"`
	}
	var facx []facto //todo:由结果可知想让COALESCE()这个MySQL函数发挥作用，gorm语句不能最后直接Find()来承接查询结果。
	material.MyDb.Table("factories").Select("COALESCE(name,?)", "my_name").
		Where("id>?", 9).Find(&facx)
		//SELECT COALESCE(name,'my_name') FROM `factories` WHERE id>9
	fmt.Println("facsxxx=", facx) //facsxxx= [{ 0} { 0} { 0}]


	//todo:Select()的可变入参填多个字段时也可以这样写（不建议这样写）。
	facs = []factory{}
	material.MyDb.Select("name", "id").Where("output>?", 9).Find(&facs)
	//SELECT `name`,`id` FROM `factories` WHERE output>9
	fmt.Println("facs7=", facs)//facs7= [{10 sanliu100  0  } {11 wuqi999  0  } {12 wuqi88  0  }]

	facs = []factory{}
	material.MyDb.Select("name", "id").Find(&facs)
	//SELECT `name`,`id` FROM `factories`
	fmt.Println("facs8=", facs) //facs8=查到了表中id、name这两个整列。即facs切片中每个结构体中只有Id、Name两个字段有值，其他字段为零值。
}

//todo:----------------------------------------------智能Select()部分字段-------------------------------------------------
//									（最好用上Model()或Session(&gorm.Session{QueryFields: true})）
// 请逐个看下面五种方式！！！！！
func AutoSelectFind() {
	//-----------------------用新结构体老字段，因新结构体而需指明表名---------------------
	type miniFac struct {
		Name   string //和数据库表的column表头要对应
		Output int    //和数据库表的column表头要对应
	}
	var factoryi []miniFac
	material.MyDb.Table("factories").Where("id<?", 3).Find(&factoryi)
	//SELECT * FROM `factories` WHERE id<3 //可以看见没有指明字段
	fmt.Println("facs9=", factoryi) //facs9= [{sanliu1 1} {sanliu2 2}]

	factoryi = []miniFac{}
	material.MyDb.Model(&factory{}).Where("id<?", 3).Find(&factoryi)
	//SELECT `factories`.`name`,`factories`.`output` FROM `factories` WHERE id<3  //可以看见指明字段了
	fmt.Println("facs10=", factoryi) //facs10= [{sanliu1 1} {sanliu2 2}]

	type miniFa struct {
		Name   string //和数据库表的column表头要对应
		Output int    //和数据库表的column表头要对应
		Gender string //todo:对应不上的字段，如果gorm语句中是指明Table则没关系，如果gorm语句是指明Model或使用了QueryFields则有关系！
		//todo:缺的字段也没关系，不接收就是
	}
	var mini []miniFa
	material.MyDb.Table("factories").Where("id<?", 3).Find(&mini)
	//SELECT * FROM `factories` WHERE id<3  //可以看见没有指明字段
	fmt.Println("mini=", mini) //mini= [{sanliu1 1 } {sanliu2 2 }]

	var minis []miniFa
	material.MyDb.Model(&factory{}).Where("id<?", 3).Find(&minis)
	//SELECT `factories`.`name`,`factories`.`output`,`factories`.`gender` FROM `factories` WHERE id<3  //可以看见指明字段了
	//sql语句中有数据表不存在的列，gorm会报错说有位置的列，同时故查询结果也是空的。
	fmt.Println("minis=", minis)//minis= []


	var mnf []miniFa
	material.MyDb.Session(&gorm.Session{QueryFields: true}).
		Table("factories").Where("id<?", 3).Find(&mnf)
	//SELECT `factories`.`name`,`factories`.`output`,`factories`.`gender` FROM `factories` WHERE id<3 //可以看见指明字段了
	//sql语句中有数据表不存在的列，gorm会报错说有位置的列，同时故查询结果也是空的。
	fmt.Println("mnf=", mnf)//mnf= []
}
