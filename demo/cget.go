package main

import (
	"fmt"
	"log"
)

// cget 实现了字典多层取值，KEY不存在则返回默认值
func cget(data map[string]any, args []string, failed any) any {
	temp := data
	for i, a := range args {
		// 检查 key 是否存在
		if val, exists := temp[a]; exists {
			// 如果这是最后一个 key，返回对应的值
			if i == len(args)-1 {
				return val
			}
			// 如果值是 map 类型，继续下一步，否则返回 failed
			if nestedMap, ok := val.(map[string]any); ok {
				temp = nestedMap
			} else {
				log.Printf("KEY %q VALUE %v not is map, %v", a, val, args)
				return failed
			}
		} else {
			log.Printf("KEY %q miss, %v", a, args)
			return failed
		}
	}
	return nil
}

func main() {
	// 示例数据
	type A = map[string]any

	data := A{
		"level1.2": 404,
		"level1": A{
			"level2": A{
				"level3": 3,
			},
			"level2.1": []int{1, 2, 3},
		},
	}

	result := cget(data, []string{"level1", "level2.1", "level3"}, "gg")
	fmt.Println(result)

	result = cget(data, []string{"level1", "222"}, "gg")
	fmt.Println(result)

	result = cget(data, []string{"level1.2", "222"}, "gg")
	fmt.Println(result)


}
