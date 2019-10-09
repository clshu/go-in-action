package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	w := newWords()

	for _, fname := range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			if err := tallyWords(filename, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(fname)
	}
	wg.Wait()
	fmt.Println("Words that appear more than once:", i+1)

}

type words struct {
	found map[string]int
}

// Create a new words instance
func newWords() *words {
	return &words{found: map[string]int()}
}
