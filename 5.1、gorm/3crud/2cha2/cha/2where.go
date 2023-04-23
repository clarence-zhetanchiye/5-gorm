package cha

import (
	material "crud/0material"
	"fmt"
)

//todo:------------------------------------------------------where-----------------------------------------------------
//参见https://gorm.io/zh_CN/docs/query.html的string条件]
func FindByWhere() {
	var fas []factory                             //todo:由于可能查询出多条记录故申明为切片
	material.MyDb.Where("output=?", 9).Find(&fas) //SELECT * FROM `factories` WHERE output=9
	fmt.Println("fas1=", fas)                     //fas1= [{9 sanliu9  9 one country}]

	var fa factory                               //todo:由于只会有一条记录故不用申明为切片
	material.MyDb.Where("output=?", 9).Find(&fa) //SELECT * FROM factories WHERE output=9
	fmt.Println("fa2=", fa)                      //fa2= {9 sanliu9  9}

	fas = []factory{}
	material.MyDb.Where("output>=?", 8).Find(&fas) //SELECT * FROM `factories` WHERE output>=8
	fmt.Println("fas3=", fas)                      //fas3= 查出了id为8、9、10、11、12的内容

	fas = []factory{}
	material.MyDb.Where("name LIKE ?", "%100%").Find(&fas) //当然也可以 NOT LIKE
	//SELECT * FROM `factories` WHERE name LIKE '%100%'
	fmt.Println("fas4=", fas) //fas4= [{10 sanliu100 broom 100   }]

	fas = []factory{}
	material.MyDb.Where("output>=? AND name=?", 8, "sanliu9").Find(&fas)
	//SELECT * FROM `factories` WHERE output>=8 AND name='sanliu9'
	fmt.Println("fas5=", fas) //fas5=查出了id为9的一条记录。

	fas = []factory{}
	material.MyDb.Where("output IN ?", []int{7, 8, 9}).Find(&fas) //SELECT * FROM `factories` WHERE output IN (7,8,9)
	fmt.Println("fas6=", fas)                                     //fas6= 查出了id为7、8、9这几条内容。

	fas = []factory{}
	material.MyDb.Where("name IN ?", []string{"sanliu1", "sanliu2"}).Find(&fas)
	//SELECT * FROM `factories` WHERE name IN ('sanliu1','sanliu2')
	fmt.Println("fas66=", fas) //fas66= 查出了name是'sanliu1'或'sanliu2'的几条数据。

	fas = []factory{}
	material.MyDb.Where("output BETWEEN ? AND ? ", 7, 9).Find(&fas)
	//SELECT * FROM `factories` WHERE output BETWEEN 7 AND 9
	fmt.Println("fas7=", fas) //fas7= 查出了id为7、8、9这三条内容。

	fas = []factory{}
	material.MyDb.Where("output =? OR output=? ", 7, 8).Find(&fas)
	//SELECT * FROM `factories` WHERE output =7 OR output=8
	fmt.Println("fas8=", fas) //fas8= 查出了id为7、8这两条内容

	fas = []factory{}
	material.MyDb.Where("product IS NOT NULL").Find(&fas) //SELECT * FROM `factories` WHERE product IS NOT NULL
	fmt.Println("fas9=", fas)                             //fas9=把全部都取了出来，因为数据库表中值为空白表示值是空字符串，如果是null数据库表中会填充为null。

	fas = []factory{}
	material.MyDb.Where("output > ? AND output != ?", 8, 9).Find(&fas) //!=也可以用<> 都表示不等于
	//SELECT * FROM `factories` WHERE output > 8 AND output != 9
	fmt.Println("fas10=", fas) //fas10=查出了id为10、11、12这几条记录。

	fas = nil
	material.MyDb.Find(&fas)   // SELECT * FROM `factories`
	fmt.Println("fas11=", fas) //fas11=把所有数据从数据库读取出来了。虽然fas是nil。

}

//todo：-----------------------------------------Where(采用struct或map)----不推荐使用--------------------------------------
//注意此时Where()里struct内的零值的字段不会加入到sql语句的where条件里！！！！
//而Where()里map内的零值的字段则是会加入到sql语句的where条件里的。
func FindByWhereStruct() {
	var fac factory //因为知道只能查到一条记录，故不申明切片，即使下面用的Find
	material.MyDb.Where(&factory{Name: "sanliu1", Output: 1}).Find(&fac)
	//SELECT * FROM `factories` WHERE `factories`.`name` = 'sanliu1' AND `factories`.`output` = 1
	fmt.Println("fac=", fac) //fac= {1 sanliu1  1 two country}

	fac = factory{}
	material.MyDb.Where(&factory{Name: "sanliu1", Output: 1}).First(&fac)
	//SELECT * FROM `factories` WHERE `factories`.`name` = 'sanliu1' AND `factories`.`output` = 1
	//ORDER BY `factories`.`id` LIMIT 1
	fmt.Println("fac12=", fac) //fac12= {1 sanliu1  1 two country}

	//不建议用下面这个,直接用fas6就行
	var facs []factory
	material.MyDb.Where([]int64{1, 2, 3}).Find(&facs) //SELECT * FROM `factories` WHERE `factories`.`id` IN (1,2,3)
	fmt.Println("fac13=", facs)//fac13=查出了id为1、2、3这三条记录。

	//也不建议使用下面这样的。
	//还有Where里是map的，此时map里的零值字段是会被加入到sql语句中的where条件的。
	fac = factory{}
	material.MyDb.Where(map[string]interface{}{"name": "sanliu1", "product": ""}).
		Find(&fac) //SELECT * FROM `factories` WHERE `name` = 'sanliu1' AND `product` = '';
	fmt.Println("fac14=", fac) //fac14= {1 sanliu1  1 two country}
}
