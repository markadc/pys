package spider

import (
	"fmt"
	"strings"
)

// 根据params构建完整的URL
func MakeURL(baseURL string, params S) string {
	if !strings.Contains(baseURL, "?") {
		baseURL += "?"
	} else {
		baseURL += "&"
	}
	for key, value := range params {
		baseURL += fmt.Sprintf("%s=%s&", key, value)
	}
	return baseURL[:len(baseURL)-1]
}
