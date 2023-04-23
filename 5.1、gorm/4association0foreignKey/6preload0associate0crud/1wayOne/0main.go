//--操作父结构体若要在crud时跳过或只针对关联结构体字段，可以使用 Select 或 Omit--------------------------------
//------这种方式相对于Association下的Find,Append,Replace,Delete,Clear,Count的
//增查删改是一种相对原始的方式
package main
/*
import "gorm.io/gorm/clause"

func main() {
	user := User{
		Name:            "jinzhu",
		BillingAddress:  Address{Address1: "Billing Address - Address 1"},
		ShippingAddress: Address{Address1: "Shipping Address - Address 1"},
		Emails:          []Email{
			{Email: "jinzhu@example.com"},
			{Email: "jinzhu-2@example.com"},
		},
		Languages:       []Language{
			{Name: "ZH"},
			{Name: "EN"},
		},
	}

	db.Select("Name").Create(&user)
	//只想字段Name对应的数据库表新增一条只写入Name字段值的数据。
   //即  INSERT INTO "users" (name) VALUES ("jinzhu");
//但其他的字段会同时插入null????

	db.Omit("BillingAddress").Create(&user)
	// 创建 user 时，跳过向BillingAddress表新增数据。
但同时BillingAddress表会同时插入一行null？？

	db.Omit(clause.Associations).Create(&user)
	// 创建user时，跳过向各个关联表新增数据，只向users表新增INSERT INTO "users"(name)VALUES("jinzhu");
//但会同时在其他各个关联表中插入一行null？？？

	//---------todo:crud可参照此进行--------
}
 */
