package main

import (
	"fmt"
)

func New(deviceType, name string) IDevice {
	switch deviceType {
	case "NAS":
		return &NAS{name: name}
	case "SAN":
		return &SAN{name: name}
	case "FAS":
		return &FAS{name: name}
	}

	return nil
}

func GetType(i interface{}) {
	fmt.Printf(">> 型別為: %T\n", i)
}

type IDevice interface {
	/**
	 *  回傳產品類型.
	 */
	Type() string

	/**
	 *  回傳產品名稱.
	 */
	Name() string

	/**
	 * 產生指定長度的string-integer map, key的格式為 <產品類型>_<產品名稱>_<編號>, value為隨機的integer (1~1000)
	 * Example: {"NAS_Will_1": 361, "NAS_Will_2": 56, "NAS_Will_3": 114}
	 */
	NewMap(len int) error

	/**

	 * 回傳當前的map
	 */
	GetCurrentMap() map[string]int

	/**
	 * 啟動一個計數器, SAN每0.5秒計數一次, NAS每10秒計數一次, FAS每3秒計數一次
	 */
	StartCounter() error

	/**
	 * 停止計數器
	 */
	StopCounter() error

	/**
	 * 取得目前計數器的值
	 */
	GetCurrentCount() int

	/**
	 * 拋出panic, panic訊息格式為"<產品類型> <產品名稱> throw panic"
	 */
	Panic()

	/**
	 * 印出一段Destruct的訊息, "<產品類型> <產品名稱> destruct success"
	 * Example: NAS Will destruct success
	 */
	Destruct()
}
