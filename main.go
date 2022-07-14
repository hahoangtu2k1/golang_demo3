package main

import "fmt"

func main() {
	fmt.Println("----------------------bài1------------------")
	routineWaitGroup()
	goMutex()
	chanRoutine()
	fmt.Println("----------------------bài2------------------")
	// mu.Lock()
	// wg.Add(1)
	go mapMutex()
	// go mapMutex()
	// go mapMutex()
	// wg.Wait()
	// mu.Unlock()
	fmt.Println("----------------------bài3------------------")
	errFunc()
	fmt.Println("----------------------bài4------------------")
	test4()
}
