/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/10/23 21:16
  */

package main

import (
	"bufio"
	"os"
	"fmt"
)

type stack struct {
	chars	[]uint8		// 栈
	size	int			// 栈大小
}

// 创建一个栈
func newStack() *stack {
	newStack := new(stack)
	newStack.size = 0
	newStack.chars = make([]uint8, 0)
	return newStack
}

// 入栈
func (s *stack) push(e uint8) {
	s.chars = append(s.chars, e)
	s.size ++
}

// 出栈
func (s *stack) pop() {
	s.chars = s.chars[:s.size-1]
	s.size --
}

// 判断栈是否为空
func (s *stack) isEmpty() bool {
	if s.size == 0 {
		return true
	}
	return false
}

func solution(line string) string {
	// 在此处理单行数据
	s := newStack()
	//fmt.Println(reflect.TypeOf(line[0]))
	for i := 0; i < len(line); i++ {
		v := line[i]

		// 左括号入栈
		if v == '(' || v == '[' || v == '{' {
			s.push(v)
		}
		// 右括号出栈
		if v == ')' || v == ']' || v == '}' {
			if s.isEmpty() == false {
				if v == ')' {
					if s.chars[s.size - 1] == '(' {
						s.pop()
					} else {
						return "0"
					}
				}
				if v == ']' {
					if s.chars[s.size - 1] == '[' {
						s.pop()
					} else {
						return "0"
					}
				}
				if v == '}' {
					if s.chars[s.size - 1] == '{' {
						s.pop()
					} else {
						return "0"
					}
				}
			} else {
				return "0"
			}
		}
	}
	if s.isEmpty() == false {
		return "0"
	}

	// 返回处理后的结果
	return "1"
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}