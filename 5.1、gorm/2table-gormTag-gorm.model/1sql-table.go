//todo--------------------------标准库db连接用sql语句建表----------------------------------
//            （实际生产中是项目开始前在一个.sql文件中写sql建表语句，导入MySQL完成建表）

package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
func main(){
	var dsn3 = "root:413188ok@tcp(localhost:3306)/bo2ke4?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn3)
	if err != nil {
		fmt.Println("err=",err)
		return
	}

	//todo:下面的sql语句中id就不能省略auto_increment，否则sql2中不写id的值就会报错；
	// 即不会因为id是主键而默认自增，自增是需要专门指出来的

	sql1 := `create table if not exists new_table(
			id int primary key auto_increment,
			name varchar(10))`
	_,err = db.Exec(sql1)
	if err != nil {
		fmt.Println("err=",err)
		return
	}
	sql2 := "insert into new_table (name) values('jack')"
	exec, err := db.Exec(sql2)
	if err != nil {
		fmt.Println("err27=",err)
		return
	}
	res, err := exec.RowsAffected()
	fmt.Printf("exec=%v,%v",res,err)
	fmt.Println("end")
}
/*
+----+------+
| id | name |
+----+------+
|  1 | jack |
+----+------+
*/