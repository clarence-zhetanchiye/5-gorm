package main

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	material "scopes/1material"
	"strconv"
)
//todo:分页器------------------------------------------------------------------------------------------------------------
func main() {
	material.GetDB()
	r := &http.Request{Header: http.Header{}}
	r.Header.Set("page", "2")
	r.Header.Set("page_size", "2")
	var gs []material.Good
	material.MyDb.Scopes(Paginate(r)).Find(&gs) //上面Paginate中的.Offset(..).Limit(..)就会写进这个gorm语句里。
	//SELECT * FROM `goods` LIMIT 2 OFFSET 2
	fmt.Println("gs=", gs)//gs= [{3 banana 3 300 333} {4 banana 6 400 444}]
}
func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(r.Header.Get("Page"))
		fmt.Println("页码=", page)//页码= 2
		if page == 0 {
			page = 1
		}
		pageSize, _ := strconv.Atoi(r.Header.Get("Page_size"))
		fmt.Println("页容量=", pageSize)//页容量= 2
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
