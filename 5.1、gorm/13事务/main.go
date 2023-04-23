package main

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	material "transaction/1material"
)

func main() {
	material.GetDB()
	material.InsertGoods()
	material.GormDB.Transaction(func(tx *gorm.DB) error { //todo:点进去就可以看到Transaction里面本质上仍是手动事务。
		var g material.Good
		tx.Table("goods").Where("id=?", 1).Find(&g)
		fmt.Println("g=", g)
		return nil
	}, &sql.TxOptions{}) //todo:sql.TxOptions内有两个字段，用于指定是否只读和事务的孤立级。
}
