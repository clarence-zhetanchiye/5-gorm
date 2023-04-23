package material

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
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

type Soft struct {
	Id   int
	Name string
	Age  int
	Food string
	Delete gorm.DeletedAt
}

func InsertSoft() {
	MyDb.Exec("DROP TABLE IF EXISTS softs;")
	MyDb.AutoMigrate(&Soft{})
	for i := 1; i <= 2; i++ {
		MyDb.Create(&Soft{Name: "tiger" + strconv.Itoa(i), Age: i, Food: "sheep"})
	}
}