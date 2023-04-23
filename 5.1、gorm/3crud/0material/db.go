package material

import(
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

