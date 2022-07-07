package main

func main() {
	routineWaitGroup()
	goMutex()
	chanRoutine()
	mu.Lock()
	wg.Add(3)
	go mapMutex()
	go mapMutex()
	go mapMutex()
	wg.Wait()
	mu.Unlock()
	errFunc()
	test4()
}
