package main
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
//todo---------------------- gorm迁移建表后插入数据失败（实际生产不是采用迁移建表的）---------------------------------
//字段和标签为 CommentID int `gorm:"type:bigint(20);primaryKey;autoIncrement"`运行后报错：
//Error 1062: Duplicate entry '0' for key 'PRIMARY'
//加上 autoIncrement或autoIncrement:true标签后仍报错：
//failed to parse  as default value for int, got error: strconv.ParseInt: parsing "": invalid syntax

var dsn2 = "root:413188ok@tcp(localhost:3306)/bo2ke4?charset=utf8mb4&parseTime=True&loc=Local"

type YourComment struct {
	CommentID int `gorm:"type:bigint(20);primaryKey;autoIncrement"`//todo:失败是因为本字段写了primaryKey后又写autoIncrement
	Sentence string `gorm:"type:varchar(50)"`
}

func main(){
	db,err := gorm.Open(mysql.Open(dsn2),&gorm.Config{})
	if err != nil {
		fmt.Println("err=",err)
	}
	db.Exec("drop table if exists your_comments")
	err = db.AutoMigrate(&YourComment{}) //todo:见表成功。但实际建的表如末尾所示，comment_id有个DEFAULT`0`
	if err != nil {
		fmt.Println("迁移err=", err)
		return
	}
	db.Create(&YourComment{Sentence:"hello"}) //todo:这一行插入成功，之后的失败，因为CommentId的值又是默认的0，重复了。
	db.Create(&YourComment{Sentence:","})
	db.Create(&YourComment{Sentence:"World"})
	fmt.Println("结束")
}
/*
CREATE TABLE `your_comments` (
  `comment_id` bigint(20) NOT NULL DEFAULT '0',
  `sentence` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`comment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
*/