package main
import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)
func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go routines:",runtime.NumGoroutine())
	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i :=0; i< gs; i++ {
		go func() {
			atomic.AddInt64(&counter,1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			time.Sleep(time.Second)
			runtime.Gosched()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Go routine", runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}