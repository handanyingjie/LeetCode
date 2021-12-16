/*
判断字符是否唯一
按位与 & 只有两个操作数对应位同为1时，结果为1，其余全为0。（或者是只要有一个操作数为0，结果就为0）
按位或 ｜只有两个操作数对应位同为0时，结果为0，其余全为1（或者是只要有一个操作数为1，结果就为1）
按位异或 ^ 同0，异1

*/
package main

import "fmt"

func isUniqueChar(str string) bool {
	mark := 0
	for _, v := range str {
		moveBit := v - 'a'
		// 第一次到这个位置 mark 的二进制位全为0,这个时候v恰好也是a,则 'a' - 'a' 为0，
		// 则moveBit二进制位全为0， 0 & 0 = 0 则走else, 0 ｜ 0 还为0，所以需要左移一位，即1 << moveBit

		if (mark & (1 << moveBit)) != 0 {
			return false
		} else {
			mark |= 1 << moveBit
		}
	}
	return true
}

func main() {
	str := "abc"
	fmt.Println(isUniqueChar(str))
}
