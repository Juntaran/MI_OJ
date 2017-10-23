/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/10/23 18:18
  */

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func solution(line string) string {
	// 在此处理单行数据
	s := strings.Split(line, ",")
	length := len(s)
	if length == 1 {
		return s[0]
	}

	nums := make([]int, length)
	for i, v := range s {
		nums[i], _ = strconv.Atoi(v)
	}

	middle := length / 2 + 1
	var max_id int
	var min_id int
	var judge int

	for i := 0; i < length-1; i++ {
		if nums[i] > nums[i+1] {
			max_id = i
			min_id = i + 1
			judge = 1
		}
	}

	// 返回处理后的结果
	if judge == 0 {
		// 没有旋转
		return strconv.Itoa(nums[middle - 1])
	}

	last := length - min_id
	if last < middle {
		return strconv.Itoa(nums[middle - last - 1])
	}
	if last > middle {
		return strconv.Itoa(nums[middle + max_id])
	}
	return strconv.Itoa(nums[length - 1])
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}