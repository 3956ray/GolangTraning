package device

type NAS struct {
	DeviceType, name string
	Names            string
}

type FAS struct {
	DeviceType, name string
	Names            string
}

type SAN struct {
	DeviceType, name string
	Names            string
}

// 機器列表
var DeviceList []IDevice

var M map[string]int
