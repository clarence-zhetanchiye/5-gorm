//todo--------事先使用sql建表语句进行建表时，要配合gorm默认的蛇形对应规则，以让gorm语句可成功进行crud--------------

package main
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MyComment struct {
	CommentID uint64 //todo:不写gorm标签指定与表列的对应关系，而是先建表并在建表时配合gorm的默认蛇形对应规则。
	Sentence string
	gorm.Model
}
func main(){
	var dsn = "root:413188ok@tcp(localhost:3306)/bo2ke4?charset=utf8mb4&parseTime=True&loc=Local"
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println("err=",err)
	}

	//todo 手动建表时表、列命名要对应上gorm的蛇形规则，才能让gorm语句的crud能成功的用。

	db.Exec(`create table my_comments (
	comment_id bigint(10) primary key auto_increment,   
	sentence varchar(400) not null)`) //todo 用非GORM默认的ID字段作为自增主键，以下成功运行

	err2 := db.Create(&MyComment{Sentence:"hello"}).Error
	if err != nil {
		fmt.Println("err2=",err2)
	}
	db.Create(&MyComment{Sentence:"hey"})
	err3 := db.Create(&MyComment{CommentID:3,Sentence:"hi"}).Error
	if err3 != nil {
		fmt.Println("err3=",err3)
	}
	/*
	+------------+----------+
	| comment_id | sentence |
	+------------+----------+
	|          1 | hello    |
	|          2 | hey      |
	|          3 | hi       |
	+------------+----------+
	*/
	fmt.Println("end结束")
}
