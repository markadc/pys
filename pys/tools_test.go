package pys

import (
	"fmt"
	"testing"
)

func TestMakeURL(t *testing.T) {
	baseURL := "https://www.baidu.com"
	params := map[string]string{"name": "CLOS", "age": "22"}
	fmt.Println(MakeURL(baseURL, params))

}


