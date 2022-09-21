package main

import (
	"fmt"
	"time"
)

var TimeCount *time.Timer

//type treeNode struct {
//	value       int
//	left, right *treeNode
//}
//
//var root treeNode
//
//func createTreeNode(value int) *treeNode {
//	return &treeNode{value: value}
//}
//
//func Newtr(len int) ITreenode {
//	return &root
//}
//
//func (t *treeNode) Type(len int) {
//	fmt.Println(t.value)
//}
//
//type ITreenode interface {
//	Type(len int)
//}

func main() {

	//for {
	//	//	顯示菜單
	//	if !menu() {
	//		break
	//	}
	//	//	選取菜單
	//	//menuSelect()
	//	//	顯示選項菜單
	//
	//}
	TimeCount = time.NewTimer(time.Second)

	fmt.Println(TimeCount)
	TimeCount.Stop()
	fmt.Println(TimeCount.C)
}
