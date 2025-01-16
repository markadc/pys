package main

import (
	. "CSDN/pys"
	"encoding/json"
	"fmt"
)

type Data struct {
	Scene      string `json:"scene"`
	Size       int    `json:"size"`
	SearchText string `json:"search_text"`
	Cursor     int    `json:"cursor"`
	Extra      S      `json:"extra"`
	Filters    A      `json:"filters"`
}

func main() {
	//data := A{"scene": "PCSquareFeed", "size": 30, "search_text": "", "cursor": 0, "extra": S{"new_session_strategy": "1", "search_id": "", "session_id": ""}, "filters": S{}}

	extra := S{
		"new_session_strategy": "1",
		"search_id":            "",
		"session_id":           "",
	}
	filters := A{}

	data := Data{
		Scene:      "PCSquareFeed",
		Size:       30,
		SearchText: "",
		Cursor:     0,
		Extra:      extra,
		Filters:    filters,
	}

	bs, _ := json.Marshal(data)
	jsonStr := fmt.Sprintf("`%s`", bs)
	fmt.Println(jsonStr)

	a, _ := json.Marshal(nil)
	fmt.Println(string(a))
}
