package main

import (
	material "crud/0material"
	"fmt"
	"strings"
)
//todo:----------------------------------------------Raw()--------------------------------------------------------------
//									 入参只能是原生查询语句。结束方法建议是Find()
func main() {
	material.GetDB()
	insertGoods()

	var g good
	material.MyDb.Raw("SELECT * FROM goods WHERE id = ?", 1).Find(&g) //todo:可见Find()等价于下面的Scan()
	fmt.Println("g1=", g)//g1= {1 apple 5 100 111}

	g = good{}
	material.MyDb.Raw("SELECT * FROM goods WHERE id = ?", 1).Scan(&g)
	fmt.Println("g2=", g)//g2= {1 apple 5 100 111}

	var gs []good
	material.MyDb.Raw("SELECT * FROM goods WHERE name = ?", "apple").Find(&gs)
	fmt.Println("gs1=", gs)//gs1= [{1 apple 5 100 111} {2 apple 8 200 222}]

	gs = []good{}
	material.MyDb.Raw("SELECT * FROM goods WHERE name = ?", "apple").Scan(&gs)//SELECT * FROM goods WHERE name = 'apple'
	fmt.Println("gs2=", gs)//gs2= [{1 apple 5 100 111} {2 apple 8 200 222}]

	//todo:Raw()只能用于查询：因为Raw()点进源码可知，它只是将入参加进了gormDB.Statement.SQL中，还缺少finish方法，且即使写上非查询的
	// 结束方法，也是非法的gorm语句写法。
	//material.MyDb.Raw("UPDATE goods SET name = ? WHERE id = ?", "x", "peach") //todo:不会被执行
	//
	//material.MyDb.Raw("INSERT INTO goods VALUES(7, 'seven', 7, 700, 7777)") //todo:不会被执行
	//
	//material.MyDb.Raw("WHERE id = ?", 6) //todo:不会被执行
	//material.MyDb.Raw("WHERE id = ?", 6).Delete(&good{}) //todo:报错说WHERE conditions required
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