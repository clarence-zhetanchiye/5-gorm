package cha

import (
	material "crud/0material"
	"fmt"
)

//todo:--------------------------------------------------Join连表-------------------------------------------------------
type TeachStudents struct {
	Id int
	Name string
	Subject string

	StudentId int `gorm:"column:sId"` //这里的标签和下面的gorm语句中Select()里的AS后面的别名对应上就行。
	StudentName string `gorm:"column:sName"`
	StudentTeachId int `gorm:"column:sTeacherId"`
}
func Join() {
	//todo:连表查询。若不同的表有同名的字段，记得加上表明和起不同的别名，并和承接的结构体的gorm:column标签对应上。
	var teachStus []TeachStudents       //todo:点进Joins看源码和注释
	material.MyDb.Table("teachers").Joins("INNER JOIN students ON students.teacher_id=teachers.id").
		Select("teachers.id, teachers.name, teachers.subject, " + //todo:因两个表的字段名称冲突，因此字段名前需加表名且加上AS别名
			"students.id AS sId, students.name AS sName, students.teacher_id AS sTeacherId").
		Find(&teachStus)
	//SELECT teachers.id, teachers.name, teachers.subject, students.id AS sId, students.name AS sName,
	//students.teacher_id AS sTeacherId FROM `teachers` INNER JOIN students ON students.teacher_id=teachers.id
	fmt.Println("teachStus1=", teachStus)
	/*
	teachStus1= [{1 smith math 1 jack 1} {1 smith math 2 tom 1} {1 smith math 3 bob 1}
				{2 green english 4 alice 2} {2 green english 5 lily 2}
				{3 trump sport 6 robot 3}]
	*/


	//todo:在JOIN表的同时对哪些数据参与到JOIN进行筛选。即本例中的 AND students.score>?", 60
	teachStus = []TeachStudents{}
	material.MyDb.Table("teachers").
		Joins("INNER JOIN students ON students.teacher_id=teachers.id AND students.score>?", 60).
		Select("teachers.id, teachers.name, teachers.subject, " +
			"students.id AS sId, students.name AS sName, students.teacher_id AS sTeacherId").
		Find(&teachStus)
	//SELECT teachers.id, teachers.name, teachers.subject, students.id AS sId, students.name AS sName, students.teac
	//her_id AS sTeacherId FROM `teachers` INNER JOIN students ON students.teacher_id=teachers.id AND students.score>60
	fmt.Println("teachStus2=", teachStus)
	//teachStus2= [{1 smith math 1 jack 1} {1 smith math 2 tom 1} {1 smith math 3 bob 1}]


	//todo:多表连接
	var teachStuSex []StuTeachSex
	j := material.MyDb.Table("teachers").
		Joins("INNER JOIN students ON students.teacher_id=teachers.id AND students.score<?", 60).
		Joins("INNER JOIN sexes ON students.gender_id=sexes.id").
		Select("teachers.id, teachers.name, teachers.subject, students.name AS sName, students.score, sexes.gender")

	j.Find(&teachStuSex)
	//SELECT teachers.id, teachers.name, teachers.subject, students.name AS sName, students.score, sexes.gender FROM
	//`teachers` INNER JOIN students ON students.teacher_id=teachers.id AND students.score<60 INNER JOIN sexes ON
	//students.gender_id=sexes.id
	fmt.Println("teachStus3=", teachStuSex)
	//teachStus3= [{2 green english 0 alice 0 female} {2 green english 0 lily 0 female} {3 trump sport 0 robot 0 male}]


	//todo:对JOIN后形成的连接表进行Where筛选查询。
	teachStuSex = []StuTeachSex{}
	j.Where("students.id>?", 5).Find(&teachStuSex)
	//SELECT teachers.id, teachers.name, teachers.subject, students.name AS sName, students.score, sexes.gender FROM
	//`teachers` INNER JOIN students ON students.teacher_id=teachers.id AND students.score<60 INNER JOIN sexes ON
	//students.gender_id=sexes.id WHERE students.id>5
	fmt.Println("teachStus4=", teachStuSex)
	//teachStus4= [{3 trump sport 0 robot 0 male}]



	//todo:Join一个衍生表。也就是Join一个查询子句。
	//query := db.Table("order").Select("MAX(order.finished_at) as latest").
	//	Joins("left join user user on order.user_id = user.id").Where("user.age > ?", 18).Group("order.user_id")
	//
	//res := db.Model(&Order{}).Joins("JOIN (?) q on order.finished_at = q.latest", query).Scan(&results)
	//
	//SELECT `order`.`user_id`,`order`.`finished_at` FROM `order` join
	//(
	//	SELECT MAX(order.finished_at) as latest FROM `order` left join user user on order.user_id = user.id
	//	WHERE user.age > 18 GROUP BY `order`.`user_id`
	//) q
	//on order.finished_at = q.latest

	//todo：下面这种Joins的入参的形式，是用于预加载。所谓的预加载，和结构体继承及数据表隶属有关，实际生产一般不使用有继承结构体的结构体迁移建表，
	// 因此生产中也不会这样使用Joins，也不用管预加载这个概念。参见GORM文档/CRUD接口/查询/Joins预加载的示例 和 GORM文档/关联/预加载
	/*
		db.Joins("Company").Find(&users)
		SELECT `users`.`id`,`users`.`name`,`users`.`age`,
		`Company`.`id` AS `Company__id`,
		`Company`.`name` AS `Company__name`
		FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;
	*/
}
type StuTeachSex struct {
	Id int
	Name string
	Subject string

	StudentId int `gorm:"column:sId"` //这里的标签和下面的gorm语句中Select()里的AS后面的别名对应上就行。
	StudentName string `gorm:"column:sName"`
	StudentTeachId int `gorm:"column:sTeacherId"`

	Gender string
}