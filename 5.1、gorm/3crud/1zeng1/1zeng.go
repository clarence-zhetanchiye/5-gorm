package main

import (
	material "crud/0material"
	"fmt"
	"gorm.io/gorm"
)

type Student struct{//todo:未大写的字段，不能存进数据库 可以参见2ji4cheng2.go
	gorm.Model //todo:悬浮可知，ID在建表时会通过数据库设定为自增。
	Name string
	Age int
}

func main() {
	//todo:初始化MyDb。若不先初始化，下面的material.MyDb都是空的
	material.GetDB()

	//todo:采用蛇形模式自动迁移建立数据库表，如已有表[例如该行代码已运行过一次],再运行该句代码则不操作。
	// 结构体有字段是继承的时，【只要是继承的字段内的字段首字母是大写的】创建表时同样予以继承
	material.MyDb.Exec("DROP TABLE IF EXISTS students")
	material.MyDb.AutoMigrate(&Student{})

	//以下运行后的情况，见最尾部的表数据展示。
	add1()
	add11()
	add2()
	add22()
	add3()

	add4()  //不建议用
	add5()  //不建议用
}

//todo:新增一条完整的记录-------------------------------------------------------
//继承的gorm.Model内的字段会自动赋值
func add1(){
	stu1 := Student{Name: "jack",Age:27}
	result := material.MyDb.Create(&stu1)//todo:会自动根据(&stu1)内的变量stu1的类型名，对应到数据库相应的表。ID已经在迁移建表中被设定为自增
	fmt.Println(result.RowsAffected)//1
	fmt.Println(result.Error)//nil
	fmt.Println("stu1.ID=",stu1.ID)//stu1.ID= 1 todo:执行了Create(&stu1)后，gorm让这里ID已有值。gorm无需额外执行sql来获得该ID，MySQL自动支持，详见标准库。
	fmt.Println("stu1.CreatedAt=",stu1.CreatedAt)//stu1.CreatedAt=  2022-11-15 15:27:27.043 +0800 CST

}

//todo:新增一条字段不完整的记录 ----------------------------------------------------
//由于未赋值字段在结构体变量中已有零值，故存入后数据库表对应的值为相应字段的类型零值；gorm.Model的那些字段仍然是自动赋值
func add11(){
	material.MyDb.Create(&Student{Name:"ha"})//todo:此时未赋值的Age字段是零值，故在数据库表中插入零值；gorm.Model结构体仍自动赋值。

}

//todo:只将某些字段的值存进数据库------------------------------------------------------
//除了Select()中指定的字段外，其他字段的值都是null；但gorm.Model的那些字段仍然会自动赋值
func add2(){
	stu2 := Student{Name:"tom",Age:22}
	material.MyDb.Select("name").Create(&stu2)//todo: 结果是数据库表中除了name有值以及gorm.Model自动赋值，其他列是null。
	fmt.Println("stu2.age=",stu2.Age)//stu2.age= 22// 但数据库表中age是null。
	fmt.Println("stu2.ID=",stu2.ID)//stu2.ID=3 //执行了Create(&stu2)后，gorm让这里的ID有了值。ID在表中自增
}

//todo:跳过某些字段的值不存进数据库------------------------------------------------------
//在Omit()中指定的字段为null外，其他字段不受影响；但gorm.Model的那些字段仍然会自动赋值
func add22(){
	stu22 := Student{Name:"tony",Age:222}
	material.MyDb.Omit("name").Create(&stu22) //todo: 结果是数据库表中除了name是null，其他列都不受影响，字段是默认零值的就插入零值，gorm.Model仍然自动赋值。
	fmt.Println("stu22.age=",stu22.Age)//stu22.age= 222   //数据库中的name是null
	fmt.Println("stu22.ID=",stu22.ID)//stu22.ID=4 //执行了Create(&stu2)后，gorm让这里的ID有了值。ID在表中自增
}


//todo:批量插入---------------------------------------------------------------------
func add3(){
	stu4 := []Student{
		{Name: "xiaoGang",Age: 11},
		{Name: "xiaoHong",Age:14},
	}
	res := material.MyDb.Create(&stu4)
	fmt.Println("res.RowsAffected=",res.RowsAffected)//res.RowsAffected= 2
	fmt.Println("stu4[1].ID=",stu4[1].ID)//stu4[1].ID=6   //执行了Create(&stu2)后，gorm让这里的ID有了值。ID在表中自增
}

//todo:使用CreateInBatches分批进行批量创建，参见GORM文档/CRUD借口/创建-----------------------
func add33() {
	var students []Student
	for i:=1; i<=10000; i++ {
		students = append(students, Student{Name: fmt.Sprintf("jinzhu+%d", i)})
	}
	material.MyDb.CreateInBatches(students, 100)// 指定每批的数量，一批中一次存入100条
}

//---------------------------------------------------------以下不建议用---------------------------------------------------
//todo:用map的方式进行插入需注意！！！
//1、map类型只能是map[string]interface{}，且键值对中，键对应结构体字段、数据库的列，值对应
// 结构体字段、数据库列的值，故不能myMap["xiaoWang"]=25 myMap["xiaoZhao"]=26 然后Create
// 2、除了这里赋值的字段和迁移建表时设定自增的ID，其他继承(包括gorm.Model)的hook自动赋值字段【也称关联associate】在数据库里全为null
func add4(){//可以对照官方文档的例子看
	myMap := make(map[string]interface{})//todo:这样的一个map，只能对应一个结构体，对应插入数据库一条记录
	myMap["name"]="xiaoWang"
	myMap["age"]=25
	res := material.MyDb.Model(&Student{}).Create(&myMap) //这里的gorm语句中的Model()是用来指定要操作的数据库表的。
	fmt.Println("res.RowsAffected=",res.RowsAffected)//res.RowsAffected= 1
}

//todo:用map切片进行批量插入
func add5(){//可以对应看官方文档的示例
	myMap := make([]map[string]interface{},2)
	myMap[0]=make(map[string]interface{})
	myMap[0]["name"]="xiaoQian"
	myMap[0]["age"]=24
	myMap[1]=make(map[string]interface{})
	//myMap[1]["name"]="xiaoSun"  //这里注释不运行，在表中值就会是null
	myMap[1]["age"]=29
	res := material.MyDb.Model(&Student{}).Create(&myMap)
	fmt.Println("res.RowsAffected=",res.RowsAffected)//res.RowsAffected= 2
	//todo:会发现和add4()一样，除了id和这里赋值了的字段【77行由于没赋值数据库表里也会是null】，
	// 其他继承的hook自动赋值的字段都是null
}
/*
id       create_at          updated_at      deleted_at		name       age
1	2022-11-15 10:14:24	2022-11-15 10:14:24		null		jack		27
2	2022-11-15 10:14:24	2022-11-15 10:14:24		null    	ha			0
3	2022-11-15 10:14:24	2022-11-15 10:14:24		null		tom			null
4	2022-11-15 10:14:24	2022-11-15 10:14:24		null		null		222
5	2022-11-15 10:14:24	2022-11-15 10:14:24		null		xiaoGang	11
6	2022-11-15 10:14:24	2022-11-15 10:14:24		null		xiaoHong	14
7			null				null			null		xiaoWang	25
8			null				null			null		xiaoQian	24
9			null				null			null		   null		29
*/