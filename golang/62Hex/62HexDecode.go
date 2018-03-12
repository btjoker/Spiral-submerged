package main

import (
	"math"
	"strings"
)

const base = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var base64 = map[rune]int64{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'a': 10, 'b': 11, 'c': 12, 'd': 13, 'e': 14, 'f': 15, 'g': 16, 'h': 17, 'i': 18, 'j': 19, 'k': 20, 'l': 21, 'm': 22, 'n': 23, 'o': 24, 'p': 25, 'q': 26, 'r': 27, 's': 28, 't': 29, 'u': 30, 'v': 31, 'w': 32, 'x': 33, 'y': 34, 'z': 35, 'A': 36, 'B': 37, 'C': 38, 'D': 39, 'E': 40, 'F': 41, 'G': 42, 'H': 43, 'I': 44, 'J': 45, 'K': 46, 'L': 47, 'M': 48, 'N': 49, 'O': 50, 'P': 51, 'Q': 52, 'R': 53, 'S': 54, 'T': 55, 'U': 56, 'V': 57, 'W': 58, 'X': 59, 'Y': 60, 'Z': 61}

// Base62ToBase10 62 进制转 10 进制
func Base62ToBase10(str string) (result int64) {
	str = strings.TrimLeft(strings.TrimSpace(str), "0")
	length := len(str) - 1
	for index, char := range str {
		result += base64[char] * int64(math.Pow(62, float64(length-index)))
	}
	return
}

// Base10ToBase62 10 进制转  62 进制
func Base10ToBase62(number int64) string {
	if number == 0 {
		return "0"
	}
	var result []byte
	for number > 0 {
		result = append(result, base[number%62])
		number = int64(number / 62)
	}
	return ReverseString(result)
}

// ReverseString 反转字符串
func ReverseString(strArr []byte) string {
	tail := len(strArr) - 1
	for i, l := 0, len(strArr)/2; i < l; i++ {
		strArr[i], strArr[tail-i] = strArr[tail-i], strArr[i]
	}
	return string(strArr)
}
