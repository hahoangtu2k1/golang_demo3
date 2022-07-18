package main

import (
	"fmt"
	"log"
	"time"
)

var chann = make(chan bool)

func routineWaitGroup() {
	fmt.Println("RoutineWaitGroup")
	wg.Add(1)
	log.Print("hello 1")
	go func() {
		time.Sleep(100 * time.Millisecond)
		log.Print("hello 3")
		wg.Done()
	}()
	log.Print("hello 2")
	wg.Wait()
}
func goMutex() {
	fmt.Println("go Mutex")
	mu.Lock()
	log.Print("hello 1")
	go func() {
		time.Sleep(100 * time.Millisecond)
		log.Print("hello 3")
		mu.Unlock()
	}()
	mu.Lock()
	log.Print("hello 2")
	mu.Unlock()
}
func chanRoutine() {
	fmt.Println("chanRoutine")
	log.Print("hello 1")
	go func() {
		time.Sleep(100 * time.Millisecond)
		log.Print("hello 3")
		chann <- true
	}()
	<-chann
	log.Print("hello 2")

}
