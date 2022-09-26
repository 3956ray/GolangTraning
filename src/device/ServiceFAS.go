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
	//view.count = 0
	////每三秒計數一次
	//view.timer = time.NewTimer(time.Second * 3)
	//
	//select {
	//case <-view.timer.C:
	//	view.count++
	//}
	return nil
}

func (fas *FAS) StopCounter() error {
	//if view.timer.Stop() {
	//	fmt.Println(">> 計數器關閉成功")
	//} else {
	//	fmt.Println(">> 計數器未啟動")
	//}
	return nil
}

func (fas *FAS) GetCurrentCount() int {
	//return view.count
	return 0
}

func (fas *FAS) Panic() {
	panic("throw panic")
}

func (fas *FAS) Destruct() {
}
