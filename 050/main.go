/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/8/10 22:23 
  */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func solution(line string) string {
	// 在此处理单行数据
	s := strings.Split(line, ",")
	length := len(s)

	heights := make([]int, length)
	for i := 0; i < length; i++ {
		heightI, _ := strconv.Atoi(s[i])
		heights[i] = heightI
	}

	left  := make([]int, length)
	right := make([]int, length)

	left[0] = 0
	right[length-1] = length - 1


	for i := 1; i < length; i++ {
		temp := i
		for {
			if temp > 0 && heights[temp-1] >= heights[i] {
				temp = left[temp - 1]
			} else {
				break
			}
		}
		left[i] = temp
	}

	for i := length-2; i >= 0; i-- {
		temp := i
		for {
			if temp < length-1 && heights[temp+1] >= heights[i] {
				temp = right[temp + 1]
			} else {
				break
			}
		}
		right[i] = temp
	}

	ret := 0
	for i := 0; i < length; i++ {
		ret = max(ret, heights[i] * (right[i] - left[i] + 1))
	}

	// 返回处理后的结果
	return strconv.Itoa(ret)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}