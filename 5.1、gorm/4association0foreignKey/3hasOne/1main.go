package main
import (
	"fmt"
	material "guanlian/0material"
)
//todo:---------------------------------------------采用自定义外键的方式---------------------------------------------------

//has one 关系是指，“主”有且仅有一个“从”，是一一对应关系。即一个购物车有且仅有一个购物项

//“主” 【相当于购物车】
type user22 struct {
	ID int     //todo:gorm将ID默认为primary_key和Auto_increment
	Address string
	Name string `gorm:"unique;size:64"`	//todo:由于要求unique和primaryKey的字段是长度固定的，故string需指定长度，
										//todo 否则建表时string类型字段会被gorm默认建成TEXT/BLOB可变长度的column。

	Card CreditCard22 `gorm:"foreignKey:Holder;references:Name"`  //todo:gorm标签意为
	//todo:gorm标签意为：“从”里的字段holder作为外键，去存放和“主”的Name字段一样的内容。整个gorm标签都需写在继
	// 承字段company后面。
	// “主”中被参照的字段【references】，必须是primaryKey或unique；且和外键【foreignKey】字
	// 段的数据类型必须一致
	// 如gorm标签中不指定foreignKey，则默认是“从”结构体中的“主”的结构体名加上其primaryKey字段【如继
	// 承gorm.Model，则是加上gorm.Model中的ID】名组合成的字段作为外键，如本例中的userID。
	// 当gorm标签不指明references,则默认是“主”的primaryKey字段【一般定义“主”的id为primaryKey,如“主”继
	// 承了gorm.Model，则是gorm.Model内的ID】被参照的references

}
//“从”  【相当于购物项】
type CreditCard22 struct{
	ID int
	Code string
	Holder string `gorm:"size:64"`//这是自定义外键，为和references字段类型一致，故也许指定长度
	UserID int  //如不指定外键，这是默认外键
}

func main() {
	material.GetDb()
	material.MyDb.AutoMigrate(&user22{},&CreditCard22{})

	one := user22{
		Name:"jack",
		Address: "america",
		Card:CreditCard22{Code: "111"},
	}
	material.MyDb.Create(&one)
	fmt.Println("运行结束")
}
/*
user22表如下：
+----+---------+------+
| id | address | name |
+----+---------+------+
|  1 | america | jack |
+----+---------+------+
credit_card22表如下：
+----+------+--------+---------+
| id | code | holder | user_id |
+----+------+--------+---------+
|  1 | 111  | jack   |       0 |
+----+------+--------+---------+
*/
