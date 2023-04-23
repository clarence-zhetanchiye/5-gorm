package cha

import (
	material "crud/0material"
	"fmt"
)

//todo---------------------------------------------------一般不用下面这两种方式-----------------------------------------------

//根据主键查询------------------------------------------------Find--------------------------------------------------------
// 主键最好是整型，若是string则需注意SQL注入的安全问题。
func FindById() {
	//todo:Find()是查询符合条件的数据库表记录的所有条
	// Find()不给查询条件参数，就是查询所有记录
	var fs2 []factory
	material.MyDb.Find(&fs2)//SELECT * FROM `factories`
	fmt.Println("fs1=",fs2)//fs1=全部查询出来了

	var f factory
	res := material.MyDb.Find(&f, 5)// SELECT * FROM `factories` WHERE `factories`.`id` = 5
	fmt.Println("err=",res.Error)//err= <nil>
	fmt.Println("f2=",f)//f2= {5 sanliu5  5 three country}

	var fs []factory
	res = material.MyDb.Find(&fs,[]int{7,8,9})//SELECT * FROM `factories` WHERE `factories`.`id` IN (7,8,9)
	fmt.Println("res3=",res.RowsAffected)//res3= 3
	fmt.Println("fs3=",fs)//fs3=查出了id为7、8、9这几条记录。


	//todo:First()是查询符合条件的数据库表记录的第一条。
	// First不给查询参数，就是按主键[有主键时]或第一个字段[无主键时]的升序排序后查询第一条
	var f2 factory
	material.MyDb.First(&f2)//SELECT * FROM `factories` ORDER BY `factories`.`id` LIMIT 1
	fmt.Println("f4=",f2)//f4=查出了id为1的一条数据。

	var f0 factory
	//下一行的material.MyDb指明了是从哪个数据库查询，&f0中f0的类型指明了是从哪个数据库的表查询
	res = material.MyDb.First(&f0, 3) //SELECT * FROM `factories` WHERE `factories`.`id` = 3 ORDER BY `factories`.`id` LIMIT 1
	fmt.Println("res1=",res.RowsAffected)//res1= 1
	fmt.Println("f5=",f0)//f5=查出了id为3的一条记录。
}
func FindByCondition(){
	var f11 []factory
	material.MyDb.Find(&f11,"output=?",2)//SELECT * FROM `factories` WHERE output=2
	fmt.Println("f6=",f11) //查出了id为2这一条记录

	material.MyDb.Find(&f11,"output IN ?",[]int{1,2})//SELECT * FROM `factories` WHERE output IN (1,2)
	fmt.Println("f7=",f11) //查出了id为1、2这两条记录
}