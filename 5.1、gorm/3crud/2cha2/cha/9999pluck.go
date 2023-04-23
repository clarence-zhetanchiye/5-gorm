package cha

import (
	material "crud/0material"
	"fmt"
)

//todo --------------------------------------------Pluck查一列（等价于Select一列）------------------------------------------
// 以下实现都可以通过Select()配合Find()来实现。
// 以下直接从GORM官网复制来，放心用
func Pluck() {
	var score []int64
	material.MyDb.Model(&Student{}).Pluck("score", &score)
	//SELECT `score` FROM `students`
	fmt.Println("score1=", score) //score1= [90 80 70 30 50 40]

	var names []string
	material.MyDb.Table("students").Pluck("name", &names)
	// SELECT `name` FROM `students`
	fmt.Println("names2=", names) //names2= [jack tom bob alice lily robot]

	names = []string{}
	material.MyDb.Table("students").Where("id>?", 4).Pluck("name", &names)
	// SELECT `name` FROM `students` WHERE id>4
	fmt.Println("names3=", names)//names3= [lily robot]

	//todo:查询一列，且展示一列里不重复的内容。Distinct Pluck
	var gender []int64
	material.MyDb.Model(&Student{}).Distinct().Pluck("gender_id", &gender)
	// SELECT DISTINCT `gender_id` FROM `students`
	fmt.Println("gender=", gender)//gender= [1 2]



	//todo:一次获取多列，则只能采用Select("column1, column2, ...").Where(...)

}