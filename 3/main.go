package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency map[string]int, mu *sync.Mutex) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Server returning error status code: " + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	text := string(body)
	mu.Lock()
	for _, word := range strings.Fields(text) {
		if _, exists := frequency[word]; exists {
			frequency[word]++
		} else {
			frequency[word] = 1
		}
	}
	mu.Unlock()
	fmt.Println("Completed:", url)
}

func main() {
	var frequency = make(map[string]int)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var start_time = time.Now()
	for i := 1000; i <= 1050; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		wg.Add(1)
		go func() {
			countLetters(url, frequency, &mu)
			wg.Done()
		}()
	}
	wg.Wait()
	for key, value := range frequency {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	fmt.Println("Total time:", time.Now().Sub(start_time))
}
