package cha

import (
	material "crud/0material"
	"fmt"
)

//todo:--------------------------------------------------------OR-------------------------------------------------------
//根据其他字段查询。 where 包含or  or里面的参数填写方式和where一样
// 不建议用这个，建议就用read2中的fas8就可以替代
func FindWhereOr() {
	var facts []factory
	material.MyDb.Where("output=?",1).Or("name=?","sanliu2").Find(&facts)
	//SELECT * FROM `factories` WHERE output=1 OR name='sanliu2'
	fmt.Println("facts1=",facts)//facts1= 查出了id为1、2这两条记录。

	facts = []factory{}
	material.MyDb.Where("output=?",1).Or(factory{Name:"sanliu2",Output: 2}).Find(&facts)
	//SELECT * FROM `factories` WHERE output=1 OR (`factories`.`name` = 'sanliu2' AND `factories`.`output` = 2)
	fmt.Println("facts2=",facts)//facts2= 查出了id为1、2这两条记录。

	facts = []factory{}
	material.MyDb.Where("output=?", 1).Or(map[string]string{"name":"sanliu2", "product":""}).
		Find(&facts) //SELECT * FROM `factories` WHERE output=1 OR (`name` = 'sanliu2' AND `product` = '')
	fmt.Println("facts3=",facts) //facts3= 查出了id为1、2这两条记录。
}