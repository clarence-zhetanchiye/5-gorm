package cha

import (
	material "crud/0material"
	"fmt"
)

//todo:-----------------------------------------------------Count计数---------------------------------------------------
func Count() {
	var num int64
	material.MyDb.Table("factories").Where("output>?", 9).Count(&num)
	//SELECT count(*) FROM `factories` WHERE output>9
	fmt.Println("num=", num)//num= 3

	var count int64
	material.MyDb.Model(&factory{}).Where("output>?", 9).Count(&count)
	//SELECT count(*) FROM `factories` WHERE output>9
	fmt.Println("count=", count)//count= 3

	var amount int64
	material.MyDb.Table("factories").Distinct("name").Where("output>?", 9).Count(&amount)
	//SELECT COUNT(DISTINCT(`name`)) FROM `factories` WHERE output>9
	fmt.Println("amount=", amount)//amount= 3

	var number Number
	material.MyDb.Table("factories").Select("COUNT(DISTINCT name) AS num").
		Where("output>?", 9).Find(&number)
	//SELECT COUNT(DISTINCT name) AS num FROM `factories` WHERE output>9
	fmt.Println("number=", number)//number= {3}
}
type Number struct {
	Num int64
}