//todo: 当使用 DeletedAt 这个含义的字段创建唯一复合索引时，必须通过 gorm.io/plugin/soft_delete 等插件将字段定义为时间戳（一串数字）
// 这种的数据格式
package main

import (
	material "crud/0material"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"strconv"
)

//todo:下面的
// DeletedAtField:DeletedAt标签指定相应字段记录软删除。这个标签是在"gorm.io/plugin/soft_delete"中被解析识别。
// softDelete:flag标签指定了gorm在删除时给该字段赋值0或1，1表示已软删除。
// softDelete:milli标签指定了gorm在删除时给该字段赋值毫秒
// softDelete:nano标签指定了gorm在删除时给该字段赋值纳秒
type Flag struct {
	Id int
	Name string
	//todo:以下记录软删除时间的字段，默认只会让既有字段顺序中的第一个字段真正记录软删除时间，可以逐个注释逐个试！
	DeleteAt gorm.DeletedAt //todo:gorm语句识别到结构体中有该gorm.DeletedAt类型后会自动启动软删除。悬浮该类型可知本质是time.Time类型。
	DelZero soft_delete.DeletedAt `gorm:"DeleteAtField:DeletedAt"`//todo:悬浮该类型可知本质是一个uint数字类型。
	DelOne soft_delete.DeletedAt `gorm:"softDelete:milli,DeletedAtField:DeletedAt"`
	DelTwo soft_delete.DeletedAt `gorm:"softDelete:nano,DeletedAtField:DeletedAt"`
	IsDel soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
}
func main() {
	material.GetDB()
	insertFlag()

	//todo:删除
	material.MyDb.Table("flags").Where("id=?", 1).Delete(&Flag{})

	//todo:看一看软删除后，记录的软删除时间的字段情况。或者直接看数据库表也行。
	var f Flag
	material.MyDb.Unscoped().Table("flags").Where("id=?", 1).Find(&f)
	fmt.Println("f=", f)
}


func insertFlag() {
	material.MyDb.Exec("DROP TABLE IF EXISTS flags;")
	material.MyDb.AutoMigrate(&Flag{})
	for i := 1; i <= 2; i++ {
		material.MyDb.Create(&Flag{Name: "tiger" + strconv.Itoa(i)})
	}
}