package main

import (
	"log"
	"sync"
)

func errFunc() {
	m := make(map[int]int)
	mu := sync.Mutex{}
	for i := 0; i < 1000; i++ {

		go func() {
			for j := 1; j < 10000; j++ {
				mu.Lock()
				if _, ok := m[j]; ok {
					delete(m, j)
					continue
				}
				m[j] = j * 10
				mu.Unlock()
			}
		}()
	}

	log.Print("done")
}
