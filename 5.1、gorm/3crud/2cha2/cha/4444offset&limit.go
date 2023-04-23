package cha

import (
	material "crud/0material"
	"fmt"
)

//todo:--------------------------------------------------LIMIT OFFSET分页------------------------------------------------
// LIMIT表示选出连续的几条，OFFSET表示跳过几条，sql语句中可以只有LIMIT，但不能只有OFFSET，sql语句中LIMIT在先OFFSET在后。
// 更多可见官网 GORM/CRUD/查询/Limit & Offset
func Limit() {
	//todo:数据表中既有顺序分页显示
	var fy []factory
	material.MyDb.Offset(5).Limit(2).Find(&fy)
	//SELECT * FROM `factories` LIMIT 2 OFFSET 5; 也即SELECT * FROM factories LIMIT 5,2 即数据表的既有顺序下选择第6、7条
	fmt.Println("fy1=",fy)//fy1= [{6 sanliu6  6 one people} {7 sanliu7  7 two country}]

	//todo:对Where查询结果排序后分页显示
	fy = []factory{}
	material.MyDb.Where("output>?",4).Order("id DESC").Offset(1).Limit(2).Find(&fy)
	//SELECT * FROM `factories` WHERE output>4 ORDER BY id DESC LIMIT 2 OFFSET 1
	fmt.Println("fy2=",fy)//fy2= [{11 wuqi999 brush 999  } {10 sanliu100 broom 100  }]


	//todo:还有下面这一类不常见的用法。
	var users1, users2 []User
	material.MyDb.Limit(2).Find(&users1).Limit(-1).Where("name=?", "plus").Find(&users2)
	//这里的.Limit(-1)取消了.Limit(10)这个条件，因此这一句gorm语句其实产生并执行了两条sql语句。
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL LIMIT 2;  这是对于user1
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND name='plus';  这是对于users2
	fmt.Println("users1=", users1) //users1=查出了id为1、2为顺序的这两条数据。
	fmt.Println("users2=", users2) //users2=查出了id为9的一条数据。

	var usersx []User
	material.MyDb.Limit(2).Find(&User{}).Limit(-1).Limit(1).Find(&usersx)
	//SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL LIMIT 2
	//SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL LIMIT 1
	fmt.Println("usersx=", usersx) //usersx=查出了id为1的一条数据。

	//类似的，offset也有
	var users3, users4 []User
	material.MyDb.Limit(1).Offset(7).Find(&users3).Offset(-1).Offset(8).Find(&users4)
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL LIMIT 1 OFFSET 7; (users3)
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL LIMIT 1 OFFSET 8; (users4)
	fmt.Println("users3=", users3)//users3=查出了id为8的一条数据
	fmt.Println("users4=", users4)//users4=查出了id为9的一条数据


	//另外有下面这个分页器的实现，参见 GORM文档/教程/Scope/分页
/*
   func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
     return func (db *gorm.DB) *gorm.DB {
       q := r.URL.Query()
       page, _ := strconv.Atoi(q.Get("page"))
       if page == 0 {
         page = 1
       }

       pageSize, _ := strconv.Atoi(q.Get("page_size"))
       switch {
       case pageSize > 100:
         pageSize = 100
       case pageSize <= 0:
         pageSize = 10
       }

       offset := (page - 1) * pageSize
       return db.Offset(offset).Limit(pageSize)
     }
   }

   db.Scopes(Paginate(r)).Find(&users) //上面Paginate中的.Offset(..).Limit(..)就会写进这个gorm语句里。
   db.Scopes(Paginate(r)).Find(&articles) //上面Paginate中的.Offset(..).Limit(..)就会写进这个gorm语句里。
*/
}