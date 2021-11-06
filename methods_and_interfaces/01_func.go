package main

import (
	"fmt"
	"math"
)

/*  方法
    1、Go 没有类。不过你可以为结构体类型定义方法。
    2、方法就是一类带特殊的 接收者 参数的函数。
    3、方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
    4、记住：方法只是个带接收者参数的函数。

    为非结构体定义方法
    1、你也可以为非结构体类型声明方法。
    2、你只能为在同一包内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 int 之类的内建类型）
        的接收者声明方法。

    指针接收者
    1、你可以为指针接收者声明方法。
    2、这意味着对于某类型 T，接收者的类型可以用 *T 的文法。（此外，T 不能是像 *int 这样的指针。）
    3、指针接收者的方法可以修改接收者指向的值（就像 Scale 在这做的）。
        由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。

    方法与指针重定向（一）
    1、带指针参数的函数必须接受一个指针
        var v Vertex
        ScaleFunc(v, 5)  // 编译错误！
        ScaleFunc(&v, 5) // OK
    2、而以指针为接收者的方法被调用时，接收者既能为值又能为指针：
        var v Vertex
        v.Scale(5)  // OK
        p := &v
        p.Scale(10) // OK
            对于语句 v.Scale(5)，即便 v 是个值而非指针，带指针接收者的方法也能被直接调用。
            也就是说，由于 Scale 方法有一个指针接收者，为方便起见，
            Go 会将语句 v.Scale(5) 解释为 (&v).Scale(5)。

    3、方法与指针重定向（二）
        1、接受一个值作为参数的函数必须接受一个指定类型的值：
            var v Vertex
            fmt.Println(AbsFunc(v))  // OK
            fmt.Println(AbsFunc(&v)) // 编译错误！
        2、而以值为接收者的方法被调用时，接收者既能为值又能为指针：
            var v Vertex
            fmt.Println(v.Abs()) // OK
            p := &v
            fmt.Println(p.Abs()) // OK
                这种情况下，方法调用 p.Abs() 会被解释为 (*p).Abs()。

    4、选择值或者指针作为接收者
        使用指针接收者的原因有二：
            1、首先，方法能够修改其接收者指向的值。
            2、其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。
            3、通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。

*/

// ========================================================================
// 在此例中，Abs 方法拥有一个名为 v，类型为 Vertex 的接收者。

// type Vertex struct {
//     X, Y float64
// }
//
// func (v Vertex) Abs() float64 {
//     return v.X+v.Y
// }
//
// func main() {
//     v := Vertex{3, 4}
//     fmt.Println(v.Abs())
//     v = Vertex{3, 6}
//     fmt.Println(v.Abs())
// }

// ========================================================================
// 为非结构体定义方法

// type MyFloat float64
//
// func (f MyFloat) Abs() float64 {
//     if f < 0 {
//         return float64(-f)
//     }
//     return float64(f)
// }
//
// func main() {
//     f := MyFloat(-3)
//     fmt.Println(f.Abs())
// }

// ========================================================================
// 指针接收者

// type Vertex struct {
//     X, Y float64
// }
//
// func (v Vertex) Abs() float64 {
//     return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }
//
// func (v *Vertex) Scale(f float64) {
//     v.X = v.X * f
//     v.Y = v.Y * f
// }
//
// func main() {
//     v := Vertex{3, 4}
//     v.Scale(10)
//     fmt.Println(v.Abs())
// }

// ========================================================================
// 指针与函数

// type Vertex struct {
//     X, Y float64
// }
//
// func Abs(v Vertex) float64 {
//     return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }
//
// func Scale(v *Vertex, f float64) {
//     v.X = v.X * f
//     v.Y = v.Y * f
// }
// func (v *Vertex) Scale(f float64) {
//     v.X = v.X * f
//     v.Y = v.Y * f
// }
// func main() {
//     v := Vertex{3, 4}
//     Scale(&v,10)
//     fmt.Println(Abs(v))
// }

// ========================================================================
// 方法与指针重定向

// type Vertex struct {
//     X, Y float64
// }
//
// func (v *Vertex) Scale(f float64) {
//     v.X = v.X * f
//     v.Y = v.Y * f
// }
//
// func ScaleFunc(v *Vertex, f float64) {
//     v.X = v.X * f
//     v.Y = v.Y * f
// }
//
// func main() {
//     v := Vertex{3, 4}
//     v.Scale(2)
//     ScaleFunc(&v, 10)
//
//     p := &Vertex{4, 3}
//     p.Scale(3)
//     ScaleFunc(p, 8)
//
//     fmt.Println(v, p)
// }

// ========================================================================
// 指针重定向2
// type Vertex struct {
//     X, Y float64
// }
//
// func (v Vertex) Abs() float64 {
//     return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }
//
// func AbsFunc(v Vertex) float64 {
//     return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }
//
// func main() {
//     v := Vertex{3, 4}
//     fmt.Println(v.Abs())     // 使用指针
//     fmt.Println(AbsFunc(v))  // 传值的方式
//
//     p := &Vertex{4, 3}
//     fmt.Println(p.Abs())        // 使用指针
//     fmt.Println(AbsFunc(*p))    // 传值的方式
// }

// ========================================================================
// 选择值或者指针作为接收者
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
