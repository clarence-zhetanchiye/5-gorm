package cha

import (
	material "crud/0material"
	"fmt"
)

//todo:--------------------------------------------------Group分组------------------------------------------------------
//todo:Group分组及分组后对组或组内的Having筛选。参见GORM文档/CRUD接口/查询/

/*
"id"	"created_at"		"updated_at"		"deleted_at"	"name"		"age"	"role"	"count"
1	"21/11/2022 21:45:43"	"21/11/2022 21:45:43"	null		"jack"		12		"root"		1
2	"23/11/2022 16:54:04"	"23/11/2022 16:54:09"	null		"tom"		11		"user"		6
3	"23/11/2022 16:54:31"	"23/11/2022 16:54:35"	null		"bob"		30		"user"		3
4	"23/11/2022 16:54:50"	"23/11/2022 16:54:53"	null		"jimy"		11		"user"		7
5	"23/11/2022 16:55:40"	"23/11/2022 16:55:43"	null		"teacher"	30		"root"		5
6	"23/11/2022 17:06:10"	"23/11/2022 17:06:12"	null		"monitor"	12		"root"		8
7	"23/11/2022 17:09:39"	"23/11/2022 17:09:42"	null		"alice"		11		"user"		9
8	"23/11/2022 19:00:27"	"23/11/2022 19:00:30"	null		"ck"		11		"user"		10
9	"24/11/2022 17:36:25"	"24/11/2022 17:36:27"	null		"plus"		11		"user"		11
*/
func GroupBy() {
	//分组--------------------------------------------------------------------------------------------------------------
	var users1 []User
	material.MyDb.Table("users").Where("id<?", 9).Group("role").Find(&users1)
	//SELECT * FROM `users` WHERE id<9 AND `users`.`deleted_at` IS NULL GROUP BY `role`
	fmt.Println("users1=", users1) //查出了id为1、2的为顺序的行

	//分组且撑开每个组的组员-----------------------------------------------------------------------------------------------
	var users2 []User
	material.MyDb.Table("users").Where("id<?", 9).Group("role, id").Find(&users2)
	//SELECT * FROM `users` WHERE id<9 AND `users`.`deleted_at` IS NULL GROUP BY role, id
	fmt.Println("users2=", users2) //查出了id为1、5、6、2、3、4、7、8的为顺序的行

	//更多分组，见自己整理的分组sql的示例word文档。


	//分组且针对组进行Having筛选--------------------------------------------------------------------------------------------
	var users3 []User
	material.MyDb.Table("users").Where("id<?", 9).Group("age").
		Having("age<?", 30).Find(&users3)
	//SELECT * FROM `users` WHERE id<9 AND `users`.`deleted_at` IS NULL GROUP BY `age` HAVING age<30
	fmt.Println("users3=", users3) //查出了id为2、1为顺序的行

	var users4 []User
	material.MyDb.Table("users").Where("id<?", 9).Group("age, id").
		Having("age<?", 30).Find(&users4)
	// SELECT * FROM `users` WHERE id<9 AND `users`.`deleted_at` IS NULL GROUP BY age, id HAVING age<30
	fmt.Println("users4=", users4) //查出了id为2、4、7、8、1、6为顺序的行

	//分组且对每个组内成员进行Having筛选
	//（此时一定要在分组时就撑开组员，即Group()入参中要有逗号隔开的第二个撑开组员的分组依据）-----------------------------------------
	var users5 []User
	material.MyDb.Table("users").Where("id<?", 9).Group("age, id").
		Having("count>?", 1).Find(&users5)
	// SELECT * FROM `users` WHERE id<9 AND `users`.`deleted_at` IS NULL GROUP BY age, id HAVING count>1
	fmt.Println("users5=", users5) //查出了id为2、4、7、8、6、3、5为顺序的行

	//分组且对每个组内成员进行Having筛选，Having筛选中伴随AND、NOT、OR----------------------------------------------------------
	var users6 []User
	material.MyDb.Table("users").Where("id<?", 9).Group("age, id").
		Having("role=?", "root").Having("count=?", 1).Find(&users6)
	//SELECT * FROM `users` WHERE id<9 AND `users`.`deleted_at` IS NULL GROUP BY age, id HAVING role='root' AND count=1
	fmt.Println("users6=", users6) //查出了id为1这一行

	var users7 []User
	material.MyDb.Table("users").Where("id<?", 9).Group("age, id").
		Having("role=? AND count=?", "root", 1).Find(&users7)
	//SELECT * FROM `users` WHERE id<9 AND `users`.`deleted_at` IS NULL GROUP BY age, id HAVING role='root' AND count=1
	fmt.Println("users7=", users7) //查出了id为1这一行

	var users8 []User
	material.MyDb.Table("users").Where("id<?", 9).Group("age, id").
		Having("role=?", "root").Not("count=?", 1). //todo:这样想Having的筛选中伴随NOT的写法是不对的
		Find(&users8)
	//SELECT * FROM `users` WHERE id<9 AND NOT count=1 AND `users`.`deleted_at` IS NULL GROUP BY age, id HAVING role='root'
	fmt.Println("users8=", users8) //查出了id为6、5为顺序的行

	var users9 []User
	material.MyDb.Table("users").Where("id<?", 9).Group("age, id").
		Having("role=? AND NOT count<?", "root", 6).Find(&users9) //todo:NOT前一定要有AND
	//SELECT * FROM `users` WHERE id<9 AND `users`.`deleted_at` IS NULL GROUP BY age, id HAVING role='root' AND NOT count<6
	fmt.Println("users9=", users9)//查出了id为6的行

	var users10 []User
	material.MyDb.Table("users").Where("id<?", 9).Group("age, id").
		Having("role=? OR count=?", "root", 10).Find(&users10) //todo:要这样写Having中的OR
	//SELECT * FROM `users` WHERE id<9 AND `users`.`deleted_at` IS NULL GROUP BY age, id HAVING role='root' OR count=10
	fmt.Println("users10=", users10) //查出了id为8、1、6、5为顺序的行
}
