package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//todo:由下面的示例可知，决定最终sql所针对的数据表名，
// 优先级顺序是 .Scopes() > .Table() > TableName()string方法 >  gorm语句根据finisher中的变量类型蛇形推定。

func main() {
	startGorm()
	db.Exec("DROP TABLE IF EXISTS globe")
	db.AutoMigrate(&Ball{}) //todo:下面的 TableName()string 方法就让这里的迁移建表决定了表名称。
	//CREATE TABLE `globe` (`id` bigint AUTO_INCREMENT,`name` longtext,`people` bigint,PRIMARY KEY (`id`))

	var g Ball
	db.Where("id=?", 1).Find(&g) //SELECT * FROM `globe` WHERE id=1
	fmt.Println("g=", g)

	var b Globe
	//todo:这是显示地指定表名。gorm语句中有.Table()显示指出表名是采用显示指出的，覆盖掉根据Find(&b)中的b的类型推定。
	db.Table("balls").Where("id=?", 1).Find(&b) //SELECT * FROM `balls` WHERE id=1
	fmt.Println("b=", b)

	//todo:动态地根据需要在gorm语句中指明表名。
	var gb Ball
	db.Scopes(ChangeTableName(1)).Where("id=?", 1).Find(&gb) //SELECT * FROM `balls` WHERE id=1
	fmt.Println("gb=", gb)

	var gxb Ball
	db.Scopes(ChangeTableName(2)).Where("id=?", 1).Find(&gxb) //SELECT * FROM `globe` WHERE id=1
	fmt.Println("gxb=", gxb)

	//todo:可以看到.Scopes()的优先级
	var gyb Ball
	db.Table("xx").Scopes(ChangeTableName(3)).Where("id=?", 1).Find(&gyb)
	//SELECT * FROM `circle` WHERE id=1
	fmt.Println("gyb=", gyb)

	var gzb Ball
	db.Scopes(ChangeTableName(3)).Table("yy").Where("id=?", 1).Find(&gzb)
	//SELECT * FROM `circle` WHERE id=1
	fmt.Println("gzb=", gzb)

}

type Ball struct {
	Id     int
	Name   string
	People int
}

func (b *Ball) TableName() string { //todo:这个方法决定的表会被缓存下来以便后续使用，因此这让gorm语句中自动针对的表名不支持动态变化
	return "globe"
}

//todo:在gorm语句里使用.Scopes()并将这里的函数作为入参，就可以实现gorm语句里动态地对应数据表名。
func ChangeTableName(in int) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if in == 1 {
			return tx.Table("balls")
		}
		if in == 2 {
			return tx.Table("globe")
		}
		return tx.Table("circle")
	}
}

type Globe struct {
	Id   int
	Name string
}

var db *gorm.DB
var er error

func startGorm() {
	dsn := "root:413188ok@tcp(localhost:3306)/guan1lian2gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, er = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if er != nil {
		fmt.Println("连接数据库失败:", er)
		return
	}
}
