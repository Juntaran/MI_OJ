/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/10/23 11:16
  */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solution(line string) string {
	// 在此处理单行数据
	s := strings.Split(line, ",")

	len1 := len(s[0])
	len2 := len(s[1])

	if len(s) != 2 || len1 != len2 {
		return "0"
	}

	// 新建两个 slice
	a := make([]int32, 256)
	b := make([]int32, 256)

	for i := 0; i < len1; i++ {
		if a[s[0][i]] != b[s[1][i]] {
			return "0"
		}
		a[s[0][i]] = int32(i) + 1
		b[s[1][i]] = int32(i) + 1
	}

	// 返回处理后的结果
	return "1"
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}