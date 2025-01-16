package spider

import (
	"fmt"
	"testing"
)

func TestGenRandomUA(t *testing.T) {
	for i := 0; i < 5; i++ {
		ua := GenRandomUA()
		fmt.Println(ua)
	}
}
