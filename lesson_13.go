/**
罗马数字转整数
*/
package main

import "fmt"

var m = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) int {
	var res int
	l := len(s)
	for i := range s {
		val := m[s[i]]
		if i < l-1 && val < m[s[i+1]] {
			res -= val
		} else {
			res += val
		}
	}
	return res
}

func main() {
	var s = "MCMXCIV"
	fmt.Println(romanToInt(s))
}
