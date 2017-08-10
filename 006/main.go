/**
    * Author: Juntaran
    * Email:  Jacinthmail@gmail.com
    * Date:   2017/8/10 20:19
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
	s := strings.Split(line, ",")
	if len(s[0]) + len(s[1]) != len(s[2]) {
		return "false"
	}
	if len(s[0]) == 0 {
		if s[1] == s[2] {
			return "true"
		} else {
			return "false"
		}
	}
	if len(s[1]) == 0 {
		if s[0] == s[2] {
			return "true"
		} else {
			return "false"
		}
	}

	dp := make([][]string, len(s[1]) + 1)
	for i := range dp {
		dp[i] = make([]string, len(s[0]) + 1)
	}

	for i := 0; i <= len(s[1]); i++ {
		for j := 0; j <= len(s[0]); j++ {
			dp[i][j] = "false"
		}
	}

	dp[0][0] = "true"

	for i := 1; i < len(s[0]) + 1; i++ {
		if s[0][i-1] == s[2][i-1] {
			dp[i][0] = "true"
		}
	}

	for i := 1; i < len(s[1]) + 1; i++ {
		if s[1][i-1] == s[2][i-1] {
			dp[0][i] = "true"
		}
	}

	for i := 1; i < len(s[0]) + 1; i++ {
		for j := 1; j < len(s[1]) + 1; j++ {
			k := i + j
			if s[0][i-1] == s[2][k-1] {
				if dp[i][j] == "true" || dp[i-1][j] == "true" {
					dp[i][j] = "true"
				}
			}
			if s[1][j-1] == s[2][k-1] {
				if dp[i][j] == "true" || dp[i][j-1] == "true" {
					dp[i][j] = "true"
				}
			}
		}
	}

	// 返回处理后的结果
	return dp[len(s[0])][len(s[1])]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}
