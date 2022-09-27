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
func (fas *FAS) Type() string {
	return fas.DeviceType
}

func (fas *FAS) Name() string {
	return fas.Names
}

/*
*
產生指定長度map
*/
func (fas *FAS) NewMap(len int) error {
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

func (fas *FAS) GetCurrentMap() map[string]int {
	return M
}

func (fas *FAS) StartCounter() error {
	*(&count) = 0
	timer = time.NewTimer(time.Second * 3)
	for {
		select {
		case <-timer.C:
			*(&count) = *(&count) + 1
			//fmt.Println(count)
			timer.Reset(time.Second * 3)
		}

	}
}

func (fas *FAS) StopCounter() error {
	if timer.Stop() {
		fmt.Println(">> 計數器關閉成功")
	}
	return nil
}

func (fas *FAS) GetCurrentCount() int {
	return count
}

func (fas *FAS) Panic() {
	panic("throw panic")
}

func (fas *FAS) Destruct() {
}
