package main
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
//todo:-----------------成功的gorm迁移建表（实际生产不是采用迁移建表）-------------------------
//主键的特性是  非空、不重复、一个表若有主键则只能有一个
// 不同于sql语句需要在指定primary key之外同时指定auto_increment，
// gorm会把gorm标签写了primaryKey的字段作为主键的同时默认自增，当新增一个结构体变量进入数据库表，
// 而该结构体变量的主键字段未赋值（即为零值），gorm会默认从1（含1）自增，以符合非空、不重复的特性。
// 如果主键字段赋值了，符合主键字段类型则会被采纳，其后新增的结构体如果主键未赋值则以上一条
// 记录的主键为基础继续自增
// 就像gorm自带的Model结构体
// type Model struct {
//  ID        uint           `gorm:"primaryKey"`
//  CreatedAt time.Time
//  UpdatedAt time.Time
//  DeletedAt gorm.DeletedAt `gorm:"index"`
// }
// 中gorm对待ID那样
// 而且切记不能再在primaryKey标签之外写其他标签（即不能写`gorm:"primaryKey;autoIncrement"`），
// 否则不能纯粹地实现主键同时自增的效果


var dsn3 = "root:413188ok@tcp(localhost:3306)/bo2ke4?charset=utf8mb4&parseTime=True&loc=Local"
type HerComment struct {
	CommentID int `gorm:"primaryKey"`
	Sentence string `gorm:"type:varchar(50)"`
}

func main(){
	db,err := gorm.Open(mysql.Open(dsn3),&gorm.Config{DisableForeignKeyConstraintWhenMigrating:true})
	if err != nil {
		fmt.Println("err=",err)
	}
	db.Exec("drop table if exists her_comments")
	db.AutoMigrate(&HerComment{}) //todo:建表成功，建的表如末尾所示。

	db.Create(&HerComment{CommentID:0,Sentence: "hello"}) //这里赋0也没影响
	db.Create(&HerComment{CommentID:0,Sentence: ","}) //这里赋0也没影响
	db.Create(&HerComment{CommentID:5,Sentence: "World"})
	db.Create(&HerComment{Sentence: "!"})
	db.Create(&HerComment{CommentID:-3,Sentence: "ha"})

	/*运行结果如下：
	+------------+----------+
	| comment_id | sentence |
	+------------+----------+
	|         -3 | ha       |
	|          1 | hello    |
	|          2 | ,        |
	|          5 | World    |
	|          6 | !        |
	+------------+----------+
	*/

	fmt.Println("结束")
}

/*
CREATE TABLE `her_comments` (
  `comment_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `sentence` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`comment_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;
*/
