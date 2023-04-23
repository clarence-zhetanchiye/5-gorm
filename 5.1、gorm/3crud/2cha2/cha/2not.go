package cha

import (
	material "crud/0material"
	"fmt"
)

//todo:------------------------------------------------------Not--------------------------------------------------------
// not的性质和where是一样的，因此not和where也可以一起用
func FindWhereNot() {
	var factos []factory
	material.MyDb.Not("output<?", 9).Find(&factos)
	//SELECT * FROM `factories` WHERE NOT output<9
	fmt.Println("factos1=", factos) //factos1=查出了id为9、10、11、12这几条数据

	factos = []factory{}
	material.MyDb.Not("name NOT LIKE ?", "%wuqi%").Find(&factos)
	//SELECT * FROM `factories` WHERE NOT name NOT LIKE '%wuqi%'
	fmt.Println("factos2=", factos) //factos2= [{11 wuqi999 brush 999  } {12 wuqi88 tank 88  }]

	factos = []factory{}
	material.MyDb.Not("id in ?", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}).Find(&factos)
	//SELECT * FROM `factories` WHERE NOT id in (1,2,3,4,5,6,7,8,9)
	fmt.Println("factos3=", factos) //factos3=查出了id为10、11、12

	factos = []factory{}
	material.MyDb.Not("name IN ?", []string{"sanliu1", "sanliu2", "sanliu3", "sanliu4", "sanliu5"}).Find(&factos)
	//SELECT * FROM `factories` WHERE NOT name IN ('sanliu1','sanliu2','sanliu3','sanliu4','sanliu5')
	fmt.Println("factos4=", factos) //factos4=查出了id为6、7、8、9、10、11、12的这些记录

	//不建议用，factos3已经能实现
	//factos = []factory{}
	//material.MyDb.Not([]int{1,2,3}).Find(&factos)
	////SELECT * FROM factories WHERE id NOT IN (1,2,3)
	//fmt.Println("factos5=",factos)

	factos = []factory{}
	material.MyDb.Not(factory{Name: "sanliu1", Output: 2}).Find(&factos) //如果是.First(&factos)就只查询符合条件的排序后的第一条
	//OM `factories` WHERE (`factories`.`name` <> 'sanliu1' AND `factories`.`output` <> 2)//<>是 != 的意思
	fmt.Println("factos6=", factos) //factos6=查出了id为3、4、5、6、7、8、9、10、11、12这些记录。

	factos = []factory{}
	material.MyDb.Not(map[string]string{"product": ""}).Find(&factos)
	//SELECT * FROM `factories` WHERE `product` <> ''
	fmt.Println("factos7=", factos)//factos7=查出了id为10、11、12这几条记录。
}
