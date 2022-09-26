package main

import (
	"GolangTraning/src/view"
)

func main() {
	//初始化一個devicelist
	view.NewDeviceList()
	for {
		//	顯示菜單
		if !view.Menu() {
			break
		}
	}

}
