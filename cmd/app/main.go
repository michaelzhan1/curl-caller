package main

import (
	"fmt"

	"github.com/michaelzhan1/curlcaller/internals/curl"
)

func main() {
	res, err := curl.Get("http://localhost:3000")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}