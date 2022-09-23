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
func (nas *NAS) Type() string {
	return nas.deviceType
}

func (nas *NAS) Name() string {
	return nas.name
}

/*
*
產生指定長度map
*/
func (nas *NAS) NewMap(len int) error {
	rand.Seed(time.Now().UnixNano())
	//生成指定長度map
	m = make(map[string]int, len)
	for i := 0; i < len; i++ {
		//1-1000 == (0-999)+1
		m["NAS_Ray_"+string(i)] = rand.Intn(999) + 1
	}

	for k, v := range m {
		fmt.Printf("%v: %d", k, v)
	}
	println()
	//斷言？
	return nil
}

func (nas *NAS) GetCurrentMap() map[string]int {
	return m
}

func (nas *NAS) StartCounter() error {
	count = 0
	//每十秒計數一次
	timer = time.NewTimer(time.Second * 10)
	select {
	case <-timer.C:
		count++
	}
	fmt.Println(">> 計數器啟動成功")
	return nil
}

func (nas *NAS) StopCounter() error {
	if timer.Stop() {
		fmt.Println(">> 計數器關閉成功")
		timer.Reset(time.Second)
	} else {
		fmt.Println(">> 計數器未啟動")
	}
	return nil
}

func (nas *NAS) GetCurrentCount() int {
	fmt.Println(count)
	return count
}

func (nas *NAS) Panic() {

}

func (nas *NAS) Destruct() {
}
