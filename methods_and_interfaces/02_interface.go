package main

import "fmt"

/*
   接口
   1、接口类型 是由一组方法签名定义的集合。
   2、接口类型的变量可以保存任何实现了这些方法的值。

   接口与隐式实现
   1、类型通过实现一个接口的所有方法来实现该接口。既然无需专门显式声明，也就没有“implements”关键字。
   2、隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。
   3、因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。

   接口值
   1、接口也是值。它们可以像其它值一样传递。
   2、接口值可以用作函数的参数或返回值。
   3、在内部，接口值可以看做包含值和具体类型的元组：(value, type)
   4、接口值保存了一个具体底层类型的具体值。
   5、接口值调用方法时会执行其底层类型的同名方法。

   底层值为 nil 的接口值
   1、即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。
   2、在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 M 方法）。
   注意: 保存了 nil 具体值的接口其自身并不为 nil。

   nil 接口值
   1、nil 接口值既不保存值也不保存具体类型。
   2、为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 具体 方法的类型。

   空接口
   1、指定了零个方法的接口值被称为 *空接口：*
       interface{}
   2、空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）
   3、空接口被用来处理未知类型的值。例如，fmt.Print 可接受类型为 interface{} 的任意数量的参数。
*/

// ========================================================================
// 接口基础写法
// type Abser interface {
//     Abs() float64
// }
//
// func main() {
//     var a Abser
//     f := MyFloat(-math.Sqrt2)
//     v := Vertex{3, 4}
//
//     // 从结构体调用函数
//     a = f  // a MyFloat 实现了 Abser
//     a = &v // a *Vertex 实现了 Abser
//
//     // 下面一行，v 是一个 Vertex（而不是 *Vertex）
//     // 所以没有实现 Abser。
//     // a = v
//
//     fmt.Println(a.Abs())
// }
//
// type MyFloat float64
//
// func (f MyFloat) Abs() float64 {
//     if f < 0 {
//         return float64(-f)
//     }
//     return float64(f)
// }
//
// type Vertex struct {
//     X, Y float64
// }
//
// func (v *Vertex) Abs() float64 {
//     return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// ========================================================================
// 接口与隐式实现

// type I interface {
//     M()
// }
//
// type T struct {
//     S string
// }
//
// // 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
// func (t T) M() {
//     fmt.Println(t.S)
// }
//
// func main() {
//     // 直接从接口调用函数
//     var i I = T{"hello"}
//     i.M()
// }

// ========================================================================
// 接口值

// type I interface {
//     M()
// }
//
// type T struct {
//     S string
// }
//
// func (t *T) M() {
//     fmt.Println(t.S)
// }
//
// type F float64
//
// func (f F) M() {
//     fmt.Println(f)
// }
//
// func main() {
//     var i I
//
//     i = &T{"Hello"}
//     describe(i)
//     i.M()
//
//     i = F(math.Pi)
//     describe(i)
//     i.M()
// }
//
// func describe(i I) {
//     fmt.Printf("(%v, %T)\n", i, i)
// }

// ========================================================================
// 底层值为 nil 的接口值

// type I interface {
//     M()
// }
//
// type T struct {
//     S string
// }
//
// func (t *T) M() {
//     if t == nil {
//         fmt.Println("<nil>")
//         return
//     }
//     fmt.Println(t.S)
// }
//
// func main() {
//     var i I
//
//     var t *T
//     i = t
//     describe(i) // 只是方法里面的值为 nil 不会出现空指针异常
//     i.M()
//
//     i = &T{"hello"}
//     describe(i)
//     i.M()
// }
//
// func describe(i I) {
//     fmt.Printf("(%v, %T)\n", i, i)
// }

// ========================================================================
// nil 接口值

// type I interface {
//     M()
// }
// type T struct {
//     a string
// }
// func (t T)M()  {
//     fmt.Println("11")
// }
// func main() {
//     var i I             // 运行报错 无效内存地址或零指针取消引用
//     // var i I = T{"s"} // 能运行
//     describe(i)
//     i.M()
// }
//
// func describe(i I) {
//     fmt.Printf("(%v, %T)\n", i, i)
// }

// ========================================================================
// 空接口 （处理未知类型的值）

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
