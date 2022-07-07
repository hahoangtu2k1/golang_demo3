package main

import (
	"fmt"
)

func mapMutex() {
	var X = make(map[string]string)
	X["name"] = "Hà Hoàng Tú"
	for i := 0; i <= 1000; i++ {

		fmt.Println(X)
	}
	wg.Done()

}
