/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/10/23 20:35
  */

package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func solution(line string) string {
	// 在此处理单行数据
	n, _ := strconv.Atoi(line)

	// 第 1 3 6 10 15 ... 都是 1
	var origin, sum, temp, r int = 0, 0, 1, 1
	for i := 1; ; {
		temp += r
		if temp == 9 {
			r = -1
		}
		if temp == 1 {
			r = 1
		}
		i += r
		if sum > (n - 1) {
			origin = sum - i
			break
		}
	}
	ret := n - origin

	// 返回处理后的结果
	return strconv.Itoa(ret)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}