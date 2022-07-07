/*package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var createChannel = make(chan string, 10)

func test4() {
	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Print(err)
	}
		createChannel <- string(data)
		fmt.Println(<-createChannel)
		for fileText := <-createChannel; fileText < string(data); {

		}

}*/
package main

import (
	"bufio"
	"log"
	"os"
	"sync"
	"time"
)

type DataLine struct {
	content   string
	indentity int
}

func test4() {
	buffReadData := make(chan *DataLine, 10)
	// defer close(buffReadData)
	var wg sync.WaitGroup
	f, _ := os.Open("file.txt")
	defer f.Close()
	numberWorker := 3
	done := make([]chan bool, numberWorker)
	for i := 0; i < numberWorker; i++ {
		done[i] = make(chan bool)
		go printData(buffReadData, &wg, done[i])
	}

	scanner := bufio.NewScanner(f)

	count := 1

	for scanner.Scan() {
		wg.Add(1)
		dataLine := &DataLine{content: scanner.Text(), indentity: count}
		count++
		buffReadData <- dataLine

	}

	wg.Wait()
	for _, x := range done {
		x <- true

	}
	time.Sleep(10 * time.Second)
}

func printData(jobs chan *DataLine, wg *sync.WaitGroup, done chan bool) {
	for {
		select {
		case data := <-jobs:
			log.Printf("Hang %v : %v xong!\n", data.indentity, data.content)
			wg.Done()
		case <-done:
			log.Println("Đã xong Worker")
			break
		}
	}

}
