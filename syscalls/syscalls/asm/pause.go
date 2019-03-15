package main

import (
	"fmt"
	"runtime"
)

// Pause starts OMIT
func Pause()

func main() {
	go func() {
		for i := 0; ; i++ {
			fmt.Println(i)
		}
	}()

	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		go Pause()
	}

	select {}
}

// END OMIT
