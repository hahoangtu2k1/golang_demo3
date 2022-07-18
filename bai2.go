package main

import (
	"fmt"
	// "sync"
	"strconv"
)

var X = make(map[string]string)

func mapMutex() {
	wg.Add(3)
	go map1(X)
	go map2(X)
	go map3(X)
	wg.Wait()
	// count := 0
	for k, v := range X {
		fmt.Printf("Key: %v, value: %v\n", k, v)
		// count++

	}

	fmt.Println("hoàn thành bài 2")
}
func map1(X map[string]string) {
	m1 := "mapMutex(1)  "
	for i := 0; i < 1000; i++ {
		mu.Lock()
		create := m1 + strconv.Itoa(i)
		X[create] = strconv.Itoa(i)
		mu.Unlock()
	}
	wg.Done()
}
func map2(X map[string]string) {
	m2 := "mapMutex(2)  "
	for i := 0; i < 1000; i++ {
		mu.Lock()
		create := m2 + strconv.Itoa(i)
		X[create] = strconv.Itoa(i)
		mu.Unlock()
	}
	wg.Done()
}
func map3(X map[string]string) {
	m3 := "mapMutex(3)  "
	for i := 0; i < 1000; i++ {
		mu.Lock()
		create := m3 + strconv.Itoa(i)
		X[create] = strconv.Itoa(i)
		mu.Unlock()
	}
	wg.Done()
}
