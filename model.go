package main

type NAS struct {
	deviceType, name string
}

type FAS struct {
	deviceType, name string
}

type SAN struct {
	deviceType, name string
}

// 機器列表
var deviceList []IDevice

//// 機器製造廠，抽象工廠
//type DeviceFactory interface {
//	New(deviceType, name string) IDevice
//}
//
//// 工廠實例化：工廠創建
//type deviceFactory struct {
//	deviceType, name string
//}

// 機器製造廠，抽象工廠
//type Device interface {
//	New(deviceType, name string)
//}

//// 工廠實例化：工廠創建
//type DeviceFactory struct {
//	deviceType, name string
//}
//
//func (deviceFactory DeviceFactory) New(deviceType, name string) Device {
//	switch deviceType {
//	case "NAS":
//		return &NAS{name: name}
//	case "SAN":
//		return &SAN{name: name}
//	case "FAS":
//		return &FAS{name: name}
//	}
//}

//切片數組,存放單位為動態的struct類型

//type Device struct {
//	Dtype, name string
//}
//
//var NAS Device

//創建列表
