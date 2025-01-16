package main

import (
	"CSDN/py"
	"fmt"
)

func main() {
	headers := py.S{"Cookie": "abcd", "User-Agent": "UA" }
	response, _ := py.GET("https://blog.csdn.net/MarkAdc", headers, nil)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Request.Header)
}
