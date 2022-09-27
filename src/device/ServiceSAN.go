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

	count = 0
	timer = time.NewTimer(time.Millisecond * 500)
	for {
		select {
		case <-timer.C:
			count = count + 1
			timer.Reset(time.Millisecond * 500)
		}
	}
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
	return count
}

func (san *SAN) Panic() {
	panic("throw panic")
}

func (san *SAN) Destruct() {
}
