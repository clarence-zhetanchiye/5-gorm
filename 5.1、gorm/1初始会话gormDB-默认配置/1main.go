//todo:以下点进mysql.Open()源码中的Initialize方法中的sql.Open可知，gorm本质上仍是通过标准库 database/sql 获取到数据库连接的。

package main
import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init(){
	get() //方式一：gorm语句操作MySQL
	take() //方式二：标准库操作MySQL
}

var gormDb *gorm.DB
var err error
func get(){
	dsn := "root:413188ok@tcp(localhost:3306)/guan1lian2gorm?charset=utf8mb4&parseTime=True&loc=Local" //todo:parseTime用以正确处理time.Time,charset=utf8mb4以支持完整的UTF-8编码
	gormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		//todo:其中gorm.Open的第一个入参可替换为mysql.New(mysql.Config{})，另可见GORM官网的 入门指南/链接到数据库。
		// 第二个入参的分析见 3gorm0config.go
	if err!=nil {
		fmt.Println("连接数据库失败:",err)
		return
	}
}

var sqlDb *sql.DB
var err2 error
func take(){
	sqlDb, err2 = sql.Open("mysql", "root:413188ok@tcp(localhost:3306)/shi3yong4gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err2 !=nil {
		fmt.Println("原生连接数据库有误")
		return
	}
}

func main() {
	type animal struct{ //只查询部分字段的值，故未写上所有字段
		Name string
		Age int
	}
	var a animal //数据库里已经有animals这个表，故直接查
	gormDb.Where("name=?", "tiger5").Find(&a)
	fmt.Println("a=",a)					//a= {tiger5 5}



	sql := "select name from animals where age=?"
	row := sqlDb.QueryRow(sql, 4)
	var theName string
	row.Scan(&theName)
	fmt.Println("thename=",theName)		//thename= tiger4
}

