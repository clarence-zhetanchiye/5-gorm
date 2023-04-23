package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
//todo: 以下表明，通过获取到gorm包内调用标准库获得的那个和数据库建立的连接【此时自然也能使用标准库的操作数据库的函数方法了】,
// 借标准库设置数据库连接的函数方法，来完成对gormDB连接池的设置————本质上，gormDB和数据库的连接池，就是标准库database/sql和数据库的连接池。

func main() {
	str := "root:413188ok@tcp(localhost:3306)/guan1lian2gorm?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, _ := gorm.Open(mysql.Open(str), &gorm.Config{})

	//todo:获取到gorm通过标准库获得的那个和数据库建立的连接。
	sqlDB,err := gormDB.DB()
	if err != nil {
		fmt.Println("转为sqlDB出错")
		return
	}
	if err := sqlDB.Ping(); err != nil {
		//Ping verifies a connection to the database is still alive, establishing a connection if necessary.
		fmt.Println("ping出错=", err)
		return
	}
	defer sqlDB.Close()

	fmt.Println("数据库当前配置=", sqlDB.Stats())//数据库当前配置= {0 1 0 1 0 0s 0 0 0}

	//todo:gorm无法通过自己获得的跟数据库的连接gormDB来进行MySQL连接池的配置，GORM 需要使用标准库的 database/sql 维护连接池
	// sqlDB.SetConnMaxIdleTime(d time.Duration) 设置连接的最大闲置时间
	// sqlDB.SetConnMaxLifetime(time.Hour)  设置连接可存活的最大时间。
	// sqlDB.SetMaxIdleConns(10)    设置空闲连接池中连接的最大数量
	// sqlDB.SetMaxOpenConns(100)   设置打开数据库连接的最大数量。


	//todo:能使用标准库的操作数据库的函数方法了!
	sql := "select name from animals where age=?"
	var na string
	sqlDB.QueryRow(sql,4).Scan(&na)
	fmt.Println("na=",na)				//na= tiger4
}
