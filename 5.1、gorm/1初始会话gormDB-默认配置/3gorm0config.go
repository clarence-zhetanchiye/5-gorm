package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDB *gorm.DB
var str = "root:413188ok@tcp(localhost:3306)/guan1lian2gorm?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	//todo  gormDB, _ = gorm.Open(mysql.Open(str), &gorm.Config{})
	// 第二个入参gorm.Config{}的所有字段如下所示，如何配置参见 4gormDB初始化配置&GORM的默认，和boardmix图谱中容器二里的注释，
	// 或GORM官网/高级主题/GORM配置。
	_ = &gorm.Config{
		SkipDefaultTransaction:                   false, //配置为true则表示跳过默认事务。
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		Logger:                                   nil,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	}

	dryRun()

}

func dryRun() {
	gormDB, _ = gorm.Open(mysql.Open(str), &gorm.Config{DryRun: true})
	//todo:由下面的打印可知，DryRun为true是到生成最终sql及之前都运行但不执行sql。实际生产中这个配置没什么用处。

	type Student struct {
		Id int
		Name string
		Age int
	}
	var s Student
	if err := gormDB.Debug().Table("students").Where("id=?", 1).Find(&s).Error; err != nil {
		fmt.Println("gorm语句出错=", err)
		return
	}
	fmt.Println("s=", s)
}