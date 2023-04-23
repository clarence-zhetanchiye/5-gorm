package main

//todo:---------------------------------------------使用 UpdateColumn、UpdateColumns------------------------------------
//.UpdateColumn()和.UpdateColumns() 可以更新的同时跳过update的钩子方法，而且不追踪更新时间，其用法类似于 Update、Updates
//update的钩子方法有：
//BeforeUpdate(*gorm.DB) error   BeforeSave(*gorm.DB) error
//AfterUpdate(*gorm.DB) error    AfterSave(*gorm.DB) error

/*
// 更新单个列
db.Model(&user).UpdateColumn("name", "hello")
// UPDATE users SET name='hello' WHERE id = 111;

// 更新多个列
db.Model(&user).UpdateColumns(User{Name: "hello", Age: 18})
// UPDATE users SET name='hello', age=18 WHERE id = 111;
*/