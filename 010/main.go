/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/8/11 14:24 
  */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution(line string) string {
	// 在此处理单行数据
	n, _ := strconv.Atoi(line)
	if n==0 || n==1 || n==2 {
		return strconv.Itoa(n)
	}

	ret := make([]int, 100)
	ret[0] = 1
	ret[1] = 2

	i := 2
	for i = 2; i < n; i++ {
		ret[i] = ret[i-2] + ret[i-1]
	}
	// 返回处理后的结果
	return strconv.Itoa(ret[i-1])
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}