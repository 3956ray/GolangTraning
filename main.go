package main

import (
	"time"
)

var timer *time.Timer
var count int
var m map[string]int

//type device struct {
//}
//type BookCompany interface {
//	Create() Book
//}
//
//type Book interface {
//	ReadBook() string
//	WriteBook() string
//}
//
//// 實現book的功能
//func (c *EnglishBook) ReadBook() string {
//	return "reading a english book"
//}
//func (c *ChineseBook) ReadBook() string {
//	return "讀中文書"
//}
//func (c *EnglishBook) WriteBook() string {
//	return "i'am writing"
//}
//func (c *ChineseBook) WriteBook() string {
//	return "我在創作中"
//}
//
//type EnglishBook struct{}
//type ChineseBook struct{}
//type BeiJinCompany struct{}
//type ShanghaiCompany struct{}
//
//// 一個印刷場刷英語書，一個刷中文書
//func (c *BeiJinCompany) Create() Book {
//	return &EnglishBook{}
//}
//func (c *ShanghaiCompany) Create() Book {
//	return &ChineseBook{}
//}

func main() {
	deviceList = make([]IDevice, 0, 100)
	for {
		//	顯示菜單
		if !menu() {
			break
		}
	}
	//	fmt.Printf("%s", deviceList[0].Name())

	//var f BookCompany
	//f = new(BeiJinCompany)
	//book := f.Create()
	//fmt.Println(book.ReadBook())
}
