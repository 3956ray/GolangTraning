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
	return nas.Dtype
}

func (nas *NAS) Name() string {
	return nas.Name()
}

/*
*
產生指定長度map
*/
func (nas *NAS) NewMap(len int) error {
	rand.Seed(time.Now().UnixNano())
	//生成指定長度map
	m := make(map[string]int, len)
	for i := 0; i < len; i++ {
		//1-1000 == (0-999)+1
		m["NAS_Ray_"+string(i)] = rand.Intn(999) + 1
	}

	for k, v := range m {
		fmt.Printf("%v: %d", k, v)
	}

	//斷言？
	return nil
}

func (nas *NAS) GetCurrentMap() map[string]int {
	return nil
}

func (nas *NAS) StartCounter() error {
	TimeCount = time.NewTimer(time.Second * 10)
	return nil
}

func (nas *NAS) StopCounter() error {
	TimeCount.Stop()
	return nil
}

func (nas *NAS) GetCurrentCount() int {
	return 0
}

func (nas *NAS) Panic() {

}

func (nas *NAS) Destruct() {
}
