package main

import (
	"fmt"
	"log"
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
		var entity, err = os.ReadDir(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		for _, e := range entity {
			wg.Add(1)
			go func() {
				outputFile(os.Args[2]+"/"+e.Name(), &search_for)
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
		log.Fatal(err)
	}
	if strings.Contains(string(data), *search_for) {
		fmt.Println(fileName, " contains subsstring ", *search_for)
	}

}
