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
func (fas *FAS) Type() string {
	return fas.deviceType
}

func (fas *FAS) Name() string {
	return fas.name
}

/*
*
產生指定長度map
*/
func (fas *FAS) NewMap(len int) error {
	rand.Seed(time.Now().UnixNano())
	//生成指定長度map
	m := make(map[string]int, len)
	for i := 0; i < len; i++ {
		//1-1000 == (0-999)+1
		m["FAS_Ray_"+string(i)] = rand.Intn(999) + 1
	}

	for k, v := range m {
		fmt.Printf("%v: %d", k, v)
	}

	//斷言？
	return nil
}

func (fas *FAS) GetCurrentMap() map[string]int {
	return m
}

func (fas *FAS) StartCounter() error {
	count = 0
	//每三秒計數一次
	timer = time.NewTimer(time.Second * 3)

	select {
	case <-timer.C:
		count++
	}
	return nil
}

func (fas *FAS) StopCounter() error {
	if timer.Stop() {
		fmt.Println(">> 計數器關閉成功")
	} else {
		fmt.Println(">> 計數器未啟動")
	}
	return nil
}

func (fas *FAS) GetCurrentCount() int {
	return count
}

func (fas *FAS) Panic() {

}

func (fas *FAS) Destruct() {
}
