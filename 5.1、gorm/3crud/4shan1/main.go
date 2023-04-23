package main

import (
	material "crud/0material"
	"gorm.io/gorm"
	"strconv"
)

func main() {
	material.GetDB()
	insertAnimals()

	deleteByWhere()
	deleteAll()

}
//todo:删除就用这里的,不用管官网其他的若干冗余的方式。但关于软删除,详见官网。
// GORM官网/CRUD接口/删除/返回删除行的数据 ，对于这部分，MySQL本身并不支持。

//todo:--------------------------------------------------使用where()来删除------------------------------------------------
func deleteByWhere() {
	material.MyDb.Where("id=?", 1).Delete(&animal{})
	//DELETE FROM `animals` WHERE id=1

	material.MyDb.Where("id IN ?", []int{6, 7}).Delete(&animal{})
	//DELETE FROM `animals` WHERE id IN (6,7)

	material.MyDb.Where("name=? AND age=?", "tiger4", 5).Delete(&animal{})
	//DELETE FROM `animals` WHERE name='tiger4' AND age=5

	material.MyDb.Where("name LIKE ?", "%10%").Delete(&animal{})
	//DELETE FROM `animals` WHERE name LIKE '%10%'

	//Where()内的条件也可以直接挪到Delete里面，但不建议这样使用--------------------------------
	material.MyDb.Delete(&animal{}, "name=?", "tiger8")
	//DELETE FROM `animals` WHERE name='tiger8'

	material.MyDb.Delete(&animal{}, "name LIKE ?", "%2%")
	//DELETE FROM `animals` WHERE name LIKE '%2%'

	//Where()内的条件是主键Id时也可以直接挪到Delete里面，但不建议这样使用------------------------
	//material.MyDb指明了数据库名,Delete内的&animal{}指明了表名
	material.MyDb.Delete(&animal{}, 1) //DELETE FROM `animals` WHERE `animals`.`id` = 1 //todo：虽id=1的记录已不存在，本次删并不报错或其他反应。
	material.MyDb.Delete(&animal{}, "2")//DELETE FROM `animals` WHERE `animals`.`id` = '2' //todo:删除不存在的内容不会报错或怎样。

	material.MyDb.Delete(&animal{}, []int{3, 4})//DELETE FROM `animals` WHERE `animals`.`id` IN (3,4)
}

//todo:-----------------------------------------------------删除全表-----------------------------------------------------
func deleteAll() {
	//全局删除不能像下面这样写，否则会报错；除非在初始获取gormDB时gorm.Config{}的AllowGlobalUpdate设为true
	material.MyDb.Delete(&animal{})//DELETE FROM `animals` 虽打印出来但未真正执行，而且gorm报错说WHERE conditions required

	//todo:方式一，使用原生SQL
	material.MyDb.Exec("DELETE FROM animals")//DELETE FROM animals

	//todo:方式二，在gorm语句中使用Where("1=1")
	material.MyDb.Where("1=1").Delete(&animal{}) //DELETE FROM `animals` WHERE 1=1

	//todo:方式三，在Session中设置AllowGlobalUpdate设为true
	material.MyDb.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&animal{})//DELETE FROM `animals`
}

type animal struct {
	Id   int
	Name string
	Age  int
	Food string
}

func insertAnimals() {
	material.MyDb.Exec("DROP TABLE IF EXISTS animals;")
	material.MyDb.AutoMigrate(&animal{})
	for i := 1; i <= 10; i++ {
		material.MyDb.Create(&animal{Name: "tiger" + strconv.Itoa(i), Age: i, Food: "sheep"})
	}
}