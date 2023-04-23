package main

import material "crud/0material"

func save() {
	//todo:新增——-----------------------------------------若入参结构体的主键为零值，则Save()就是新增该入参结构体为数据表的一条记录。
	row1 := good{
		Id:     0, //todo:主键是零值或数据表中尚不存在
		Name:   "new",
		Price:  1,
		Amount: 1,
		Code: 99999999, //唯一键，和数据表中已有数据不冲突
	}
	material.MyDb.Save(&row1) //INSERT INTO `goods` (`name`,`price`,`amount`,`code`) VALUES ('new',1,1,99999999)

	//todo:该Save()的效果是新增还是修改（入参是结构体时），只看主键Id，而不管unique键，即和3crud/1zeng1/3Upsert.go中的upsert不一样。
	row11 := good{
		Id:     0, //主键是零值或数据表中尚不存在
		Name:   "new",
		Price:  1,
		Amount: 1,
		Code: 111, //todo:和数据表中已有数据冲突，因此并不会成功插入而会报错，且会浪费一次主键的自增。
	}
	material.MyDb.Save(&row11)
	//INSERT INTO `goods` (`name`,`price`,`amount`,`code`) VALUES ('new',1,1,111)
	//todo:报错：Duplicate entry '111' for key 'code'


	//todo:更改——--------------------若入参结构体的主键有值且数据表中已存在该主键，则Save()会保存入参结构体的所有的字段，即使字段是零值
	row2 := good{
		Id: 1, //todo:主键，该主键的值在数据表中已经存在
		Name: "newName",
		Price: 0,
		Amount: 0,
		Code: 1000000, //唯一键，和数据表中已有数据不冲突
	}
	material.MyDb.Save(&row2) //UPDATE `goods` SET `name`='newName',`price`=0,`amount`=0,`code`=1000000 WHERE `id` = 1

	//todo:当Save()的入参为数组或切片时---------------------------------------Save()等价于3crud/1zeng1/3Upsert.go中的upsert。
	row3 := []good{
		{
			Id: 0, //主键是零值或数据表中尚不存在
			Name: "x",
			Price: 1,
			Amount: 1,
			Code: 121, //唯一键，和数据表中已有数据不冲突
		},
		{
			Id: 0, //主键是零值或数据表中尚不存在
			Name: "xx",
			Price: 2,
			Amount: 2,
			Code: 232, //唯一键，和数据表中已有数据不冲突
		},
		{
			Id: 2, //主键，该值在数据表中已存在。因此该结构体的效果是更改数据表中相应的一行内容。
			Name: "new-second",
			Price: 0,
			Amount: 0,
			Code: 6969, //唯一键，和数据表中已有数据不冲突
		},
		{
			Id: 0,
			Name: "unique",
			Price: 11,
			Amount: 11,
			Code: 666, //todo:数据表中id为6那行的code也是666，即unique键冲突。因此实际效果是更新了id为6冲突的这一行，即是Upsert的效果
		},
	}
	//todo:由结果可知入参的切片中，前面两个结构体是新增，第三个是修改数据表中id为2的一行。第四个是修改数据表中code为666的那行数据。
	material.MyDb.Save(&row3)
	//INSERT INTO `goods` (`name`,`price`,`amount`,`code`,`id`)
	//VALUES ('x',1,1,121,DEFAULT),('xx',2,2,232,DEFAULT),('new-second',0,0,6969,2),('unique',11,11,666,DEFAULT)
	//ON DUPLICATE KEY UPDATE `name`=VALUES(`name`),`price`=VALUES(`price`),`amount`=VALUES(`amount`),`code`=VALUES(`code`)

}
