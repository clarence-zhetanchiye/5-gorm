package main

import (
	material "crud/0material"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)
//todo:----------------------------------------------Row()&Rows()--------------------------------------------------------------
//									     转向到标准库中的row.Scan(...)用法。仅用于gorm的查询语句。
func main() {
	material.GetDB()
	insertSoft()

	//将gorm语句转为标准库的row.Scan()
	rowExample()
	rowCount()

	//将gorm语句转为标准库的rows.Scan()
	rowsExample()

	//将gorm语句转为标准库的row.Scan()，但在获取值时又转为使用gorm的gormDb.ScanRows()。实际没什么用处，等价于gorm语句最终Find()
	rowStruct()
	rowsStruct()
}

func rowExample() {
	row := material.MyDb.Table("softs").Select("name, age").Where("id=?", 1).Row()
	//SELECT name, age FROM `softs` WHERE id=1
	var n string
	var a int
	row.Scan(&n, &a) //todo:row.Scan(...)是按gorm语句中Select的字段顺序逐一地赋给自己的入参的，因此接收值的变量名可自定义。
	fmt.Println("name=", n, "age=", a) //name= tiger1   age= 1

	row2 := material.MyDb.Raw("SELECT name, age From softs WHERE Id = ?", 1).Row()
	//SELECT name, age From softs WHERE Id = 1
	name, age := "", 0
	row2.Scan(&name, &age)
	fmt.Println("name=", name, "age=", age) //name= tiger1   age= 1
}

func rowCount() {
	var num int64
	material.MyDb.Table("softs").Where("name LIKE ?", "%tiger%").Count(&num)
	//SELECT count(*) FROM `softs` WHERE name LIKE '%tiger%'
	fmt.Println("num=", num)//num= 2

	//todo:下面这样写，就不需要额外设置一个结构体让里面的一个字段接收COUNT值，再通过Find()获得值了。
	var number int
	row := material.MyDb.Table("softs").Select("COUNT(*) AS n").Row()
	//SELECT COUNT(*) AS n FROM `softs`
	row.Scan(&number)
	fmt.Println("number=", number)//number= 2
}

func rowsExample() {
	var nm string
	var ag int
	rows, err := material.MyDb.Table("softs").Select("name, age").
		Where("name LIKE ?", "%tiger%").Rows()
	//SELECT name, age FROM `softs` WHERE name LIKE '%tiger%'
	if err != nil {
		fmt.Println("rows err=", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&nm, &ag)//todo:rows.Scan(...)是按gorm语句中Select的字段顺序逐一地赋给自己的入参的，因此接收值的变量名可自定义。
		fmt.Println("nm=", nm, "ag=", ag)
		//nm= tiger1 ag= 1
		//nm= tiger2 ag= 2
	}


	var nms string
	var ags int
	rows, err = material.MyDb.Raw("SELECT name, age FROM softs WHERE name LIKE ?", "%tiger%").Rows()
	//SELECT name, age FROM softs WHERE name LIKE '%tiger%'
	if err != nil {
		fmt.Println("rows err=", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&nms, &ags)
		fmt.Println("nms=", nms, "ags=", ags)
		//nms= tiger1 ags= 1
		//nms= tiger2 ags= 2
	}

}

func rowStruct() {
	//todo:虽然知道只会查到一条记录，但由于gorm的ScanRows(,)的第一个入参必须是标准库的*sql.Rows类型，因此最后必须是 .Rows()
	rows, err := material.MyDb.Table("softs").Select("name, age").Where("id = ?", 1).Rows()
	//SELECT name, age FROM `softs` WHERE id = 1
	if err != nil {
		fmt.Println("rows err=", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var s Soft
		material.MyDb.ScanRows(rows, &s)//todo:第一个入参必须是*sql.Rows类型。
		fmt.Println("s=", s)//s= {0 tiger1 1  {0001-01-01 00:00:00 +0000 UTC false}}
	}

	rows, err = material.MyDb.Raw("SELECT name, age FROM softs WHERE id= ?", 1).Rows()
	//SELECT name, age FROM softs WHERE id= 1
	if err != nil {
		fmt.Println("rows err=", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var f Soft
		material.MyDb.ScanRows(rows, &f)
		fmt.Println("f=", f)//f= {0 tiger1 1  {0001-01-01 00:00:00 +0000 UTC false}}
	}
}

func rowsStruct() {
	rows, err := material.MyDb.Table("softs").Select("name, age").
		Where("name LIKE ?", "%tiger%").Rows()
	//SELECT name, age FROM `softs` WHERE name LIKE '%tiger%'
	if err != nil {
		fmt.Println("rows err=", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var s Soft
		material.MyDb.ScanRows(rows, &s)//todo:第一个入参必须是*sql.Rows类型。
		fmt.Println("s=", s)
		//s= {0 tiger1 1  {0001-01-01 00:00:00 +0000 UTC false}}
		//s= {0 tiger2 2  {0001-01-01 00:00:00 +0000 UTC false}}
	}


	rows, err = material.MyDb.Raw("SELECT name, age FROM softs WHERE name LIKE ?", "%tiger%").Rows()
	//SELECT name, age FROM softs WHERE name LIKE '%tiger%'
	if err != nil {
		fmt.Println("rows err=", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var f Soft
		material.MyDb.ScanRows(rows, &f)
		fmt.Println("f=", f)
		//f= {0 tiger1 1  {0001-01-01 00:00:00 +0000 UTC false}}
		//f= {0 tiger2 2  {0001-01-01 00:00:00 +0000 UTC false}}
	}
}

type Soft struct {
	Id     int
	Name   string
	Age    int
	Food   string
	Delete gorm.DeletedAt
}

func insertSoft() {
	material.MyDb.Exec("DROP TABLE IF EXISTS softs;")
	material.MyDb.AutoMigrate(&Soft{})
	for i := 1; i <= 2; i++ {
		material.MyDb.Create(&Soft{Name: "tiger" + strconv.Itoa(i), Age: i, Food: "sheep"})
	}
}
