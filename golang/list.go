package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// List .
type List []int

// Len 方法返回长度
func (list List) Len() int {
	return len(list)
}

// Less 方法返回 j 下标 是否小于 k 下标的值
func (list List) Less(j, k int) bool {
	return list[j] < list[k]
}

// Swap 方法交换 j 下标 k 下标的值
func (list List) Swap(j, k int) {
	list[j], list[k] = list[k], list[j]
}

// Filter 方法创建一个新数组, 其包含通过所提供函数实现的测试的所有元素。
func (list List) Filter(fn func(int) bool) List {
	result := make(List, 0, list.Len())
	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map 方法创建一个新数组，其结果是该数组中的每个元素都调用一个提供的函数后返回的结果。
func (list List) Map(fn func(int) int) List {
	result := make(List, 0, list.Len())
	for _, v := range list {
		result = append(result, fn(v))
	}
	return result
}

// ForEach 方法对数组的每个元素执行一次提供的函数。
func (list List) ForEach(fn func(int)) {
	for _, v := range list {
		fn(v)
	}
}

// Sum 方法返回总值
func (list List) Sum() int {
	var result int
	for _, v := range list {
		result += v
	}
	return result
}

// Reduce 方法对累加器和数组中的每个元素（从左到右）应用一个函数，将其减少为单个值。
func (list List) Reduce(fn func(int, int) int) int {
	var result int
	for _, v := range list {
		result = fn(result, v)
	}
	return result
}

// Every 方法测试数组的所有元素是否都通过了指定函数的测试。
func (list List) Every(fn func(int) bool) bool {
	for _, v := range list {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Some 方法测试数组中的某些元素是否通过由提供的函数实现的测试。
func (list List) Some(fn func(int) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
	}
	return false
}

// Reverse 方法将数组中元素的位置颠倒。
func (list List) Reverse() List {
	result := make(List, 0, list.Len())
	for i := list.Len() - 1; i >= 0; i-- {
		result = append(result, list[i])
	}
	return result
}

// Push 方法将一个或多个元素添加到数组的末尾。
func (list List) Push(i int) {
	list = append(list, i)
}

// Pop 方法从数组中删除最后一个元素，并返回该元素的值。
func (list List) Pop() int {
	result := list[list.Len()-1]
	list = list[:list.Len()-1]
	return result
}

// Shift 方法从数组中删除第一个元素，并返回该元素的值。
func (list List) Shift() int {
	result := list[0]
	list = list[1:]
	return result
}

// UnShift 方法将一个或多个元素添加到数组的开头，并返回新数组的长度。
func (list List) UnShift(i int) {
	result := make(List, 0, list.Len())
	result = append(result, i)
	result = append(result, list...)
	list = result
}

// Concat 方法用于合并两个或多个数组。此方法不会更改现有数组，而是返回一个新数组。
func (list List) Concat(item List) List {
	result := make(List, 0, list.Len()*2)
	result = append(result, list...)
	result = append(result, item...)
	return result
}

// Includes 方法用来判断一个数组是否包含一个指定的值，如果是，酌情返回 true或 false。
func (list List) Includes(i int) bool {
	for _, v := range list {
		if v == i {
			return true
		}
	}
	return false
}

// Keys 方法返回一个新的Array迭代器，它包含数组中每个索引的键。
func (list List) Keys() List {
	result := make(List, 0, list.Len())
	for k := range list {
		result = append(result, k)
	}
	return result
}

// IndexOf 方法返回在数组中可以找到一个给定元素的第一个索引，如果不存在，则返回-1。
func (list List) IndexOf(i int) int {
	for k, v := range list {
		if v == i {
			return k
		}
	}
	return -1
}

// LastIndexOf 方法返回指定元素在数组中的最后一个的索引，如果不存在则返回 -1。
func (list List) LastIndexOf(v int) int {
	for i := list.Len() - 1; i >= 0; i-- {
		if list[i] == v {
			return i
		}
	}
	return -1
}

// FindIndex 方法返回数组中满足提供的测试函数的第一个元素的索引。否则返回-1。
func (list List) FindIndex(fn func(int) bool) int {
	for k, v := range list {
		if fn(v) {
			return k
		}
	}
	return -1
}

// Fill 方法用一个固定值填充一个数组中从起始索引到终止索引内的全部元素, 改变原数组。
func (list List) Fill(i int) {
	for k := range list {
		list[k] = i
	}
}

// Join 方法将数组（或一个类数组对象）的所有元素连接到一个字符串中。
func (list List) Join(sep string) string {
	result := make([]string, 0, list.Len())
	for _, v := range list {
		result = append(result, strconv.Itoa(v))
	}
	return strings.Join(result, sep)
}

// ToArray 方法返回一个int数组
func (list List) ToArray() []int {
	return []int(list)
}

// ToString 方法返回字符串形式数组
func (list List) ToString() string {
	return "[" + list.Join(",") + "]"
}

// NewList 方法返回一个List实例
func NewList(list []int) List {
	return List(list)
}

func main() {
	list := List(rand.Perm(15))

	every := list.Every(func(i int) bool {
		if i%2 == 0 {
			return true
		}
		return false
	})

	some := list.Some(func(i int) bool {
		if i == 0 {
			return true
		}
		return false
	})

	filter := list.Filter(func(i int) bool {
		if i > 2 && i != 2 && i != 3 && i != 5 {
			return true
		}
		if i%2 != 0 && i%3 != 0 && i%5 != 0 {
			return true
		}
		return false
	})

	findIndex := list.FindIndex(func(i int) bool {
		if i*i*i < 10 {
			return true
		}
		return false
	})

	list.ForEach(func(i int) {
		fmt.Println("ForEach: ", i)
	})

	listMap := list.Map(func(i int) int {
		return i * i
	})

	link := list.Concat(NewList(rand.Perm(10)).Map(func(i int) int {
		return i * 3
	}).Filter(func(i int) bool {
		if i%9 == 0 {
			return true
		}
		return false
	}))

	fmt.Println("List: ", list)
	fmt.Println("Every: ", every)
	fmt.Println("Some: ", some)
	fmt.Println("FindIndex: ", findIndex)
	fmt.Println("LastIndexOf: ", list.LastIndexOf(3))
	fmt.Println("Includes: ", list.Includes(8))
	fmt.Println("IndexOf: ", list.IndexOf(9))

	fmt.Println("Reverse: ", list.Reverse())
	fmt.Println("Keys: ", list.Keys())
	fmt.Println("Concat:", list.Concat(List(rand.Perm(5))))
	fmt.Println("Join: ", list.Join(", "))
	fmt.Println("Map: ", listMap)
	fmt.Println("Filter: ", filter)
	fmt.Println("ToString: ", list.ToString())
	fmt.Println("ToArray: ", list.ToArray())
	fmt.Println("link: ", link)

	list.Fill(9)
	fmt.Println("Fill: ", list)
}
