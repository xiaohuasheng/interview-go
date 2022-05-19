package gopractice

import (
	"fmt"
)

type BitMap []byte

const byteSize = 8 //定义的bitmap为byte的数组，byte为8bit

func NewBitMap(n uint) BitMap {
	return make([]byte, n/byteSize+1)
}
func (bt BitMap) set(n uint) {
	if n/byteSize > uint(len(bt)) {
		fmt.Println("大小超出bitmap范围")
		return
	}
	byteIndex := n / byteSize   //第x个字节（0,1,2...）
	offsetIndex := n % byteSize //偏移量(0<偏移量<byteSize)
	//bt[byteIndex] = bt[byteIndex] | 1<<offsetIndex //异或1（置位）
	//第x个字节偏移量为offsetIndex的位 置位1
	bt[byteIndex] |= 1 << offsetIndex //异或1（置位）
}
func (bt BitMap) del(n uint) {
	if n/byteSize > uint(len(bt)) {
		fmt.Println("大小超出bitmap范围")
		return
	}
	byteIndex := n / byteSize
	offsetIndex := n % byteSize
	bt[byteIndex] &= 0 << offsetIndex //清零
}
func (bt BitMap) isExist(n uint) bool {
	if n/byteSize > uint(len(bt)) {
		fmt.Println("大小超出bitmap范围")
		return false
	}
	byteIndex := n / byteSize
	offsetIndex := n % byteSize
	//fmt.Println(bt[byteIndex] & (1 << offsetIndex))
	return bt[byteIndex]&(1<<offsetIndex) != 0 //TODO：注意：条件是 ！=0，有可能是：16,32等
}
func main2() {
	//a1 := byte('a')
	//fmt.Println(a1)
	bitMap := NewBitMap(20)
	bitMap.set(20)
	bitMap.set(1)
	bitMap.set(2)
	bitMap.set(3)
	bitMap.set(4)
	bitMap.set(5)
	fmt.Println("20:", bitMap.isExist(20))
	fmt.Println("5:", bitMap.isExist(5))
	bitMap.del(20)
	fmt.Println("20:", bitMap.isExist(20))
	fmt.Println("5:", bitMap.isExist(5))
	fmt.Println("1:", bitMap.isExist(1))
	fmt.Println("3:", bitMap.isExist(3))
	fmt.Println("6:", bitMap.isExist(6))
	fmt.Println("7:", bitMap.isExist(7))
}
