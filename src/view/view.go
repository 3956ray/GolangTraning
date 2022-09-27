package view

import (
	"GolangTraning/src/device"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
*
菜單
*/
func Menu() bool {
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
			//使用bool進行判斷是否結束此循環
			if !menuOption3() {
				break
			}
		}
	case 4:
		menuOption4()
	case 5:
		menuOption5()
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
		if name == "" {
			fmt.Printf("\x1b[%dmError: \x1b[0m請輸入正確的產品名稱\n", 31)
			return
		}
		//    創建新產品
		device.DeviceList = append(device.DeviceList, NewDevice("SAN", name))

	} else if optionDeviceType == 1 {
		//輸入產品名稱
		var name string
		fmt.Printf("<< 請輸入產品名稱：")
		fmt.Scanf("%s\n", &name)
		if name == "" {
			fmt.Printf("\x1b[%dmError: \x1b[0m請輸入正確的產品名稱\n", 31)
			return
		}
		//    創建新產品
		device.DeviceList = append(device.DeviceList, NewDevice("NAS", name))

	} else if optionDeviceType == 2 {
		//輸入產品名稱
		var name string

		fmt.Printf("<< 請輸入產品名稱：")
		fmt.Scanf("%s\n", &name)
		if name == "" {
			fmt.Printf("\x1b[%dmError: \x1b[0m請輸入正確的產品名稱\n", 31)
			return
		}
		//    創建新產品
		device.DeviceList = append(device.DeviceList, NewDevice("FAS", name))
	} else {
		//字體顔色為紅色"\x1b[%dm string \x1b[0m", 31  -->d輸入數字表示的顔色
		fmt.Printf("\x1b[%dmError: \x1b[0m請輸入正確的產品類型\n", 31)
	}
}

func menuOption2() {
	//如果產品為空，輸出產品空
	if len(device.DeviceList) > 1 {
		for i := 1; i < len(device.DeviceList); i++ {
			fmt.Printf("        %d. 產品類型：%s 產品名稱：%s\n", i, device.DeviceList[i].Type(), device.DeviceList[i].Name())
		}
	} else {
		fmt.Println("<< 尚未新增產品")
	}
}

func menuOption3() bool {
	//如果產品為空，輸出為新增產品
	var deviceNumber int

	//檢測是否存在並選取
	if len(device.DeviceList) > 1 {
		//顯示全部產品
		menuOption2()

		//輸入產品編號
		fmt.Printf("<< 請輸入指定產品列表內的編號：")
		fmt.Scanf("%d\n", &deviceNumber)

		if deviceNumber < len(device.DeviceList) && deviceNumber > 0 {
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

				var option int
				fmt.Scanf("%d \n", &option)

				//檢車scanf輸入情況
				//n, err := fmt.Scanf("%d \n", &option)
				//fmt.Println("n:", n, "err", err)

				switch option {
				case 1:
					//	回傳產品類型
					fmt.Println(">> 當前產品類型: " + device.DeviceList[deviceNumber].Type())
				case 2:
					//	回傳產品名稱
					fmt.Println(">> 當前產品名稱:" + device.DeviceList[deviceNumber].Name())
				case 3:
					//	生成指定長度map
					var len int
					fmt.Print("請輸入Map的長度：")
					fmt.Scanf("%d\n", &len)
					device.DeviceList[deviceNumber].NewMap(len)
				case 4:
					//  檢查計數器是否為0
					//	啟動一個計數器, SAN每0.5秒計數一次, NAS每10秒計數一次, FAS每3秒計數一次
					if !device.FlagTimer {
						go device.DeviceList[deviceNumber].StartCounter()
						device.FlagTimer = true
						fmt.Println(">> 計數器啟動成功")
					} else {
						fmt.Println(">> 計數器正在執行中")
					}
				case 5:
					if !device.FlagTimer {
						fmt.Println(">> 計數器未啟動")
					} else {
						device.DeviceList[deviceNumber].StopCounter()
					}

				case 6:
					//	取得目前計數器的值
					fmt.Printf(">> 目前計數為: %d\n", device.DeviceList[deviceNumber].GetCurrentCount())
				case 7:
					mapdata := map[string]interface{}{
						"Type":       device.DeviceList[deviceNumber].Type(),
						"Name":       device.DeviceList[deviceNumber].Name(),
						"CurrentMap": device.M,
					}
					//fmt.Println(mapdata)

					//var mapdata sync.Map
					//
					//mapdata.Store("Type", device.DeviceList[deviceNumber].Type())
					//mapdata.Store("Name", device.DeviceList[deviceNumber].Name())
					//mapdata.Store("CurrentMap", device.M)
					//格式化json輸出
					mapData, err := json.MarshalIndent(mapdata, "", "\t")

					if err != nil {
						fmt.Println("failed", err)
					}
					fmt.Println("map: ", string(mapData))

					//寫入json文檔
					WriteBytes(mapData)
				case 8:
					//	拋出panic,並recover
					defer func() {
						if e := recover(); e != nil {
							fmt.Printf("Panic recover %s %s %s\n", device.DeviceList[deviceNumber].Type(), device.DeviceList[deviceNumber].Name(), e)
							fmt.Println(">> 返回前列表")
						}
					}()
					device.DeviceList[deviceNumber].Panic()
				case 9:
					//將timer結束，flagTimer設置爲false表示未使用
					if device.FlagTimer {
						device.DeviceList[deviceNumber].StopCounter()
					}
					device.FlagTimer = false

					//	結束
					return false
				}
			}
		} else {
			fmt.Println("<< 產品不存在")
			return false
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

	switch optionFour {
	case 0:
		device.GetType(device.DeviceList)
	case 1:

		if len(device.DeviceList) > 1 {
			var deviceNumber int
			//顯示全部產品
			menuOption2()

			//輸入產品編號
			fmt.Printf("<< 請輸入指定產品列表內的編號：")
			fmt.Scanf("%d\n", &deviceNumber)

			//不在此區間即表示產品不存在
			if deviceNumber < len(device.DeviceList) && deviceNumber > 0 {
				device.GetType(device.DeviceList[deviceNumber])
			} else {
				fmt.Println("<< 產品不存在")
			}

		} else {
			//	為空，返回
			fmt.Println("<< 尚未新增產品")
		}
	default:
		device.GetType(device.DeviceList)
	}
	if optionFour == 0 {
		//	返回產品列表

	}
}

func menuOption5() {
	for i := 1; i < len(device.DeviceList); i++ {
		fmt.Printf(">> %s %s destruct success\n", device.DeviceList[i].Type(), device.DeviceList[i].Name())
	}
}

/*
*
創建設備
*/
func NewDevice(deviceType, name string) device.IDevice {
	switch deviceType {
	case "NAS":
		return &device.NAS{DeviceType: deviceType, Names: name}
	case "SAN":
		return &device.SAN{DeviceType: deviceType, Names: name}
	case "FAS":
		return &device.FAS{DeviceType: deviceType, Names: name}
	}

	return nil
}

/*
*
json數據寫入。並生成json文件
*/
func WriteBytes(b []byte) {

	jsonfilepath := "./"
	//默認檔名：
	filename := "info"
	filetype := ".json"
	fmt.Printf(">> 請輸入檔案名稱(default: info.json)：")
	fmt.Scanf("%s\n", &filename)

	fmt.Printf("Filename: \x1b[%dm%s\x1b[0m\n", 34, filename)
	fmt.Println(">> " + filename + "已開啟")

	//	寫入到路徑下
	err := ioutil.WriteFile(jsonfilepath+filename+filetype, b, 0644)
	if err != nil {
		fmt.Println("Writefile Error =", err)
		return
	}
	fmt.Println(">> " + filename + "已寫入")
	fmt.Println(">> " + filename + "已關閉")

	//讀取檔案并輸出
	//fileContent, err := ioutil.ReadFile(jsonfilepath + filename)
	//if err != nil {
	//	fmt.Println("Read file err =", err)
	//	return
	//}
	//
	fmt.Printf("Filename: \x1b[%dm%s\x1b[0m \x1b[%dmsuccess\x1b[0m\n", 34, filename, 32)
	//fmt.Printf("Read \x1b[%dmsuccess\x1b[0m = %s\n", 32, string(fileContent))

}

/*
*
初始化devicelist
*/
func NewDeviceList() {
	device.DeviceList = make([]device.IDevice, 1, 100)
}
