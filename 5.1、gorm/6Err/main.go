package main

import (
	material "err/0material"
	"fmt"
)

func main() {
	//todo:下面的四个示例可以证明-------------------------------------------------------------------------------------------
	// 仅当First()、Last()、Take()方法找不到记录时，GORM会返回ErrRecordNotFound错误（jinzhu版的gorm则不是这样）。
	//可以这样来判断返回的错误是否是该错误err :
	//errors.Is(err, gorm.ErrRecordNotFound)

	material.GetDB()

	var ss []material.Soft
	result := material.MyDb.Table("softs").Where("id=?", 3).Find(&ss)
	//SELECT * FROM `softs` WHERE id=3 AND `softs`.`delete` IS NULL
	fmt.Println("result.Err=", result.Error)//result.Err= <nil>
	fmt.Println("ss=", ss)//ss= []

	var s material.Soft
	res := material.MyDb.Table("softs").Where("id=?", 3).Find(&s)
	//SELECT * FROM `softs` WHERE id=3 AND `softs`.`delete` IS NULL
	fmt.Println("res.Err=", res.Error)//res.Err= <nil>
	fmt.Println("s=", s)//s= {0  0  {0001-01-01 00:00:00 +0000 UTC false}}

	var sx material.Soft
	r := material.MyDb.Table("softs").Where("id=?", 3).First(&sx)
	//SELECT * FROM `softs` WHERE id=3 AND `softs`.`delete` IS NULL ORDER BY `softs`.`id` LIMIT 1
	fmt.Println("r.Err=", r.Error)//r.Err= record not found
	fmt.Println("sx=", sx)//sx= {0  0  {0001-01-01 00:00:00 +0000 UTC false}}

	var sl material.Soft
	rs := material.MyDb.Table("softs").Where("id=?", 3).Order("softs.id").Limit(1).Find(&sl)
	//SELECT * FROM `softs` WHERE id=3 AND `softs`.`delete` IS NULL ORDER BY softs.id LIMIT 1
	fmt.Println("rs.Err=", rs.Error)//rs.Err= <nil>
	fmt.Println("sl=", sl)//sl= {0  0  {0001-01-01 00:00:00 +0000 UTC false}}



	//todo:gorm中申明的报错有如下这些：
	//ErrRecordNotFound
	//ErrInvalidTransaction
	//ErrNotImplemented
	//ErrMissingWhereClause
	//ErrUnsupportedRelation
	//ErrPrimaryKeyRequired
	//ErrModelValueRequired
	//ErrInvalidData
	//ErrUnsupportedDriver
	//ErrRegistered
	//ErrInvalidField
	//ErrEmptySlice
	//ErrDryRunModeUnsupported
	//ErrInvalidDB
	//ErrInvalidValue
	//ErrInvalidValueOfLength
	//ErrPreloadNotAllowed

}
