package main

import "strconv"

func main() {
	if matchLuhn("6550020001000003287") {
		println("有效")
	} else {
		println("无效")
	}
}

// matchLuhn 银行卡号验证Luhn算法
func matchLuhn(cardNo string) bool {
	var sum int

	cardNoLength := len(cardNo)
	cardNoArr := make([]int, cardNoLength)

	// 转为数字数组
	for k, v := range cardNo {
		cardNoArr[k], _ = strconv.Atoi(string(v))
	}

	// 从右往左 偶数位乘2
	for i := cardNoLength - 2; i >= 0; i -= 2 {
		cardNoArr[i] <<= 1
		cardNoArr[i] = cardNoArr[i]/10 + cardNoArr[i]%10
	}

	// 总和
	for i := 0; i < cardNoLength; i++ {
		sum += cardNoArr[i]
	}
	return sum%10 == 0
}
