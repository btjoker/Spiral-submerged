package main

import (
	"math/rand"
	"strings"
	"testing"
)

func Test_ScrollingText(t *testing.T) {
	// 随机两个数字, 然后用 slice 切出字符串, 然后组合成字符串判断是否相等
	testText := "abcdefghijkmlnopqrstuvwxyz"
	testHelp := func(text string) (result []string) {
		item := []byte(strings.ToUpper(text))

		for range item {
			result = append(result, string(item))
			item = append(item, item[0])
			item = item[1:]
		}
		return
	}

	swap := func(x, y int) (int, int) {
		if x > y {
			return y, x
		}
		return x, y
	}

	for range testText {
		t.Run("", func(t *testing.T) {
			x, y := swap(rand.Intn(25), rand.Intn(25))
			item := testText[x:y]
			result := strings.Join(ScrollingText(item), ",")
			want := strings.Join(testHelp(item), ",")
			if result != result {
				t.Errorf("ScrollingText() = %v, want %v", result, want)
			}
		})
	}
}

func Benchmark_ScrollingText(b *testing.B) {
	b.StopTimer()
	item := "abcdefghijkmlnopqrstuvwxyz"
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ScrollingText(item)
	}
}
