package cha

import (
	material "crud/0material"
	"fmt"
	"gorm.io/gorm/clause"
)

//todo:-----------------------------------------------------Order排序---------------------------------------------------
// 对查询结果排序地显示，ASC升序，DESC降序，默认升序
func Order() {
	//todo:单排序
	var fctr []factory
	material.MyDb.Where("output>?",10).Order("output DESC").Find(&fctr)
	//SELECT * FROM `factories` WHERE output>10 ORDER BY output DESC
	fmt.Println("fctr=",fctr)//fctr= [{11 wuqi999 brush 999  } {10 sanliu100 broom 100  } {12 wuqi88 tank 88  }]


	//todo:双排序
	var users []User
	material.MyDb.Where("id<?", 7).Order("age, count DESC").Find(&users)
	//SELECT * FROM `users` WHERE id<7 AND `users`.`deleted_at` IS NULL ORDER BY age, count DESC
	fmt.Println("users1=", users)//users1=查出了id为4、2、6、1、5、3为顺序的这几条数据。
	//从结果看，可以视为第一个order依据字段是分组且组间排序的依据，第二个order字段是组内排序的依据字段。

	//todo:双排序也可以这样写
	users = []User{}
	material.MyDb.Where("id<?", 7).Order("age").Order("count DESC").Find(&users)
	//SELECT * FROM `users` WHERE id<7 AND `users`.`deleted_at` IS NULL ORDER BY age, count DESC
	fmt.Println("users2=", users) //users2=查出了id为4、2、6、1、5、3为顺序的这几条数据。


	//todo:还有这样的用法。
	users = []User{}
	material.MyDb.Where("id>?", 6).Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{9, 8, 7}}, WithoutParentheses: true},
	}).Find(&users)
	//SELECT * FROM `users` WHERE id>6 AND `users`.`deleted_at` IS NULL ORDER BY FIELD(id,9,8,7)
	fmt.Println("users3=", users) //users3=查出了id为9、8、7为顺序的这几条数据。
}