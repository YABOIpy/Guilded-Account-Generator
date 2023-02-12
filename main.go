package main

import (
	"fmt"
	"gen/src"
	"strconv"
	"sync"
)

var (
	c = src.Guilded{}
)

func main() {

	var wg sync.WaitGroup
	var threads int
	fmt.Print("[threads]>: ")
	fmt.Scanln(&threads)
	wg.Add(threads)
	for i := 0; i < threads; i++ {
		go func(i int) {
			defer wg.Done()
			token, str, err := c.GetAuth()
			c.Err(err)
			if len(token) > 0 {
				data := str + token
				c.WriteFile("tokens.txt", data)
			} else {
				fmt.Println("Failed")
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("generated: " + strconv.Itoa(threads) + " tokens")
}
