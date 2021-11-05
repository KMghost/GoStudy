package main

import (
	"fmt"
	"time"
)

/*
   switch 是编写一连串 if - else 语句的简便方法。它运行第一个值等于条件表达式的 case 语句。
   Go 只运行选定的 case，而非之后所有的 case。实际上，Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。
   除非以 fallthrough 语句结束，否则分支会自动终止。Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。

   switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。

   switch i {
       case 0:
       case f():
       }
   在 i==0 时 f 不会被调用。

   没有条件的 switch 同 switch true 一样。


*/

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
