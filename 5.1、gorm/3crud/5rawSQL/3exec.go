package main

import (
	material "crud/0material"
	"fmt"
	"gorm.io/gorm"
	"strings"
)
//todo:----------------------------------------------Exec(,)--------------------------------------------------------------
//									     入参只能是非原生查询语句。没有结束方法。
func main() {
	material.GetDB()
	insertTeachers()

	material.MyDb.Exec("INSERT INTO students VALUES (7, 'seven', 1, 1, 1)")

	var t Student
	//todo:Exec()不能用于原生查询。点进Exec()的源码可知其入参被加进了gormDB.Statement.SQL中后，就直接被执行了，不涉及结束方法。
	material.MyDb.Exec("SELECT * FROM students WHERE id=", 1).Find(&t) //todo:报语法错误。
	fmt.Println("t=", t)//t= {0  0 0 0}

	material.MyDb.Exec("UPDATE students SET name=? WHERE id = ?", "haha", 1)
	material.MyDb.Exec("UPDATE students SET score=? WHERE id = ?", gorm.Expr("score * ? + ?", 10, 1), 2)
	//UPDATE students SET score=score * 10 + 1 WHERE id = 2

	material.MyDb.Exec("DELETE FROM students WHERE id = ?", 3)


}
func insertTeachers() {
	material.MyDb.Exec("DROP TABLE IF EXISTS students")
	if err := material.MyDb.AutoMigrate(&Student{}); err != nil {
		fmt.Println("迁移建学生表出错:", err)
		return
	}
	for _, v := range strings.Split(insertStudent, ";") {
		material.MyDb.Exec(v)
	}
}
type Student struct {
	Id int
	Name string
	TeacherId int
	Score int
	GenderId int
}
var insertStudent = `
INSERT INTO students VALUES (1, 'jack', 1, 90, 1);
INSERT INTO students VALUES (2, 'tom', 1, 80, 1);
INSERT INTO students VALUES (3, 'bob', 1, 70, 1);
INSERT INTO students VALUES (4, 'alice', 2, 30, 2);
INSERT INTO students VALUES (5, 'lily', 2, 50, 2);
INSERT INTO students VALUES (6, 'robot', 3, 40, 1)
`