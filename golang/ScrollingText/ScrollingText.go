package main

import "strings"

func main() {}

// ScrollingText 滚动字符
// 思路就是追加第一个字符, 然后切片
func ScrollingText(text string) (result []string) {
	item := []byte(strings.ToUpper(text))

	for range item {
		result = append(result, string(item))
		item = append(item, item[0])
		item = item[1:]
	}
	return
}