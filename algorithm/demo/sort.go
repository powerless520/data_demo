package main

import (
	"fmt"
	"strings"
)

//给你一个字符串 s 和一个 长度相同 的整数数组 indices 。
//请你重新排列字符串 s ，其中第 i 个字符需要移动到 indices[i] 指示的位置。
//返回重新排列后的字符串。

func main() {
	//s = "codeleet", indices = [4,5,6,7,0,2,1,3]
	s := "codeleet"
	indics := []int{4, 5, 6, 7, 0, 2, 1, 3}

	fmt.Println(sortStr1(s, indics))
}

func sortStr(s string, indices []int) string {
	if len(s) != len(indices) {
		return ""
	}
	params := map[int]string{}
	for k, v := range indices {
		params[v] = s[k : k+1]
	}

	sortStr := ""
	for i := 0; i < len(s); i++ {
		sortStr += params[i]
	}

	return sortStr
}

func sortStr1(s string, indices []int) string {
	if len(s) != len(indices) {
		return ""
	}

	str := make([]string, len(indices))

	for k, _ := range indices {
		str[indices[k]] = s[k : k+1]
	}

	return strings.Join(str, "")
}
