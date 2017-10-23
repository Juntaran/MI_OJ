/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/10/23 20:56
  */

package main

import "fmt"

func main() {

	for n := 1; n < 1001; n++ {
		// 第 1 3 6 10 15 ... 都是 1
		var origin, sum int = 0, 0
		for i := 1; ; i++ {
			sum += i
			if sum > (n - 1) {
				origin = sum - i
				break
			}
		}
		ret := n - origin
		fmt.Println(n, ":", ret)
	}
}
