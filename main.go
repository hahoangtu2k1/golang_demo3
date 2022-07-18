package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var mu = sync.Mutex{}

func main() {
	fmt.Println("----------------------------------bài1-----------------------------------")
	fmt.Println("")
	routineWaitGroup()
	fmt.Println("")
	goMutex()
	fmt.Println("")
	chanRoutine()
	fmt.Println("")
	fmt.Println("----------------------------------bài2-----------------------------------")
	fmt.Println("")
	mapMutex()
	fmt.Println("")
	fmt.Println("----------------------------------bài3-----------------------------------")
	fmt.Println("")
	fmt.Println("")
	errFunc()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("----------------------------------bài4-----------------------------------")
	fmt.Println("")
	test4()
}
