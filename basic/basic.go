package main

// todo 一、导入导出
// 1、分组导入

// import (
//    "fmt"
//    "math"
// )

// todo 2、单个导入  （ 导入不使用的包会报错 ）

// import "fmt"

// todo 注意：小写字母开头为私有变量，大写字母开头为公共变量
// ========================================================================

// todo 二，函数
// 1、函数可以没有参数或者有多个参数（类型在变量名的后面）

// import "fmt"
//
// func add(x int, y int) int {
//    return x + y
// }
//
// func main() {
//    fmt.Println(add(42, 13))
// }

// todo 2、当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。
// x int, y int
// 被缩写为
//
// x, y int

// todo 3、函数可以返回任意数量的返回值
// import "fmt"
//
// func swap(x, y string) (string, string) {
//    return y, x
// }
//
// func main() {
//    a, b := swap("hello", "world")
//    fmt.Println(a, b)
// }

/*todo 4、Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
  返回值的名称应当具有一定的意义，它可以作为文档使用。
  没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
  直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。
*/
// func split(sum int) (x, y int) {
//    x = sum * 4 / 9
//    y = sum - x
//    return
// }
//
// func main() {
//    fmt.Println(split(17))
// }

// ========================================================================

// todo 三、变量
// todo 1、var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后。
// var c, python, java bool
// var i int

/*todo 2、变量初始化
  变量声明可以包含初始值
  如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。
*/
// var c, python, java = true, false, "no!"

// todo 3、短变量声明
/*  := 可在类型明确的地方代替 var 声明。
    函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。
*/
// k := 3
// c, python, java := true, false, "no!"

// ========================================================================

// todo 四、基本类型
/*todo
  bool
  string
  int  int8  int16  int32  int64
  uint uint8 uint16 uint32 uint64 uintptr
  byte // uint8 的别名
  rune // int32 的别名
      // 表示一个 Unicode 码点
  float32 float64
  complex64 complex128
  int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。
  当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。
*/

//
// var (
//	MaxInt uint64     = 1<<64 - 1
//    z      complex128 = cmplx.Sqrt(-5 + 12i)
// )
//
// func main() {
//    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
//    fmt.Printf("Type: %T Value: %v\n", z, z)
// }

/*todo  2、没有明确初始值的变量声明会被赋予它们的 零值。
  数值类型为 0，
  布尔类型为 false，
  字符串为 ""（空字符串）。
*/

// todo  3、类型转换

// var i int = 42
// var f float64 = float64(i)
// var u uint = uint(f)
// 或者
// i := 42
// f := float64(i)
// u := uint(f)

// todo 4、类型推导
// todo 在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），变量的类型由右值推导得出。
// todo 不过当右边包含未指明类型的数值常量时，新变量的类型就可能是 int, float64 或 complex128 了，这取决于常量的精度：

// i := 42           // int
// f := 3.142        // float64
// g := 0.867 + 0.5i // complex128

/*todo 5、常量
  常量的声明与变量类似，只不过是使用 const 关键字。
  常量可以是字符、字符串、布尔值或数值。
  常量不能用 := 语法声明。
*/

/*todo 6、数值常量
  数值常量是高精度的 值。
  一个未指定类型的常量由上下文来决定其类型。
  再尝试一下输出 needInt(Big) 吧。
  （int 类型最大可以存储一个 64 位的整数，有时会更小。）
  （int 可以存放最大64位的整数，根据平台不同有时会更少。）
*/
