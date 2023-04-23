package main

import (
	"fmt"
	"gorm.io/gorm"
	material "scopes/1material"
)
//todo:由下可知 .Scopes(...) 的可变入参所构成的sql语句中的条件之间的关系是AND---------------------------------------------------

func main() {
	material.GetDB()
	material.InsertGoods()

	var g material.Good
	//todo:点进.Scopes源码，猜测可知MyDb链式调用Scopes(fx)方法时，会将MyDb自己作为参数传给Scopes(fx)的fx的入参
	material.MyDb.Table("goods").Scopes(smallSeven).Where("id=?", 1).Find(&g)
	//SELECT * FROM `goods` WHERE id=1 AND id<7
	fmt.Println("g=", g)//g= {1 apple 5 100 111}

	var gs []material.Good
	material.MyDb.Table("goods").Scopes(smallSeven, apple).Order("id DESC").Find(&gs)
	//SELECT * FROM `goods` WHERE id<7 AND name='apple' ORDER BY id DESC
	fmt.Println("gs=", gs)//gs= [{2 apple 8 200 222} {1 apple 5 100 111}]

	gs = nil
	material.MyDb.Table("goods").Scopes(smallSeven, judgePart(2)).Limit(3).Find(&gs)
	//SELECT * FROM `goods` WHERE id<7 AND name IN ('apple','banana') LIMIT 3
	fmt.Println("gs2=", gs)//gs2= [{1 apple 5 100 111} {2 apple 8 200 222} {3 banana 3 300 333}]

	gs = nil
	material.MyDb.Table("goods").Scopes(smallSeven, judgePart(0)).Where("name=?", "orange").Find(&gs)
	//SELECT * FROM `goods` WHERE name='orange' AND id<7
	fmt.Println("gs3=", gs) //gs3= [{5 orange 9 500 555}]

	var x material.Good
	material.MyDb.Table("goods").Scopes(smallSeven, small7Apple).Find(&x)
	//SELECT * FROM `goods` WHERE id<7 AND name='apple' AND id<7
	fmt.Println("x=", x)//x= {1 apple 5 100 111}
}

func smallSeven(db *gorm.DB) *gorm.DB {
	return db.Where("id<?", 7)
}
func apple(db *gorm.DB) *gorm.DB {
	return db.Where("name=?", "apple")
}
func judgePart(in int) func (db *gorm.DB)*gorm.DB {
	if in == 2 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("name IN (?)", []string{"apple", "banana"})
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
func small7Apple(db *gorm.DB) *gorm.DB {
	return db.Scopes(smallSeven).Where("name=?", "apple")
}

