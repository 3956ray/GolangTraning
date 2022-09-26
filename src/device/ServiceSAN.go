package device

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
	return san.DeviceType
}

func (san *SAN) Name() string {
	return san.Names
}

/*
*
產生指定長度map
*/
func (san *SAN) NewMap(len int) error {
	rand.Seed(time.Now().UnixNano())
	//生成指定長度map
	M = make(map[string]int, len)
	for i := 0; i < len; i++ {
		//1-1000 == (0-999)+1
		M[fmt.Sprintf("NAS_Ray_%d", i)] = rand.Intn(999) + 1
	}

	fmt.Println(M)
	return nil
}

func (san *SAN) GetCurrentMap() map[string]int {
	return M
}

func (san *SAN) StartCounter() error {
	//view.count = 0
	////每0.5秒計數一次
	//view.timer = time.NewTimer((time.Second * 1) / 2)
	//select {
	//case <-view.timer.C:
	//	view.count++
	//}
	return nil
}

func (san *SAN) StopCounter() error {
	//if view.timer.Stop() {
	//	fmt.Println(">> 計數器關閉成功")
	//} else {
	//	fmt.Println(">> 計數器未啟動")
	//}
	return nil
}

func (san *SAN) GetCurrentCount() int {
	//fmt.Println(view.count)
	//return view.count
	return 0
}

func (san *SAN) Panic() {
	panic("throw panic")
}

func (san *SAN) Destruct() {
}
