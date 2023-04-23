package main
import (
	"gorm.io/gorm"
	material "guanlian/0material"
)
//todo:---------------------------------------------采用自定义外键的方式---------------------------------------------------

//has many关系，就是一个“主”有任意多个“从”。如一个购物车有多个购物项

//“主”
type User33 struct {
	gorm.Model  //下面的gorm指定了自定义的外键和references后，这里的ID作为原本gorm默认的references就不再起作用
				//但gorm还是将ID默认为primary_key和Auto_increment
	Name string
	UserCode int  `gorm:"unique"`  //int类型在建表成为字段后是长度固定的类型，故无需指定size，对看3hasOne/main.go:14
						//不能少了这里的unique，因为被参照的字段在数据表中必须是不重复的；也不能指定
						// 为primaryKey因为gorm.Model里面的ID已经被指定为primaryKey了，不然导致复合主键。
	CreditCards []CreditCard33 `gorm:"foreignKey:CardCode;references:UserCode"`
						//todo:切记，gorm中写foreignKey【而非foreign_key】,写CardCode【而非card_code】

	//todo:gorm标签意为：“从”里的字段CardCode作为外键，去存放和“主”的UserCode字段一样的内容。整个gorm标签都
	// 需写在继承字段company后面。
	// “主”中被参照的字段【references】，必须是primaryKey或unique；且和外键【foreignKey】字
	// 段的数据类型必须一致
	// 如gorm标签中不指定foreignKey，则默认是“从”结构体中的“主”的结构体名加上其primaryKey字段【如继
	// 承gorm.Model，则是加上gorm.Model中的ID】名组合成的字段作为外键，如本例中的User33ID。
	// 当gorm标签不指明references,则默认是“主”的primaryKey字段【一般定义“主”的id为primaryKey,如“主”继
	// 承了gorm.Model，则是gorm.Model内的ID】是被参照的references
}

//“从”
type CreditCard33 struct {
	gorm.Model  //gorm将ID默认为primary_key和Auto_increment
	Number string
	CardCode int  //自定义指定的外键，故不能再自主赋值
	User33ID int  //gorm指定了自定义的外键和references后，这个原本gorm默认的外键就不再起作用，得到的表格证实了这一点
}

func main() {
	material.GetDb()
	material.MyDb.AutoMigrate(&User33{},&CreditCard33{})//&CreditCard{}不能缺，不然不会创建两个表

	one := User33{
		Name: "jack",
		UserCode: 999,
		CreditCards: []CreditCard33{
			{Number: "111"},
			{Number: "222"},
		},
	}
	material.MyDb.Create(&one)
}
/*
user33表如下
+----+---------------------+---------------------+------------+------+-----------+
| id | created_at          | updated_at          | deleted_at | name | user_code |
+----+---------------------+---------------------+------------+------+-----------+
|  1 | 2020-09-12 11:54:46 | 2020-09-12 11:54:46 | NULL       | jack |       999 |
+----+---------------------+---------------------+------------+------+-----------+
credit_card33表如下
+----+---------------------+---------------------+------------+--------+-----------+----------+
| id | created_at          | updated_at          | deleted_at | number | card_code | user33_id|
+----+---------------------+---------------------+------------+--------+-----------+----------+
|  1 | 2020-09-12 11:54:46 | 2020-09-12 11:54:46 | NULL       | 111    |       999 |        0 |
|  2 | 2020-09-12 11:54:46 | 2020-09-12 11:54:46 | NULL       | 222    |       999 |        0 |
+----+---------------------+---------------------+------------+--------+-----------+----------+
*/

