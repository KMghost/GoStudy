package main

import (
	"fmt"
	"math"
)

/*
   Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。
   同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。
   在 if 的简短语句中声明的变量同样可以在任何对应的 else 块中使用。

*/

// func pow(x, n, lim float64) float64 {
//     if v := math.Pow(x, n); v < lim {
//         return v
//     }
//     return lim
// }
//
// func main() {
//     fmt.Println(
//       ow(3, 2, 10),
//     )
// }

// ========================================================================
// 使用牛顿公式计算平方根

func unum(z, new_z float64) float64 {
	/*求出绝对值*/
	var condition float64
	if z-new_z < 0 {
		condition = (z - new_z) * (-1.0)
	} else {
		condition = (z - new_z)
	}
	return condition
}

func Sqrt(x float64) float64 {
	// 新的结果减去旧的结果偏差小于 0.001 则输出结果
	z := 1.0
	new_z := z - (z*z-x)/(2*z)
	var condition float64
	condition = unum(z, new_z)
	for condition > 0.001 {
		z = new_z
		new_z = new_z - (new_z*new_z-x)/(2*new_z)
		condition = unum(z, new_z)
	}
	return new_z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))

}
