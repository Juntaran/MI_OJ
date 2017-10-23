/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/10/23 15:19
  */

package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"math"
)

// 两个正整数相减，num1 - num2  如果 num1 < num2 会先算 num2 - num1 再加负号
func bigReduce(num1, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	lenMax := int(math.Max(float64(len1), float64(len2)))
	var ret []int
	var result string

	// 判断 num1 和 num2 大小
	var judge int  = 0
	if len1 > len2 {
		judge = 1
	} else if len1 < len2 {
		judge = -1
	} else {
		// len1 == len2
		for i := 0; i < len1; i++ {
			// 从左向右逐位判断 num1 和 num2 谁大
			if num1[i] == num2[i] {
				continue
			} else if num1[i] > num2[i] {
				judge = 1
				break
			} else {
				judge = -1
				break
			}
		}
	}
	switch judge {
	case 0:			// num1 == num2
		return "0"
	case 1:			// num1 > num2
		// do nothing
		break
	case -1:		// num1 < num2
		num1, num2 = num2, num1
		len1, len2 = len2, len1
		result = "-"
	}

	var a = make([]uint8, lenMax+1)
	var b = make([]uint8, lenMax+1)
	var c = make([]int, lenMax+1)

	for i := 0; i < len1; i++ {
		a[len1-i-1] = num1[i] - '0'
	}
	for i := 0; i < len2; i++ {
		b[len2-i-1] = num2[i] - '0'
	}

	for i := 0; i < lenMax; i++ {
		c[i] = int(a[i]) - int(b[i]) + c[i]
		if c[i] < 0 {
			c[i] += 10
			c[i+1] --
		}
	}
	var mark int
	for i := lenMax - 1; i >= 0; i-- {
		if c[i] != 0 {
			mark = i
			break
		}
	}
	for j := mark; j >= 0; j-- {
		ret = append(ret, c[j])
	}
	for _, v := range ret {
		result += strconv.Itoa(v)
	}

	return result
}

func solution(line string) string {
	// 在此处理单行数据
	s := strings.Split(line, " - ")
	num1 := s[0]
	num2 := s[1]

	// 返回处理后的结果
	return bigReduce(num1, num2)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}