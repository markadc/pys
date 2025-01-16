package pys

import (
	"fmt"
	"testing"
)

func TestGget(t *testing.T) {
		data := A{
			"level1.2": 404,
			"level1": A{
				"level2": A{
					"level3": 3,
				},
				"level2.1": []int{1, 2, 3},
			},
		}

		result := Nget(data, []string{"level1", "level2.1", "level3"}, "gg")
		fmt.Println(result)

		result = Nget(data, []string{"level1", "222"}, "gg")
		fmt.Println(result)

		result = Nget(data, []string{"level1.2", "222"}, "gg")
		fmt.Println(result)
}