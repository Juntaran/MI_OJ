/** 
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/8/10 21:38
  */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
	"sort"
)

func solution(line string) string {
	// 在此处理单行数据
	s := strings.Split(line, ";")

	a := s[0]
	b := s[1]
	c := s[2]
	d := s[3]

	coordinate1 := strings.Split(a, ",")
	coordinate2 := strings.Split(b, ",")
	coordinate3 := strings.Split(c, ",")
	coordinate4 := strings.Split(d, ",")

	a1, _ := strconv.ParseFloat(coordinate1[0], 64)
	a2, _ := strconv.ParseFloat(coordinate1[1], 64)

	b1, _ := strconv.ParseFloat(coordinate2[0], 64)
	b2, _ := strconv.ParseFloat(coordinate2[1], 64)

	c1, _ := strconv.ParseFloat(coordinate3[0], 64)
	c2, _ := strconv.ParseFloat(coordinate3[1], 64)

	d1, _ := strconv.ParseFloat(coordinate4[0], 64)
	d2, _ := strconv.ParseFloat(coordinate4[1], 64)


	AtoB := math.Sqrt(math.Pow(a1-b1, 2) + math.Pow(a2-b2, 2))
	AtoC := math.Sqrt(math.Pow(a1-c1, 2) + math.Pow(a2-c2, 2))
	AtoD := math.Sqrt(math.Pow(a1-d1, 2) + math.Pow(a2-d2, 2))
	BtoC := math.Sqrt(math.Pow(b1-c1, 2) + math.Pow(b2-c2, 2))
	BtoD := math.Sqrt(math.Pow(b1-d1, 2) + math.Pow(b2-d2, 2))
	CtoD := math.Sqrt(math.Pow(c1-d1, 2) + math.Pow(c2-d2, 2))

	lengthSlice := []float64{AtoB, AtoC, AtoD, BtoC, BtoD, CtoD}
	sort.Float64s(lengthSlice)


	// 返回处理后的结果
	if lengthSlice[0] == lengthSlice[3] {
		return "1"
	} else {
		return "0"
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}