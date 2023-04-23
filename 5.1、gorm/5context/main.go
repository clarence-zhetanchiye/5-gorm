package main

import (
	"context"
	material "context/0material"
	"fmt"
	"time"
)

//todo:-----------------------------------------------context在gorm语句中的使用-------------------------------------------
//												参见 7链式操作-与新建会话  8新建会话Sessoin
func main() {
	material.GetDB()
	material.InsertSoft()

	//todo:获取Context的一种方式。-----------------------------------------------------
	_ = material.MyDb.Statement.Context //通过这个获得的context，是boardmix导图中步骤2.5初始化中赋值的context.Background()

	//todo:获取Context也可以自己手动创建------------------------------------------------
	// 特别是创建超时的context、手动决定何时结束的context等等。
	ctx := context.Background()
	ctxT, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()


	//todo:在gorm中使用Context。------------------------------------------------------
	// 点进去看源码可知是建其一个gorm的Seesion临时会话，并设置进该ctx。
	material.MyDb.WithContext(ctx)

	//todo:单会话模式————被用于执行单次gorm语句
	var s []material.Soft
	material.MyDb.WithContext(ctxT).Find(&s)
	fmt.Println("查询结果s=", s)

	//todo:持续会话模式————用于执行一系列gorm语句
	dbCtx := material.MyDb.WithContext(ctxT)
	var ones material.Soft
	dbCtx.Table("softs").First(&ones)
	dbCtx.Table("softs").Where("id=?", 1).Update("name=?", "x")

}
