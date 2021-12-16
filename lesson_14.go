/*
最长公共前缀
*/
package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		var length int
		if len(prefix) < len(strs[i]) {
			length = len(prefix)
		} else {
			length = len(strs[i])
		}

		index := 0
		for index < length && prefix[index] == strs[i][index] {
			index++
		}
		prefix = strs[i][:index]
	}
	return prefix
}

func main() {
	//strs := []string{"flower", "flow", "flight"}
	strs1 := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs1))
}
