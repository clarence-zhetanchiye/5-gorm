//todo:在向数据库表新增的结构体内有继承结构体时的情况。
package main

import (
	material "crud/0material"
	"fmt"
)

type teacher struct {
	Name string
	Age int
	gender string //todo:小写字段是不能存进数据库的，建表时建成的表中也不会有这一列
	Hobby 
	   //todo:匿名继承字段内的字段会被视为和结构体的其他字段是一样的。但不大写而写为hobby则匿名继承字段整个都不会存进数据库。
	   // 匿名继承字段内如有和父字段重名的，匿名继承字段的该子字段在创建表时会被忽略。

	//F Family
		//todo: 有名继承字段，需使用`gorm:"embedded"`标签实现和匿名继承一样的效果;
		// 或gorm的外键参照标签来实现关联(teacher结构体中继承Family结构体，teacher对应teachers表，Family对应Families表，两表有关联关系)，
		// 否则报错且不会成功建表，详见 我的说明.txt 中的第5条
}
type Hobby struct {
	Name string //todo:因和teacher.Name字段重名，故创建表时被忽略
	Frequent int
}
type Family struct{
	Num int
}
func main() {
	material.GetDB()

	//todo:自动迁移建表,结构体有字段是继承的(只要是继承的字段内的字段首字母是大写的)创建表时同样予以继承,本例中表格是三列为 name age frequent
	material.MyDb.AutoMigrate(&teacher{})


	//以下运行后的情况，见最尾部的表数据展示。
	insert0()
	insert1()
	insert2()  //不建议用

}

//todo:插入一条----------------------------------------------------------------------
//数据库表中name age frequent三列下 成功插入  laoWu 0 2
// 此时虽然未给Age字段赋值，但变量申明为结构体后其下各字段初始值均为零值
// 故运行时控制台可以看到sql语句里age是插入的0，数据库表该列也是0，而不是null
func insert0(){
	res := material.MyDb.Create(&teacher{Name: "laoWu", Hobby: Hobby{Name: "cheese", Frequent: 2}})
	fmt.Println("res.RowsAffected=",res.RowsAffected)//res.RowsAffected= 1

}

//todo:用结构体切片进行批量插入-------------------------------------------------------
//数据库表中name age frequent三列下 成功插入了  laowang 0 2 和 laoliu 35 3
func insert1(){
	teachs := []teacher{
		{Name: "laowang",Hobby:Hobby{Name: "soccer",Frequent: 2}},
		{Name: "laoliu",Age:35,Hobby:Hobby{"pingpong",3}},
	}
	res := material.MyDb.Create(&teachs)
	fmt.Println("res.RowsAffected=",res.RowsAffected) //res.RowsAffected= 2
}




//----------------------------------------------------以下不建议用--------------------------------------------------------
//todo:用map新增一条   这个需要特别注意一下注释掉的内容！！！
// 成功插入一条 laoli 66 7
func insert2(){
	res := material.MyDb.Model(&teacher{}).Create(&map[string]interface{}{//这里的gorm语句中的Model()是用来指定要操作的数据库表的。
		"name":   "laoLi",
		"age":    66,
		//"gender": "male",//这一行若写了，由于创建的数据库表没有这一列，导致整体create不进去
		//"hobby":  Hobby{Name: "taiji", Frequent: 7},//不能这样写，这样写create不进去
		//"name":"taiji",//虽然这是兴趣的名称，但和老师的姓名重名，写了会提示错误
		"frequent":7, //这里赋值，成功插入数据库表了
	})
	fmt.Println("res.RowsAffected=",res.RowsAffected)//res.RowsAffected= 1
}

/*
name		age		frequent
laoWu		0		2
laowang		0		2
laoliu		35		3
laoLi		66		7
*/
