package main

import "fmt"

/*Range
  1、for 循环的 range 形式可遍历切片或映射。
  2、当使用 for 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。
  3、可以将下标或值赋予 _ 来忽略它。
  4、若你只需要索引，忽略第二个变量即可
*/
func main() {
	a := []int{1, 4, 65, 4, 5, 6, 3, 2}
	for v := range a {
		fmt.Println(v)
	}
}
