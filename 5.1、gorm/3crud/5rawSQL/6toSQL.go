package main

import (
	material "crud/0material"
	"fmt"
	"gorm.io/gorm"
	"strings"
)
//todo:--------------------------------------------------ToSQL----------------------------------------------------------
//												返回生成的 SQL 但不执行
func main() {
	material.GetDB()
	insertSexes()

	//todo:看数据库表中数据可知，相应的sql语句实际并未执行。
	sqlString2 := material.MyDb.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Table("sexes").Where("id=?", 2).Update("gender", "x")
		//UPDATE `sexes` SET `gender`='x' WHERE id=2
	})
	fmt.Println("sql语句=", sqlString2)//sql语句= UPDATE `sexes` SET `gender`='x' WHERE id=2

	sqlString := material.MyDb.ToSQL(func(tx *gorm.DB) *gorm.DB {
		var p Sex
		defer fmt.Println("查询结果p=", p)//查询结果p= {0 } //todo:可见查询的sql实际并未真正执行。
		return tx.Table("sexes").Select("id, gender").Where("id=?", 1).Find(&p)
		//SELECT id, gender FROM `sexes` WHERE id=1
	})
	fmt.Println("sql语句=", sqlString)//sql语句= SELECT id, gender FROM `sexes` WHERE id=1
}
func insertSexes() {
	material.MyDb.Exec("DROP TABLE IF EXISTS sexes")
	if err := material.MyDb.AutoMigrate(&Sex{}); err != nil {
		fmt.Println("迁移建性别表出错:", err)
		return
	}
	for _, v := range strings.Split(insertSex, ";") {
		material.MyDb.Exec(v)
	}
}
type Sex struct {
	Id int
	Gender string
}
var insertSex = `
INSERT INTO sexes VALUES (1, 'male');
INSERT INTO sexes VALUES (2, 'female')
`
