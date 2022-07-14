package main

import (
	"fmt"
)

func mapMutex() {
	numFor := 3
	var X = make(map[string]string)
	X["name"] = "Hà Hoàng Tú"
	for j := 0; j <= numFor; j++ {
		for i := 0; i <= 1000; i++ {
			fmt.Println(X)
		}
	}
	wg.Done()

}
