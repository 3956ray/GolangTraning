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
func (nas *NAS) Type() string {
	return nas.DeviceType
}

func (nas *NAS) Name() string {
	return nas.Names
}

/*
*
產生指定長度map
*/
func (nas *NAS) NewMap(len int) error {

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

func (nas *NAS) GetCurrentMap() map[string]int {
	return M
}

func (nas *NAS) StartCounter() error {

	count = 0
	timer = time.NewTimer(time.Second * 10)
	for {
		select {
		case <-timer.C:
			count = count + 1
			timer.Reset(time.Second * 10)
		}
	}

}

func (nas *NAS) StopCounter() error {
	if timer.Stop() {
		fmt.Println(">> 計數器關閉成功")
	} else {
		fmt.Println(">> 計數器未啟動")
	}
	return nil
}

func (nas *NAS) GetCurrentCount() int {
	return count
}

func (nas *NAS) Panic() {
	panic("throw panic")
}

func (nas *NAS) Destruct() {
}
