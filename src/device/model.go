package device

import "time"

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

type JsonData struct {
	Type         string
	Name         string
	CurrentMap   map[string]int
	CurrentCount int
}

// 機器列表
var DeviceList []IDevice

var M map[string]int

// 計數
var count int

// 定時器
var timer *time.Timer

// 定時器啓動flag
var FlagTimer = false
