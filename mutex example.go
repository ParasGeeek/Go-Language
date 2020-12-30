package main
import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	increment :=0
	gs :=100
	wg.Add(gs)
	var m sync.Mutex
	for i :=0; i< gs; i++ {
		go func() {
			m.Lock()
			v :=increment
			runtime.Gosched()
			v++
			increment=v
			fmt.Println(increment)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", increment)
}