package test

import (
	"fmt"
	"testing"
)

//数组
func TestArray(test *testing.T) {
	var a1 = [5]int{7, 7, 7, 9, 9}
	for k, v := range a1 {
		fmt.Println("k:", k, "\t", "v:", v)
	}

	var b1 = [...]int{7, 7, 7, 7, 7, 7}
	fmt.Println("len(b1):", len(b1))

	//通过索引赋值
	//指定索引为1和2的值分别为7和9
	var c1 = [5]int{1: 7, 2: 9}
	for k, v := range c1 {
		fmt.Println("k:", k, "\t", "v:", v)
	}

	//指针数组
	d1 := [5]*int{0: new(int)}
	fmt.Println("d1:", d1)
	*d1[0] = 1234
	fmt.Println("d1:", *d1[0])

	d3 := &[2]int{1, 2}
	var d4 = &[2]int{}
	*d4 = *d3
	fmt.Println("d4", &d4)
	fmt.Println("d4:", *d4)

	//多维数组
	var e1 [4][2]int
	e1[0][0] = 1
	e1[0][1] = 2
	fmt.Println("e1:", e1)

}
