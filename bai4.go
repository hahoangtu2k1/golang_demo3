package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type Line struct {
	line_number int
	data        string
}

//ex4: bài tập worker pool: tạo bằng tay file dưới. file.txt sau đó đọc từng dòng file này nạp dữ liệu vào 1 buffer channel có size 10, Điều kiện đọc file từng dòng. Chỉ được sử dụng 3 go routine. Kết quả xử lý xong ỉn ra màn hình + từ xong
func test4() {
	ch := make(chan string, 10)
	finish := make(chan bool)
	access, deny := os.Open("file.txt")
	if deny != nil {
		log.Println(deny)
	}
	num := 0
	line := []*Line{}
	defer access.Close()
	scanner := bufio.NewScanner(access)
	for scanner.Scan() {
		ch <- scanner.Text()
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go chan_nel(finish, ch, scanner)
			wg.Done()
			wg.Wait()
		}
		time.Sleep(1 * time.Second)
		app := append(line, &Line{})
		num++
		for i := range app {
			fmt.Printf("%v giá trị là: %v\n", app[i].line_number, app[i].data)
		}
	}

	fmt.Println("xong")
}

func chan_nel(finish chan bool, ch chan string, scanner *bufio.Scanner) {
	for range scanner.Text() {
		fmt.Printf("Giá trị %v đã lấy ra\n", <-ch)
	}
	finish <- false
	close(finish)
	close(ch)
}

//var createChannel = make(chan string, 10)

// func test4() {
// 	wg := &sync.WaitGroup{}
// 	data, err := ioutil.ReadFile("file.txt")
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	fmt.Println(string(data))
// }

// func workPool() {

// }

// package main

// import (
// 	"bufio"
// 	"log"
// 	"os"
// 	"sync"
// 	"time"
// )

// type DataLine struct {
// 	content   string
// 	indentity int
// }

// func test4() {
// 	buffReadData := make(chan *DataLine, 10)
// 	// defer close(buffReadData)
// 	var wg sync.WaitGroup
// 	f, _ := os.Open("file.txt")
// 	defer f.Close()
// 	numberWorker := 3
// 	done := make([]chan bool, numberWorker)
// 	for i := 0; i < numberWorker; i++ {
// 		done[i] = make(chan bool)
// 		go printData(buffReadData, &wg, done[i])
// 	}

// 	scanner := bufio.NewScanner(f)

// 	count := 1

// 	for scanner.Scan() {
// 		wg.Add(1)
// 		dataLine := &DataLine{content: scanner.Text(), indentity: count}
// 		count++
// 		buffReadData <- dataLine

// 	}

// 	wg.Wait()
// 	for _, x := range done {
// 		x <- true

// 	}
// 	time.Sleep(10 * time.Second)
// }

// func printData(jobs chan *DataLine, wg *sync.WaitGroup, done chan bool) {
// 	for {
// 		select {
// 		case data := <-jobs:
// 			log.Printf("Hang %v : %v xong!\n", data.indentity, data.content)
// 			wg.Done()
// 		case <-done:
// 			log.Println("Đã xong Worker")
// 			break
// 		}
// 	}

// }
