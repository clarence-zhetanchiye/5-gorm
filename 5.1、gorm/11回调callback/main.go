//-----------------------------成功运行-------------------------
package main
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type student struct {
	ID int
	Name string
	Age int
}
//func (s *student)BeforeCreate(tx *gorm.DB)error {
//	fmt.Println("我的钩子")
//	s.Name = "mary"
//	s.Age = 18
//	//tx.Create(&s)//todo:会引起递归，这里死归了
//	return nil
//}
func (s *student)BeforeCreate(tx *gorm.DB)error {
	fmt.Println("我的钩子")
	s.Name = "mary"
	s.Age = 18  //todo:看查询出来的表格情况，会发现钩子这里的操作更改了45行的新增
	return nil
}
var (
	db *gorm.DB
	err error
)
func init() {
	dstn := "root:413188ok@tcp(localhost:3306)/guan1lian2gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dstn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库出错：",err)
		return
	}
	db.Callback().Create().Register("myCallBack",MyCallBack())
	db.Callback().Create().Before("myCallBack").Register("myCallBackTwo",MyCallBackTwo)
	//todo:鼠标悬浮每一个单词逐步深入去看，可以看到提供的连续若干函数
}
func main() {
	db.AutoMigrate(&student{})

	st := student{Name:"jack",Age: 27}
	db.Create(&st)		//todo:执行到此处 st下Create类的钩子会自动被调用
	fmt.Println("运行完毕")

}

func MyCallBack() func(*gorm.DB) {
	return func(tx *gorm.DB) {
		fmt.Println("我是CallBack")
		type teacher struct{
			MyId int
			Call string
		}
		//tx.AutoMigrate(&teacher{})  //报错 ALTER TABLE `students` ADD `my_id` bigint 同时还在students表里新增了my_id和call两列
		//tx.Create(&teacher{Call:"haha"})//这会导致又触发回调，递归死归了

//todo 手动通过海豚创建名为teachers的表并写入 1 ha 和  2 hey  两条数据
 //		    var t teacher
//				//todo 下面一行运行后提示说 1mrecord not found  而且发现导致了students表的数据插入了两条
//			tx.Table("teachers").Where("my_id=?",1).First(&t)
//			fmt.Println("t=",t)  //todo:t={0  }  故这样做有问题，大概因为tx并不是一个能独立自主的db

		var s teacher
		sqlDB, err2 := tx.DB()
		if err2 != nil {
			fmt.Println("err2=",err2)
			return
		}
		row := sqlDB.QueryRow("select * from teachers where my_id =?", 2)
		row.Scan(&s.MyId,&s.Call)
		fmt.Println("s=",s)
	}
}
func MyCallBackTwo(tx *gorm.DB){
	fmt.Println("我是CallBack2")
	var ss string
	sqlDb, _ := tx.DB()
	row := sqlDb.QueryRow("select call from teachers where my_id =?", 2)
	row.Scan(&ss)
	fmt.Println("ss=",ss)
}
/*
我的钩子
我是CallBack2
ss=
我是CallBack
s= {2 hey}
运行完毕
mysql> select * from students;
+----+------+------+
| id | name | age  |
+----+------+------+
|  1 | mary |   18 |
+----+------+------+
*/
