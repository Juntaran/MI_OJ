/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/10/23 16:15
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
	s := strings.Split(line, ",")
	length := len(s)
	nums := make([]int, length)
	for i, v := range s {
		nums[i], _ = strconv.Atoi(v)
	}

	mark := make(map[int]bool)
	for i := 0; i < length; i++ {
		mark[nums[i]] = false
	}

	var count = 0

	for i := 0; i < length; i++ {
		if mark[nums[i]] == true {
			continue
		}
		var temp int = 1
		for j := nums[i] + 1; ; j++ {
			if _, ok := mark[j]; ok {
				mark[j] = true
				temp ++
			} else {
				break
			}
		}

		for j := nums[i] - 1; ; j -- {
			if _, ok := mark[j]; ok {
				mark[j] = true
				temp ++
			} else {
				break
			}
		}

		if count < temp {
			count = temp
		}
	}

	// 返回处理后的结果
	return strconv.Itoa(count)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}