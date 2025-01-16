package main

import (
	"fmt"
	_url "net/url"
)

func BuildURL(baseURL string, params map[string]string) string {
	u, _ := _url.Parse(baseURL)
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func main() {
	url := "https://www.google.com"
	params := map[string]string{"name": "CLOS", "age": "22"}
	aa := BuildURL(url, params)
	fmt.Println(aa)
	//fmt.Println(999)
	//pys, err := http.NewRequest("GET", "http://www.bcccaiduc.com", nil)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//values := pys.URL.Query()
	//params := map[string]string{
	//	"name": "CLOS",
	//	"age":  "22",
	//}
	//for k, v := range params {
	//	values.Add(k, v)
	//}
	//pys.URL.RawQuery = values.Encode()
	//
	//reqInfo := pys.RequestInfo{}
	//fmt.Println(reqInfo.Header)
	//fmt.Println(reqInfo.Params)
}
