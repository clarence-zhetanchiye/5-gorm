package main
import (
	"gorm.io/gorm"
	material "guanlian/0material"
)
//todo:---------------------------------------------采用gorm默认外键的方式-------------------------------------------------

//has many关系，就是一个“主”有任意多个“从”。如一个购物车有多个购物项

//“主”
type User3 struct {
	gorm.Model  //内中的ID是gorm默认的references，primary_key和Auto_increment
	Name string
	CreditCards []CreditCard3 `gorm:""`
	//这里gorm标签不填东西或根本无标签，则gorm迁移建表时默认外键是“从”结构体内的字段User3ID,被参照的reference字段
	// 是“主”中的ID【在gorm.Model里】。想自定义主键和被参照字段，相应的gorm标签均要写在[]CreditCard 后面的gorm内写
}

//“从”
type CreditCard3 struct {
	gorm.Model  //gorm迁移建表时将ID默认为primary_key和Auto_increment
	Number string
	User3ID int  //gorm迁移建表时默认的外键 故不能自主赋值
}


func main() {
	material.GetDb()
	material.MyDb.AutoMigrate(&User3{},&CreditCard3{})//&CreditCard{}不能缺，不然不会创建两个表

	one := User3{
		Name: "jack",
		CreditCards: []CreditCard3{
			{Number: "111"},
			{Number: "222"},
		},
	}
	material.MyDb.Create(&one)
}
//user_id的值和users 表的id保持一致，两行数据的user_id都是1，
//因为  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
/*
user3表如下
+----+---------------------+---------------------+------------+------+
| id | created_at          | updated_at          | deleted_at | name |
+----+---------------------+---------------------+------------+------+
|  1 | 2020-09-12 11:00:31 | 2020-09-12 11:00:31 | NULL       | jack |
+----+---------------------+---------------------+------------+------+
credit_card3表如下
+----+---------------------+---------------------+------------+--------+---------+
| id | created_at          | updated_at          | deleted_at | number | user3_id |
+----+---------------------+---------------------+------------+--------+---------+
|  1 | 2020-09-12 11:00:31 | 2020-09-12 11:00:31 | NULL       | 111    |       1 |
|  2 | 2020-09-12 11:00:31 | 2020-09-12 11:00:31 | NULL       | 222    |       1 |
+----+---------------------+---------------------+------------+--------+---------+
*/