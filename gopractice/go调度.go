package gopractice

import (
	"fmt"
	"sync"
	"time"
)

/*
go test -c go调度_test.go go调度.go -o gopractice
GOMAXPROCS=2 GODEBUG=schedtrace=1000,scheddetail=1 ./gopractice
*/

func main() {
	wg := sync.WaitGroup{}
	// for i := 0; i < 1000000; i++ {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			defer wg.Done()
			time.Sleep(100 * time.Microsecond)
		}(wg, 1)
	}

	wg.Done()
	fmt.Println("finished")
}
