package main

import (
	"fmt"

	"github.com/michaelzhan1/curlcaller/internals/curl"
)

func main() {
	res, err := curl.Call("http://localhost:3000", []string{}, []string{})
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}