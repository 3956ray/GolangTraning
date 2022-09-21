package main

import (
	"fmt"
)

/*
*
菜單
*/
func menu() bool {
	fmt.Println("--------------------------------------")
	fmt.Println(">> 動作代號：")
	fmt.Println("        1. 創建新產品物件並存入物件列表")
	fmt.Println("        2. 顯示當前產品物件列表")
	fmt.Println("        3. 操作指定產品物件")
	fmt.Println("        4. 使用device.GetType檢查產品型別")
	fmt.Println("        5. 結束")
	fmt.Printf("<< 請輸入執行的動作：")

	var option int
	fmt.Scanf("%d\n", &option)
	switch option {
	case 1:
		var optionDeviceType int
		fmt.Println(">> 開始創建產品物件")
		fmt.Printf("<< 請輸入產品類型(0:SAN|1:NAS|2:FAS)：")
		fmt.Scanf("%d\n", &optionDeviceType)
		if optionDeviceType == 0 {
			//輸入產品名稱
			var name string
			fmt.Printf("<< 請輸入產品名稱：")
			fmt.Scanf("%s\n", &name)
			//    創建新產品
			New("SAN", name)
		} else if optionDeviceType == 1 {
			//輸入產品名稱
			var name string
			fmt.Printf("<< 請輸入產品名稱：")
			fmt.Scanf("%s\n", &name)
			//    創建新產品
			New("NAS", name)
		} else if optionDeviceType == 2 {
			//輸入產品名稱
			var name string
			fmt.Printf("<< 請輸入產品名稱：")
			fmt.Scanf("%s\n", &name)
			//    創建新產品
			New("FAS", name)
		} else {
			//字體顔色為紅色"\x1b[%dm string \x1b[0m", 31  -->d輸入數字表示的顔色
			fmt.Printf("\x1b[%dmError: \x1b[0m請輸入正確的產品類型\n", 31)
		}
	case 2:
		//如果產品為空，輸出產品空
		if true {

		} else {
			fmt.Println("<< 尚未新增產品")
		}
	case 3:
		//如果產品為空，輸出為新增產品
		if true {
			var mapCap int
			fmt.Scanf("%d", &mapCap)

		} else {
			fmt.Println("<< 尚未新增產品")
		}
	case 4:
		var optionFour int
		fmt.Printf("<< 請輸入檢查項目(0:產品列表 | 1:指定產品 )：")
		fmt.Scanf("%d\n", &optionFour)
		if optionFour == 0 {
			//	返回產品列表
		} else if optionFour == 1 {
			//	指定產品，為空，返回
		}

	case 5:
		fmt.Println(">> Program exit")
		return false
	default:
		fmt.Println(">> 請輸入正確的動作")
	}
	return true
}

/**
菜單3實現
*/
