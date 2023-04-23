package cha

import "gorm.io/gorm"

type factory struct {
	Id      int
	Name    string
	Product string
	Output  int
	Type    string //农业、工业、服务业
	Belong  string //国企、私企
}

type User struct {
	gorm.Model
	Name string
	Age int
	Role string
	Count int
}

type Teacher struct {
	Id int
	Name string
	Subject string
}
type Student struct {
	Id int
	Name string
	TeacherId int
	Score int
	GenderId int
}
type Sex struct {
	Id int
	Gender string
}
