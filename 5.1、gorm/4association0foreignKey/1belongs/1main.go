package main
import (
	"fmt"
	material "guanlian/0material"
)
//todo:----------------------------------------------采用自定义外键的方式--------------------------------------------------

//belongs关系，就是任意数量的“从”隶属于同一个“主”，即一个购物车有任意数量的购物项

//主  【相当于购物车】
type companyTwo struct {
	ID int   //todo:原本gorm默认其为主键和自增，但由于下两行指定了主键，故变成一个普通字段
	Name string
	Code int `gorm:"primaryKey"`
}
//从 【相当于购物项】
type employeeTwo struct {
	ID int   //gorm默认为主键、自增
	Name string
	CompanyID int  //若不自定义外键，则这是gorm迁移建表时默认的外键，自定义外键后，该字段就是普通字段
	ComCode int
	Company companyTwo `gorm:"foreignKey:ComCode;references:Code"` //todo:这里的gorm标签意为：
	//todo:在gorm迁移建表时设定“从”内的ComCode字段作为外键，去存放和“主”内的Code字段一样的内容。整
	// 个gorm标签都需写在继承字段Company后面。
	// “主”中被参照的字段【references】，必须是primaryKey或unique；且和外键【foreignKey】字
	// 段的数据类型必须一致
	// 如gorm标签中不指定foreignKey和references，则默认是“从”结构体中的--“主”的结构体名加上其primaryKey字
	// 段【如继承gorm.Model，则是加上gorm.Model中的ID】名组合成的--字段作为外键，如本例中的companyID。
	// 并默认是“主”的primaryKey字段【一般定义“主”的id为primaryKey,如“主”继承了gorm.Model，则是gorm.Model
	// 内的ID】作为被参照的references字段。
}

func main() {
	material.GetDb()
	material.MyDb.Exec("DROP TABLE IF EXISTS company_twos")
	material.MyDb.Exec("DROP TABLE IF EXISTS employee_twos")
	material.MyDb.AutoMigrate(&companyTwo{},&employeeTwo{})

	e := employeeTwo{Name:"erlang",Company: companyTwo{Name:"cpGroup",Code:2333}}
	material.MyDb.Create(&e)
	fmt.Println("运行结束")
}
/*
mysql> select * from company_twos;
+------+---------+------+
| id   | name    | code |
+------+---------+------+
|    0 | cpGroup | 2333 |
+------+---------+------+
mysql> select * from employee_twos;
+----+--------+------------+----------+
| id | name   | company_id | com_code |
+----+--------+------------+----------+
|  1 | erlang |          0 |     2333 |
+----+--------+------------+----------+
*/
