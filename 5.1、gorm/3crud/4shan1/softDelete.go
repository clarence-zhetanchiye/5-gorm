package main

import (
	material "crud/0material"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

//todo:-------------------------------------------------软删除-----------------------------------------------------------
/*
如果结构体中继承gorm.Model或结构体中有gorm.DeletedAt类型的字段(gorm.Model就包含了该类型的字段)，
将会自动启用软删除特性：拥有软删除能力的结构体实例在gorm语句中调用 Delete 时,数据不会被从数据库中真正删除，而是GORM会
将DeletedAt（或类型为gorm.DeletedAt的列）列置为当前时间, 并且你不能再通过正常的查询方法找到该记录。例如：

//删除
db.Where("age = ?", 20).Delete(&User{})
//实际sql并非删除而是UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

// 在查询时会忽略被软删除的记录，显得似乎是删除了
db.Where("age = 20").Find(&user)
// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;

//查找被软删除的记录
db.Unscoped().Where("age = 20").Find(&users)
// SELECT * FROM users WHERE age = 20;

//硬删除，即不再软删除，直接彻底删除
db.Unscoped().Delete(&User{},"age=?",20)
// DELETE FROM users WHERE age=20;

*/
func main() {
	material.GetDB()
	insertSoft()

	//todo:软删除————由于Soft{}的字段Delete是gorm.DeletedAt类型，因此删除自动是软删除。
	material.MyDb.Table("softs").Where("id=?", 1).Delete(&Soft{})
	//UPDATE `softs` SET `delete`='2022-12-07 23:51:59.144' WHERE id=1 AND `softs`.`delete` IS NULL

	//todo:查询被软删除的数据
	var s Soft
	material.MyDb.Table("softs").Where("id=?", 1).Find(&s)
	//SELECT * FROM `softs` WHERE id=1 AND `softs`.`delete` IS NULL
	fmt.Println("被软删除后查不到=", s)//被软删除后查不到= {0  0  {0001-01-01 00:00:00 +0000 UTC false}}

	//todo:强行查出被软删除的数据
	var f Soft
	material.MyDb.Unscoped().Table("softs").Where("id=?", 1).Find(&f)
	//SELECT * FROM `softs` WHERE id=1
	fmt.Println("被软删除后查到=", f)//被软删除后查到= {1 tiger1 1 sheep {2022-12-07 23:51:59 +0800 CST true}}

	//todo:硬删除----对于已经含有gorm.DeletedAt类型字段而具备软删除特性的结构体，在gorm语句中进行硬删除
	material.MyDb.Unscoped().Table("softs").Where("id=?", 1).Delete(&Soft{})
	//DELETE FROM `softs` WHERE id=1
}
type Soft struct {
	Id   int
	Name string
	Age  int
	Food string
	Delete gorm.DeletedAt
}

func insertSoft() {
	material.MyDb.Exec("DROP TABLE IF EXISTS softs;")
	material.MyDb.AutoMigrate(&Soft{})
	for i := 1; i <= 2; i++ {
		material.MyDb.Create(&Soft{Name: "tiger" + strconv.Itoa(i), Age: i, Food: "sheep"})
	}
}