package main
import "gorm.io/gorm"
//todo:------------------------------------------------暂时没懂透，还需细究------------------------------------------------


//many to many 关系，如下为例，是指一个person对应任意多个language,一个language也可以对应
//任意多个person

type person struct{
	gorm.Model
	Name string
	Langs []language  `gorm:"many2many:user_languages"`
					//Automigrate创建person时会自动的创建名为user_languages的中间表
}
type language struct{
	gorm.Model
	Character string
}

func main() {
	
}
