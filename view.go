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
		menuOption1()
	case 2:
		menuOption2()
	case 3:
		menuOption3()
	case 4:
		menuOption4()
	case 5:
		fmt.Println(">> Program exit")
		return false
	default:
		fmt.Println(">> 請輸入正確的動作")
	}
	return true
}

func menuOption1() {
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
}

func menuOption2() {
	//如果產品為空，輸出產品空
	if true {

	} else {
		fmt.Println("<< 尚未新增產品")
	}
}

/*
*
菜單3實現
*/
func menuOption3() {
	//如果產品為空，輸出為新增產品
	var deviceNumber int
	var option int
	//輸入產品編號
	fmt.Printf("<< 請輸入指定產品列表內的編號：")
	fmt.Scanf("%d\n", &deviceNumber)
	//檢測是否存在並選取
	if deviceNumber == 1 {
		for {
			fmt.Println("--------------------------------------")
			fmt.Println(">> 產品操作代號：")
			fmt.Println("        1. 顯示產品類型")
			fmt.Println("        2. 顯示產品名稱")
			fmt.Println("        3. 產生一個指定長度的Map")
			fmt.Println("        4. 啟動計數器")
			fmt.Println("        5. 停止計數器")
			fmt.Println("        6. 取得目前計數")
			fmt.Println("        7. 將產品資訊存成JSON檔案")
			fmt.Println("        8. 拋出panic")
			fmt.Println("        9. 結束")
			fmt.Printf("<< 請輸入執行的動作：")

			fmt.Scanf("%d\n", &option)
			switch option {
			case 1:
				//	回傳產品類型
				fmt.Println()
			case 2:
				//	回傳產品名稱
				fmt.Println()
			case 3:
				//	生成指定長度map
				var len int
				fmt.Print("請輸入Map的長度：")
				fmt.Scanf("%d\n", &len)
			//	NewMap(len)
			case 4:
			//	啟動一個計數器, SAN每0.5秒計數一次, NAS每10秒計數一次, FAS每3秒計數一次
			case 5:
			//	停止計數器
			case 6:
			//	取得目前計數器的值
			case 7:
			//	回傳當前的map
			case 8:
				//	拋出panic
				fmt.Println(">> 返回前列表")
				return
			case 9:
				//	結束
				return
			}
		}

	}

	if true {
		var mapCap int
		fmt.Scanf("%d", &mapCap)

	} else {
		fmt.Println("<< 尚未新增產品")
	}
}

func menuOption4() {
	var optionFour int
	fmt.Printf("<< 請輸入檢查項目(0:產品列表 | 1:指定產品 )：")
	fmt.Scanf("%d\n", &optionFour)
	if optionFour == 0 {
		//	返回產品列表
	} else if optionFour == 1 {
		//	指定產品，為空，返回
	}
}
