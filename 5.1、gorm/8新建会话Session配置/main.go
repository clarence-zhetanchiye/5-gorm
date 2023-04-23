package main

import (
	"fmt"
	"gorm.io/gorm"
	material "session/1material"
)

func main() {
	material.GetDB()
	material.InsertGoods()

	//dryRun()
	//prepareStmt()
	//newDB()
	//initialized()
	//queryFields()

	//更多的可以另行再参考 0a说明/0新建会话配置.txt 中的整理写出demo予以验证。
}

func dryRun() {
	NewSession := material.MyDb.Session(&gorm.Session{DryRun: true})
	var g material.Good
	dbSt := NewSession.Table("goods").Where("id=?", 1).Find(&g)
	//SELECT * FROM `goods` WHERE id=1
	fmt.Println("g=", g)//g= {0  0 0 0}   //todo:由此可见gorm语句对应的sql没有真正被执行。
	fmt.Println("sql=", dbSt.Statement.SQL.String())//sql= SELECT * FROM `goods` WHERE id=?
	fmt.Println("vars=", dbSt.Statement.Vars)//vars= [1]
	fmt.Println("explain=", NewSession.Dialector.Explain(dbSt.Statement.SQL.String(), dbSt.Statement.Vars))
	//explain= SELECT * FROM `goods` WHERE id='[1]'
}

func prepareStmt() {
	//下面一行可点进去看源码。另外，下面一行让NewSession.ConnPool字段就被赋为&gorm.PreparedStmtDB{}
	NewSession := material.MyDb.Session(&gorm.Session{PrepareStmt: true})

	NewSession.Table("goods").Where("id=?", 1).Update("price", 10000)
	var g material.Good
	NewSession.Table("goods").Where("id=?", 1).Find(&g)
	fmt.Println("g=", g)//g= {1 apple 10000 100 111}

	prepareManager, ok := NewSession.ConnPool.(*gorm.PreparedStmtDB)
	if !ok {
		fmt.Println("出错")
		return
	}
	fmt.Println("MySQL已预编译的sql=", prepareManager.PreparedSQL)
	//MySQL已预编译的sql= [UPDATE `goods` SET `price`=? WHERE id=?   SELECT * FROM `goods` WHERE id=?]
	fmt.Println("MySQL已预编译的sql=", prepareManager.Stmts)
	//MySQL已预编译的sql= map[SELECT * FROM `goods` WHERE id=?:{0xc00028a990 false}
	//						UPDATE `goods` SET `price`=? WHERE id=?:{0xc00028a870 true}]
	for PreparedSQL, Stmt := range prepareManager.Stmts {
		fmt.Println("k=", PreparedSQL)
		//k= UPDATE `goods` SET `price`=? WHERE id=?
		//k= SELECT * FROM `goods` WHERE id=?
		Stmt.Close()
	}
	//prepareManager.Close()//这一行里源码的本质就上面的遍历的Close()，因此不必重复Close()
}

//todo:给NewDB赋值true，会让新建会话前的Where()条件不再影响之后的gorm语句
func newDB() {
	session := material.MyDb.Where("id<?", 6).Session(&gorm.Session{})
	var r material.Good
	session.Table("goods").Where("name=?", "orange").Find(&r)
	//SELECT * FROM `goods` WHERE id<6 AND name='orange'

	sess := material.MyDb.Where("id<?", 6).Session(&gorm.Session{NewDB: true})
	var e material.Good
	sess.Table("goods").Where("name=?", "orange").Find(&e)
	//SELECT * FROM `goods` WHERE name='orange'
}

//todo:设置Initialized为true会让新建会话不再是新建，即这样新建的会话是个假的新建会话，仍会让前后两个gorm语句的链式方法如Where()叠加地互相影响。
func initialized() {
	se := material.MyDb.Where("id<?", 6).Session(&gorm.Session{Initialized: false})
	var a material.Good
	se.Table("goods").Where("name=?", "orange").Find(&a)
	//SELECT * FROM `goods` WHERE id<6 AND name='orange'

	var b material.Good
	se.Table("goods").Where("name=?", "apple").Find(&b)
	//SELECT * FROM `goods` WHERE id<6 AND name='apple'  //todo:注意看这里没有name='orange'


	sess := material.MyDb.Session(&gorm.Session{Initialized: true})
	var n material.Good
	sess.Table("goods").Where("name=?", "orange").Find(&n)
	//SELECT * FROM `goods` WHERE name='orange'

	var m material.Good
	sess.Table("goods").Where("name=?", "apple").Find(&m)
	//SELECT * FROM `goods` WHERE name='orange' AND name='apple'  //todo:注意看这里的 AND name='orange'


	session := material.MyDb.Where("id<?", 6).Session(&gorm.Session{Initialized: true})
	var r material.Good
	session.Table("goods").Where("name=?", "orange").Find(&r)
	//SELECT * FROM `goods` WHERE id<6 AND name='orange'

	var p material.Good
	session.Table("goods").Where("name=?", "apple").Find(&p)
	//SELECT * FROM `goods` WHERE id<6 AND name='orange' AND name='apple'  //todo:注意看这里的 AND name='orange'
}

//todo:让最后生成的sql语句中的显示地指明要查询的字段。详见下面的举例。
func queryFields() {
	var d material.Good
	material.MyDb.Session(&gorm.Session{}).Table("goods").Where("id=?", 1).Find(&d)
	//SELECT * FROM `goods` WHERE id=1

	var g material.Good
	material.MyDb.Session(&gorm.Session{QueryFields: true}).Table("goods").Where("id=?", 1).Find(&g)
	//SELECT `goods`.`id`,`goods`.`name`,`goods`.`price`,`goods`.`amount`,`goods`.`code` FROM `goods` WHERE id=1

	var f material.Good
	material.MyDb.Table("goods").Select("id, name").Where("id=?", 1).Find(&f)
	//SELECT id, name FROM `goods` WHERE id=1

	var h material.Good
	material.MyDb.Session(&gorm.Session{}).Table("goods").Select("id, name").Where("id=?", 1).Find(&h)
	//SELECT id, name FROM `goods` WHERE id=1

	//todo:下面这样使用，实现了智能查询部分字段。参见3crud/2cha2/cha/44select.go
	type halfGood struct {
		Id int
		Name string
		//good的其他字段不予承接。
	}
	var x halfGood
	material.MyDb.Session(&gorm.Session{QueryFields: true}).Table("goods").Where("id=?", 1).Find(&x)
	//SELECT `goods`.`id`,`goods`.`name` FROM `goods` WHERE id=1  //可以看见指明字段了
	fmt.Println("x=", x)//x= {1 apple}
}
