package main

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

func (nas *NAS) NewMap(len int) error {
	return nil
}

func (nas *NAS) GetCurrentMap() map[string]int {
	return nil
}

func (nas *NAS) StartCounter() error {
	return nil
}

func (nas *NAS) StopCounter() error {
	return nil
}

func (nas *NAS) GetCurrentCount() int {
	return 0
}

func (nas *NAS) Panic() {

}

func (nas *NAS) Destruct() {
}
