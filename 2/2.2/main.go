package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var time_start = time.Now()
	if len(os.Args) > 1 {
		var search_for string = os.Args[1]
		for i := 2; i < len(os.Args); i++ {
			wg.Add(1)
			go func() {
				outputFile(os.Args[i], &search_for)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	fmt.Println("Total time:", time.Now().Sub(time_start))

}

func outputFile(fileName string, search_for *string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return
	}
	if strings.Contains(string(data), *search_for) {
		fmt.Println(fileName, " contains subsstring ", *search_for)
	}

}
