package material

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var(
	MyDb *gorm.DB
	err error
)
func GetDb() {
	dstn := "root:413188ok@tcp(localhost:3306)/guan1lian2gorm?charset=utf8mb4&parseTime=True&loc=Local"
	MyDb, err = gorm.Open(mysql.Open(dstn), &gorm.Config{})
	if err != nil {
		fmt.Println("err=",err)
		return
	}
}
//todo
// 结构体间有继承关系【即数据库表关联】时，如未自定义指定外键和参照，AutoMigrate会按照约定的默认方式
// 自动创建数据库外键约束，参见22gorm/2foreignKey0guan1lian2/我的说明.txt:16，
// 您可以在初始化时禁用此功能：
//  db, err := gorm.Open(mysql.Open(dstn), &gorm.Config{
//  DisableForeignKeyConstraintWhenMigrating: true,
// })
