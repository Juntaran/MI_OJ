/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/8/10 19:47
  */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func solution(line string) string {
	// 在此处理单行数据
	s := strings.Split(line, " ")

	if len(s) == 1 {
		return s[0]
	}

	if len(s) <= 2 {
		return ""
	}

	var numSlice []int

	for _, each := range s {
		num, _ := strconv.Atoi(each)
		numSlice = append(numSlice, num)
	}

	ret := numSlice[0]

	for i := 1; i < len(numSlice); i++ {
		ret = ret ^ numSlice[i]
	}

	stringRet := strconv.Itoa(ret)

	// 返回处理后的结果
	return stringRet
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}