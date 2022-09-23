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
		for {
			if !menuOption3() {
				break
			}
		}
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
		NewDeviceTypeNAS("NAS")
		deviceList = append(deviceList, NewDevice("SAN", name))

	} else if optionDeviceType == 1 {
		//輸入產品名稱
		var name string
		fmt.Printf("<< 請輸入產品名稱：")
		fmt.Scanf("%s\n", &name)
		//    創建新產品
		NewDeviceTypeSAN("NAS")
		deviceList = append(deviceList, NewDevice("NAS", name))

	} else if optionDeviceType == 2 {
		//輸入產品名稱
		var name string
		fmt.Printf("<< 請輸入產品名稱：")
		fmt.Scanf("%s\n", &name)
		//    創建新產品
		deviceList = append(deviceList, NewDevice("FAS", name))
		NewDeviceTypeFAS("FAS")
	} else {
		//字體顔色為紅色"\x1b[%dm string \x1b[0m", 31  -->d輸入數字表示的顔色
		fmt.Printf("\x1b[%dmError: \x1b[0m請輸入正確的產品類型\n", 31)
	}
}

func menuOption2() {
	//如果產品為空，輸出產品空
	if len(deviceList) > 1 {
		for i := 1; i < len(deviceList); i++ {
			fmt.Printf("        %d. 產品類型：%s 產品名稱：%s\n", i, deviceList[i].Type(), deviceList[i].Name())
		}
	} else {
		fmt.Println("<< 尚未新增產品")
	}
}

/*
*
菜單3實現
*/
func menuOption3() bool {
	//如果產品為空，輸出為新增產品
	var deviceNumber int
	var option int
	//檢測是否存在並選取
	if len(deviceList) >= 1 {
		//顯示全部產品
		menuOption2()
		//輸入產品編號
		fmt.Printf("<< 請輸入指定產品列表內的編號：")
		fmt.Scanf("%d\n", &deviceNumber)

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
				fmt.Println(deviceList[deviceNumber].Type())
			case 2:
				//	回傳產品名稱
				fmt.Println(deviceList[deviceNumber].Name())
			case 3:
				//	生成指定長度map
				var len int
				fmt.Print("請輸入Map的長度：")
				fmt.Scanf("%d\n", &len)
				deviceList[deviceNumber].NewMap(len)
			case 4:
				//	啟動一個計數器, SAN每0.5秒計數一次, NAS每10秒計數一次, FAS每3秒計數一次
				deviceList[deviceNumber].StartCounter()
			case 5:
				//	停止計數器
				deviceList[deviceNumber].StopCounter()
			case 6:
				//	取得目前計數器的值
				deviceList[deviceNumber].GetCurrentCount()
			case 7:
			//	回傳當前的map
			case 8:
				//	拋出panic,並recover
				fmt.Println(">> 返回前列表")
				return false
			case 9:
				//	結束
				return false
			}
		}

	} else {
		fmt.Println("<< 尚未新增產品")
		return false
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

func NewDevice(deviceType, name string) IDevice {
	switch deviceType {
	case "NAS":
		return &NAS{deviceType: deviceType, name: name}
	case "SAN":
		return &SAN{deviceType: deviceType, name: name}
	case "FAS":
		return &FAS{deviceType: deviceType, name: name}
	}

	return nil
}
