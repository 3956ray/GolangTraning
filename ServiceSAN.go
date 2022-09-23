package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
*
输出产品名称和类型
*/
func (san *SAN) Type() string {
	return san.deviceType
}

func (san *SAN) Name() string {
	return san.name
}

/*
*
產生指定長度map
*/
func (san *SAN) NewMap(len int) error {
	rand.Seed(time.Now().UnixNano())
	//生成指定長度map
	m = make(map[string]int, len)
	for i := 0; i < len; i++ {
		//1-1000 == (0-999)+1
		m[fmt.Sprintf("SAN_Ray_%d ", i)] = rand.Intn(999) + 1
	}

	fmt.Println(m)

	return nil
}

func (san *SAN) GetCurrentMap() map[string]int {
	return m
}

func (san *SAN) StartCounter() error {
	count = 0
	//每0.5秒計數一次
	timer = time.NewTimer((time.Second * 1) / 2)
	select {
	case <-timer.C:
		count++
	}
	return nil
}

func (san *SAN) StopCounter() error {
	if timer.Stop() {
		fmt.Println(">> 計數器關閉成功")
	} else {
		fmt.Println(">> 計數器未啟動")
	}
	return nil
}

func (san *SAN) GetCurrentCount() int {
	fmt.Println(count)
	return count
}

func (san *SAN) Panic() {

}

func (san *SAN) Destruct() {
}
