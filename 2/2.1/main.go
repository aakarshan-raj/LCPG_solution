package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var time_start = time.Now()
	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			wg.Add(1)
			go func() {
				outputFile(os.Args[i])
				wg.Done()
			}()
		}
	}
	wg.Wait()
	var time_end = time.Now()
	fmt.Println("Time start:", time_start)
	fmt.Println("Time end:", time_end)

}

func outputFile(fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return
	}
	fmt.Println(string(data))
}
