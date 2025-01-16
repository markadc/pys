package main

import (
	"CSDN/pys"
	"fmt"
)

func main() {
	headers := pys.S{"Cookie": "abcd", "User-Agent": "UA" }
	response, _ := pys.GET("https://blog.csdn.net/MarkAdc", headers, nil)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Request.Header)
}
