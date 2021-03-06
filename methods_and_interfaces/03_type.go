package main

import "fmt"

/*
   类型断言
   1、类型断言 提供了访问接口值底层具体值的方式。
       t := i.(T)
           该语句断言接口值 i 保存了具体类型 T，并将其底层类型为 T 的值赋予变量 t。
           若 i 并未保存 T 类型的值，该语句就会触发一个恐慌。
           为了 判断 一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。
       t, ok := i.(T)
           若 i 保存了一个 T，那么 t 将会是其底层值，而 ok 为 true。
           否则，ok 将为 false 而 t 将为 T 类型的零值，程序并不会产生恐慌。
           请注意这种语法和读取一个映射时的相同之处。

   类型选择
   1、类型选择 是一种按顺序从几个类型断言中选择分支的结构。
   2、类型选择与一般的 switch 语句相似，不过类型选择中的 case 为类型（而非值），
       它们针对给定接口值所存储的值的类型进行比较。
   3、类型选择中的声明与类型断言 i.(T) 的语法相同，只是具体类型 T 被替换成了关键字 type。
   4、此选择语句判断接口值 i 保存的值类型是 T 还是 S。在 T 或 S 的情况下，
       变量 v 会分别按 T 或 S 类型保存 i 拥有的值。
       在默认（即没有匹配）的情况下，变量 v 与 i 的接口类型和值相同。

   Stringer
   1、fmt 包中定义的 Stringer 是最普遍的接口之一。
       type Stringer interface {
           String() string
       }
   2、Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。

*/

// ========================================================================
// 类型断言 （配合接口值使用）正确返回值，不正确返回 0
// func main() {
//     var i interface{} = "hello"
//
//     s := i.(string)
//     fmt.Println(s)
//
//     s, ok := i.(string)
//     fmt.Println(s, ok)
//
//     f, ok := i.(float64)
//     fmt.Println(f, ok)
//
//     f = i.(float64) // 报错(panic) interface conversion: interface {} is string, not float64
//     fmt.Println(f)
// }

// ========================================================================
// 类型选择
// func do(i interface{}) {
//     switch v := i.(type) {
//     case int:
//         fmt.Printf("Twice %v is %v\n", v, v*2)
//     case string:
//         fmt.Printf("%q is %v bytes long\n", v, len(v))
//     default:
//         fmt.Printf("I don't know about type %T!\n", v)
//     }
// }
//
// func main() {
//     do(21)
//     do("hello")
//     do(true)
// }

// ========================================================================
// Stringer
// type Person struct {
// 	Name string
// 	Age  int
// }
//
// func (p Person) String() string {
// 	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
// }
//
// func main() {
// 	a := Person{"Arthur Dent", 42}
// 	z := Person{"Zaphod Beeblebrox", 9001}
// 	fmt.Println(a, z)
// }

// ========================================================================
// Stringer 练习

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
