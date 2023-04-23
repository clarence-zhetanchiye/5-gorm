package material

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
)

var MyDb *gorm.DB
var err error

func GetDB() {
	dst := "root:413188ok@tcp(localhost:3306)/guan1lian2gorm?charset=utf8mb4&parseTime=True&loc=Local"
	MyDb, err = gorm.Open(mysql.Open(dst), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}

}

type Good struct{
	Id int
	Name string
	Price int
	Amount int
	Code int `gorm:"unique"`
}
var insertGoodsSql = `
INSERT INTO goods VALUES (1, 'apple', 5, 100, 111);
INSERT INTO goods VALUES (2, 'apple', 8, 200, 222);
INSERT INTO goods VALUES (3, 'banana', 3, 300, 333);
INSERT INTO goods VALUES (4, 'banana', 6, 400, 444);
INSERT INTO goods VALUES (5, 'orange', 9, 500, 555);
INSERT INTO goods VALUES (6, 'peach', 12, 600, 666)
`
func InsertGoods(){
	MyDb.Exec("DROP TABLE IF EXISTS `goods`;")
	MyDb.AutoMigrate(&Good{})
	for _, v := range strings.Split(insertGoodsSql, ";") {
		MyDb.Exec(v)
	}
}