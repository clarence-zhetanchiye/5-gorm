package cha

import (
	material "crud/0material"
	"fmt"
)

//todo:--------------------------------------------------DISTINCT去重----------------------------------------------------

//todo:注意, 如果sql语句中要用DISTINCT关键词, 它就必须紧跟SELECT并写在所有字段前面,
// 即必须是这样 SELECT DISTINCT column1, column2··· FROM table1
// 而不能像类似这样 SELECT column1, DISTINCT column2, column3 FROM···来写, 这样写sql会报语法错误
// 也就是说,
// 不能这样写来企图实现去重的同时又连带查出去重后相应行的其他字段, 要想这样实现, 只能不这样用, 而改为
// 使用GROUP BY, 即把去重依据字段改成作为分组的依据字段（参考：https://www.jianshu.com/p/0f86250967c6）,
// 但即使是这样也只适合DISTINCT一个字段的情况改为Group BY。
// 【sql语句中诸如DISTINCT(column1)或DISTINCT(column1, column2)或(DISTINCT column1)这样带括号的写法都可以认为是错误或没必要的】。
func Distinct() {
	var users []User
	material.MyDb.Distinct("age", "role").Order("id").Find(&users)
	//SELECT DISTINCT `age`,`role` FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY id
	fmt.Println("users=", users)//按序显示了 12 root; 11 user; 30 user; 30 root 这四条数据，切片中每个结构体的其他字段是零值。

	var users2 []User
	material.MyDb.Select("DISTINCT age, role").Order("id").Find(&users2)
	//SELECT DISTINCT age, role FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY id
	fmt.Println("users2=", users2)//按序显示了 12 root; 11 user; 30 user; 30 root 这四条数据，切片中每个结构体的其他字段是零值。

	wrong()//这里面的gorm语句写法是不对的
}

func wrong() {
	var users3 []User
	material.MyDb.Select("id").Distinct("age", "role").Order("id").Find(&users3)
	//不会报sql语法错误但没有实现意图。
	//SELECT DISTINCT `age`,`role` FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY id;
	fmt.Println("users3=", users3)//按序显示了 12 root; 11 user; 30 user; 30 root 这四条数据，切片中每个结构体的其他字段是零值。


	var users4 []User
	material.MyDb.Select("id, DISTINCT age, role").Order("id").Find(&users4)
	//SELECT id, DISTINCT age, role FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY id  //这个接下来报出sql语法错误
	fmt.Println("结果users4=", users4)//结果users4= []
}