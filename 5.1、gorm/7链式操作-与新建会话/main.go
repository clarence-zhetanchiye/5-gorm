package main

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	material "link/1material"
)

func main() {
	material.GetDB()
	material.InsertGoods()

	linking := material.MyDb //todo:这里的linking仍是初始会话，因此下面两个gorm语句中的链式方法Where()条件不会互相干扰。
	var aps []material.Good
	linking.Where("name=?", "apple").Find(&aps)//SELECT * FROM `goods` WHERE name='apple'
	fmt.Println("aps=", aps)//aps= [{1 apple 5 100 111} {2 apple 8 200 222}]

	var bas []material.Good
	linking.Where("name=?", "banana").Find(&bas)//SELECT * FROM `goods` WHERE name='banana'
	fmt.Println("bas=", bas)//bas= [{3 banana 3 300 333} {4 banana 6 400 444}]



	link := material.MyDb.Table("Goods") //todo:link因有链式方法Table()已不再是初始会话！故下面两个gorm语句的Where()条件会互相干扰
	var apples []material.Good
	link.Where("name=?", "apple").Find(&apples)//SELECT * FROM `Goods` WHERE name='apple'
	fmt.Println("apples=", apples)//apples= [{1 apple 5 100 111} {2 apple 8 200 222}]

	var appleBanana []material.Good
	link.Where("name=?", "banana").Find(&appleBanana)
	//SELECT * FROM `Goods` WHERE name='apple' AND name='banana' //todo:注意看这里的 name='apple' AND
	fmt.Println("appleBanana=", appleBanana)//appleBanana= []


	//todo:实际生产中，常根据不同的if条件来拼接不同的链式方法，因此实际生产中一般不会产生下面这种情况。
	//todo:下一行的lik此时已经不再是初始会话实例！因此下面两个gorm语句的链式方法Where()条件会互相干扰
	lik := material.MyDb.Table("Goods").Where("id<?", 6)
	var gs []material.Good
	lik.Where("name=?", "apple").Find(&gs)
	//SELECT * FROM `Goods` WHERE id<6 AND name='apple'
	fmt.Println("gs=", gs) //gs= [{1 apple 5 100 111} {2 apple 8 200 222}]

	var g material.Good
	lik.Where("name=?", "banana").Find(&g)
	//SELECT * FROM `Goods` WHERE id<6 AND name='apple' AND name='banana' //todo:看这里的 name='apple' AND
	fmt.Println("g=", g) //g= {0  0 0 0}



	//todo:下面的newSession是新建会话，因此下面两个gorm语句的链式方法Where()条件不会互相影响，而新建会话前的Where()条件会被两个gorm语句复用。
	newSession := material.MyDb.Where("id<?", 6).Session(&gorm.Session{})
	var ns []material.Good
	newSession.Where("name=?", "apple").Find(&ns)//SELECT * FROM `goods` WHERE id<6 AND name='apple'
	fmt.Println("ns=", ns)//ns= [{1 apple 5 100 111} {2 apple 8 200 222}]

	var nx material.Good
	newSession.Where("name=?", "banana").Find(&nx)//SELECT * FROM `goods` WHERE id<6 AND name='banana'
	fmt.Println("nx=", nx)//nx= {3 banana 3 300 333}


	//todo:下面的newSess是新建会话，因此下面两个gorm语句的Where()条件不会互相影响，而新建会话前的Where()条件会被两个gorm语句复用。
	newSess := material.MyDb.Where("id<?", 6).WithContext(context.Background())
	var ms []material.Good
	newSess.Where("name=?", "apple").Find(&ms)//SELECT * FROM `goods` WHERE id<6 AND name='apple'
	fmt.Println("ms=", ms)//ms= [{1 apple 5 100 111} {2 apple 8 200 222}]

	var mx material.Good
	newSess.Where("name=?", "banana").Find(&mx)// SELECT * FROM `goods` WHERE id<6 AND name='banana'
	fmt.Println("mx=", mx)//mx= {3 banana 3 300 333}

	//todo:由于WithContext()后面又写了一个链式方法，下面的newLast不再是新建会话，因此下面两个gorm语句的链式方法Where()条件会互相影响
	newLast := material.MyDb.WithContext(context.Background()).Where("id<?", 6)
	var msq []material.Good
	newLast.Where("name=?", "apple").Find(&msq)//SELECT * FROM `goods` WHERE id<6 AND name='apple'
	fmt.Println("msq=", msq)//msq= [{1 apple 5 100 111} {2 apple 8 200 222}]

	var mxq material.Good
	newLast.Where("name=?", "banana").Find(&mxq)//SELECT * FROM `goods` WHERE id<6 AND name='apple' AND name='banana'
	fmt.Println("mxq=", mxq)//mxq= {0  0 0 0

	//todo:更多参见 0a说明/0链式调用.txt
}
