/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/10/23 15:44
  */

package main

import (
	"math"
	"strconv"
	"bufio"
	"os"
	"fmt"
	"strings"
)

// 判断两个数大小
func whoBig(num1, num2 string) string {
	// 判断 num1 和 num2 大小
	// 传入默认 num1 大于 num2
	len1 := len(num1)
	len2 := len(num2)
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
	case 1:			// num1 > num2
		return "Y"
	default:		// num1 <= num2
		return "N"
	}
}

// 两个正整数相加
func bigAdd(num1, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	lenMax := int(math.Max(float64(len1), float64(len2)))
	var ret []uint8
	var result string

	var a = make([]uint8, lenMax + 1)
	var b = make([]uint8, lenMax + 1)
	var c = make([]uint8, lenMax + 1)

	for i := 0; i < len1; i++ {
		a[len1-i-1] = num1[i] - '0'
	}
	for i := 0; i < len2; i++ {
		b[len2-i-1] = num2[i] - '0'
	}

	for i := 0; i < lenMax; i++ {
		c[i] = a[i] + b[i]
	}

	for i := 0; i < lenMax; i++ {
		c[i+1] = c[i] / 10 + c[i+1]
		c[i] = c[i] % 10
	}
	if c[lenMax] != 0 {
		for i := lenMax; i >= 0; i-- {
			ret = append(ret, c[i])
		}
	} else {
		for i := lenMax - 1; i >= 0; i-- {
			ret = append(ret, c[i])
		}
	}

	for _, v := range ret {
		result += strconv.Itoa(int(v))
	}

	return result
}

func solution(line string) string {
	// 在此处理单行数据
	// 有三种情况，分别是分隔符为 + < >
	// 返回处理后的结果
	s := strings.Split(line, "+")
	if len(s) != 2 {
		s = strings.Split(line, "<")
		if len(s) != 2 {
			s = strings.Split(line, ">")
			// 判断 num1 是否大于 num2
			return whoBig(s[0], s[1])
		}
		// 判断 num1 是否小于 num2
		return whoBig(s[1], s[0])
	}
	// 大数相加
	return bigAdd(s[0], s[1])
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}