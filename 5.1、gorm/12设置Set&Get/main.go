package main

import (
	"fmt"
	"gorm.io/gorm"
	material "set/0material"
)

func main() {
	initData()

	flag := 123

	var s Student
	material.MyDb.Set("flag", flag).Where("id=?", 1).Find(&s)

	fmt.Println("==========")

	flagt := 456
	var t Student
	material.MyDb.InstanceSet("flag", flagt).Where("id=?", 1).Find(&t)

	/*
	student.v= 123
	---
	student.insV= <nil>
	==========
	student.v= <nil>
	---
	student.insV= 456
	*/

}

type Student struct {
	Id   int
	Name string
	Age  int
	Info
}

func (s *Student) AfterFind(tx *gorm.DB)error {
	v, _ := tx.Get("flag")
	fmt.Println("student.v=", v)

	fmt.Println("---")
	insV, _ := tx.InstanceGet("flag")
	fmt.Println("student.insV=", insV)

	return nil
}

type Info struct {
	Id   int
	Text string
}
func (i *Info) AfterFind(tx *gorm.DB)error {
	v, _ := tx.Get("flag")
	fmt.Println("info.v=", v)

	fmt.Println("+++")
	insV, _ := tx.InstanceGet("flag")
	fmt.Println("info.insV=", insV)

	return nil
}

func initData() {
	material.GetDB()
	material.MyDb.Exec("DROP TABLE IF EXISTS students")
	material.MyDb.AutoMigrate(&Student{})
	stu1 := Student{Name: "jack", Age: 27}
	material.MyDb.Create(&stu1)
}

