/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/8/11 14:59 
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
	s := strings.Split(line, " ")

	len1 := len(s[0])
	len2 := len(s[1])

	if len1 > len2 {
		return "false"
	}

	var ascii = [128]int{0}

	for i := 0; i < len2; i++ {
		ascii[s[1][i]] ++
	}

	for i := 0; i < len1; i++ {
		ascii[s[0][i]] --
		if ascii[s[0][i]] < 0 {
			return "false"
		}
	}

	// 返回处理后的结果
	return "true"
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}
