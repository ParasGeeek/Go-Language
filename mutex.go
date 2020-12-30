//#https://godoc.org/sync#Mutex.Lock
package main
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)
func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go routines:",runtime.NumGoroutine())
	counter :=0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex
	for i :=0; i< gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			time.Sleep(time.Second)
			runtime.Gosched()
			v++;
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Go routine", runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}