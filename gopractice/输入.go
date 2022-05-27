package gopractice

import (
	"fmt"
)

func input() {
	a := 0
	b := 0
	for {
		n, _ := fmt.Scan(&a, &b)
		fmt.Println("输入了几个数？", n)
		if n == 0 {
			break
		} else {
			fmt.Printf("%d\n", a+b)
		}
	}
}
