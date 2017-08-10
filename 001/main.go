/** 
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/8/10 16:56
  */

package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func solution(line string) string {
	// 在此处理单行数据
	s := strings.Split(line, " ")
	ret1, _ := strconv.Atoi(s[0])
	ret2, _ := strconv.Atoi(s[1])
	ret := strconv.Itoa(ret1 + ret2)
	// 返回处理后的结果
	return ret
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}
