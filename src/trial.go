package main

import (
	"fmt"
	"sync"
)

func main() {
	cont := 0
	v := 5
	var m sync.Mutex
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			m.Lock()
			v++
			cont = v
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(cont)
}
