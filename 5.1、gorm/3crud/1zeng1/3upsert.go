//todo: MySQL中Upsert的意思：通过判断即将插入的一行记录与表中已存在的所有行数据是否存在主键索引或unique索引冲突，
// 来决定是插入还是更新。当出现主键索引或唯一索引冲突时则进行update操作，否则进行insert操作。
// gorm实现MySQL的Upsert功能时，gorm语句里的.Clauses(clause.OnConflict{})中入参的clause.OnConflict{}结构体中只有字
// 段DoUpdates和UpdateAll是真正有用的，其他字段都没有用(除非gorm用于其他数据库)。

package main

import (
	material "crud/0material"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	gorm.Model
	Name string
	Age int
	Role string
	Count int `gorm:"unique"` //todo:gorm.Model中的ID是主键、这里的Count是唯一索引，因此这两个字段是会被用来判断新插入的数据是否冲突的。
}

func main() {
	material.GetDB()
	if err := material.MyDb.AutoMigrate(&User{}); err != nil {
		fmt.Println("自动迁移建表出错=", err)
		return
	}
	if err := material.MyDb.Exec("TRUNCATE TABLE users").Error; err != nil {
		fmt.Println("截断清空表数据出错=", err)
		return
	}

	u := User{
		Name: "jack",
		Age:11,
		Role:"user",
		Count: 1,
	}
	if err := material.MyDb.Create(&u).Error; err != nil {
		fmt.Println("新增出错=", err)
		return
	}
	fmt.Println("插入后id=", u.ID)//插入后id= 1

	//todo:要试用哪一个就取消注释哪一个
	f1(u)
	//f2(u)
	//f3(u)
	//f45(u)
	//f54(u)

	fmt.Println("完")
}

//有冲突时，什么都不做，即不插入也不更改。
func f1(u User) {
	fmt.Println("插入主键冲突的一条---------")
	u.ID = 1
	u.Count = 2
	res := material.MyDb.Clauses(clause.OnConflict{}).Create(&u)
	/*
	INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`role`,`count`,`id`)
	VALUES ('2022-11-21 21:40:53.565','2022-11-21 21:40:53.565',NULL,'jack',11,'user',2,1)
	ON DUPLICATE KEY UPDATE `id`=`id`
	*/
	if res.Error != nil {
		fmt.Println("err1=", res.Error)
		return
	}
	fmt.Println("res.Affected=", res.RowsAffected)//res.Affected= 0
	fmt.Println("res.u.id=", u.ID) //res.u.id= 1

	fmt.Println("插入唯一索引冲突的一条----------")
	u.ID = 2
	u.Count = 1
	res = material.MyDb.Clauses(clause.OnConflict{}).Create(&u)
	/*
	INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`role`,`count`,`id`)
	VALUES ('2022-11-21 21:40:53.565','2022-11-21 21:40:53.565',NULL,'jack',11,'user',1,2)
	ON DUPLICATE KEY UPDATE `id`=`id`
	*/
	if res.Error != nil {
		fmt.Println("err1=", res.Error)
		return
	}
	fmt.Println("另res.Affected=", res.RowsAffected) //另res.Affected= 0
	fmt.Println("另res.u.id=", u.ID) //另res.u.id= 2

	//本函数运行完后，数据库表中仍只有main函数中插入的那一条数据，且未发生任何改动。
}

// 有冲突时，将DoUpdates的map中指定的列更新为map指定的值。
func f2(u User) {
	fmt.Println("插入主键ID冲突的一条---------------")
	u.ID = 1
	u.Name = "mary"
	u.Role = "woman user"
	u.Count = 2
	res2 := material.MyDb.Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(map[string]interface{}{"role": "strange"}),
	}).Create(&u)
	/*
	INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`role`,`count`,`id`)
	VALUES ('2022-11-21 21:39:47.384','2022-11-21 21:39:47.384',NULL,'mary',11,'woman user',2,1)
	ON DUPLICATE KEY UPDATE `role`='strange'
	*/
	if res2.Error != nil {
		fmt.Println("err2=", res2.Error)
		return
	}
	fmt.Println("res2.Affected=", res2.RowsAffected)//res2.Affected= 2
	fmt.Println("res2.u.id=", u.ID)//res2.u.id= 1

	//本函数运行完后，数据库表中的第一行数据的role列的值由"user"变为了"strange"，而不是"woman user"，且其他列不变。
}

// 在冲突时，将指定要更新的列更新为新值
func f3(u User) {
	fmt.Println("插入unique索引Count字段冲突的一条---------------")
	u.Name = "ugly"
	u.Age = 999
	u.Count = 1
	res3 := material.MyDb.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"name", "age"}), //todo:将这里指定的name、age列更新为.Create(&u)中u相应字段的值。
	}).Create(&u)
	/*
	INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`role`,`count`,`id`)
	VALUES ('2022-11-21 21:38:19.53','2022-11-21 21:38:19.53',NULL,'ugly',999,'user',1,1)
	ON DUPLICATE KEY UPDATE `name`=VALUES(`name`),`age`=VALUES(`age`)
	*/
	if res3.Error != nil {
		fmt.Println("err3=", res3.Error)
		return
	}
	fmt.Println("res3.u.id=", u.ID)//res3.u.id= 1
	fmt.Println("res3.Affected=", res3.RowsAffected)//res3.Affected= 2

	//本函数运行完后，数据库表中count等于1的那一行数据的name列由jack变为ugly，age列由11变为999，其他不变。
}

// 在冲突时，更新所有列到新值（除了判定冲突的主键id的列）
func f45(u User) {
	fmt.Println("插入主键ID字段冲突的一条---------------")
	u.ID = 1
	u.Name = "newName"
	u.Age = 9999999
	u.Role = "newRole"
	u.Count = 111111111
	res4 := material.MyDb.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&u)
	/*
	INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`role`,`count`,`id`)
	VALUES ('2022-11-21 21:30:54.508','2022-11-21 21:30:54.508',NULL,'newName',9999999,'newRole',111111111,1)
	ON DUPLICATE KEY UPDATE `updated_at`='2022-11-21 21:30:54.512',`deleted_at`=VALUES(`deleted_at`),
	`name`=VALUES(`name`),`age`=VALUES(`age`),`role`=VALUES(`role`),`count`=VALUES(`count`)
	*/
	if res4.Error != nil {
		fmt.Println("err4=", res4.Error)
		return
	}
	fmt.Println("res4.Affected=", res4.RowsAffected)//res4.Affected= 2
	fmt.Println("res4.u.id=", u.ID)//res4.u.id= 1

	//运行完当前函数后，数据库表中ID为1的那一行除了id仍为1，其他列的值均变为新赋值的"newName"、9999999、"newRole"、包括Count的111111111
}

// 在冲突时，更新所有列到新值(除了主键id和判定冲突的unique索引的列)。
func f54(u User) {
	fmt.Println("插入unique索引Count字段冲突的一条---------------")
	u.ID = 11111111
	u.Name = "newName"
	u.Age = 9999999
	u.Role = "newRole"
	u.Count = 1
	res5 := material.MyDb.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&u)
	/*
	INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`role`,`count`,`id`)
	VALUES ('2022-11-21 21:35:01.343','2022-11-21 21:35:01.343',NULL,'newName',9999999,'newRole',1,11111111)
	ON DUPLICATE KEY UPDATE `updated_at`='2022-11-21 21:35:01.347',`delete	d_at`=VALUES(`deleted_at`),
	`name`=VALUES(`name`),`age`=VALUES(`age`),`role`=VALUES(`role`),`count`=VALUES(`count`)
	*/
	if res5.Error != nil {
		fmt.Println("err5=", res5.Error)
		return
	}
	fmt.Println("res5.Affected=", res5.RowsAffected) //res5.Affected= 2
	fmt.Println("res5.u.id=", u.ID) //res5.u.id= 11111111

	//运行完当前函数后，数据库表中ID为1的那一行除了id仍为1，count仍为1，其他列的值均变为新赋的"newName"、9999999、"newRole"
}
