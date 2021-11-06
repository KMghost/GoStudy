package main

import "fmt"

/*
   函数值
   1、函数也是值。它们可以像其它值一样传递。
   2、函数值可以用作函数的参数或返回值。

   函数的闭包
   1、Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。
           该函数可以访问并赋予其引用的变量的值，
           换句话说，该函数被这些变量“绑定”在一起。
   2、Go语言的并发时，一定要处理好循环中的闭包引用的外部变量。如下代码：
*/

// ========================================================================
// 函数值
// func compute(fn func(float64, float64) float64) float64 {
//     return fn(3, 4)
// }
//
// func main() {
//     hypot := func(x, y float64) float64 {
//         return x+y
//     }
//     fmt.Println(hypot(5, 12))
//
//     fmt.Println(compute(hypot))
//     fmt.Println(compute(math.Pow))
// }

// ========================================================================
// 函数的闭包

// func adder() func(int) int {
//     sum := 0
//     fmt.Println(sum)
//     return func(x int) int {
//         sum += x
//         return sum
//     }
// }
//
// func main() {
//     // pos, neg := adder(), adder()
//     pos := adder()   // pos 不仅仅是存储了一个函数的返回值，它同时存储了一个闭包的状态。
//     for i := 0; i < 10; i++ {
//         fmt.Println(
//             pos(i),
//             // neg(-2*i),
//         )
//     }
// }

// ========================================================================
// 在并发中使用闭包

/*
   这种现象的原因在于闭包共享外部的变量i，
       注意到，每次调用go就会启动一个goroutine，这需要一定时间；
       但是，启动的goroutine与循环变量递增不是在同一个goroutine，
       可以把 i 认为处于主goroutine中。启动一个goroutine的速度远小于循环执行的速度，
       所以即使是第一个goroutine刚起启动时，外层的循环也执行到了最后一步了。
       由于所有的goroutine共享i，而且这个i会在最后一个使用它的goroutine结束后被销毁，
       所以最后的输出结果都是最后一步的i==5。

*/
// func main() {
//     runtime.GOMAXPROCS(runtime.NumCPU())
//
//     var wg sync.WaitGroup
//     for i := 0; i < 5; i++ {
//         wg.Add(1)
//         go func() {
//             fmt.Println(i)
//             wg.Done()
//         }()
//         // time.Sleep(1 * time.Second)   // 验证上述说法
//     }
//     wg.Wait()
// }

// ========================================================================
// 解决高并发使用 闭包 的思路 （方法一）共享的环境变量作为函数参数传递
// func main() {
//     runtime.GOMAXPROCS(runtime.NumCPU())
//
//     var wg sync.WaitGroup
//     for i := 0; i < 5; i++ {
//         wg.Add(1)
//         go func(i int) {
//             fmt.Println(i)
//             wg.Done()
//         }(i)
//     }
//     wg.Wait()
// }

// ========================================================================
// 方法二 使用同名的变量保留当前的状态
// func main() {
//     runtime.GOMAXPROCS(runtime.NumCPU())
//
//     var wg sync.WaitGroup
//     for i := 0; i < 5; i++ {
//         wg.Add(1)
//         i := i       // 注意这里的同名变量覆盖
//         go func() {
//             fmt.Println(i)
//             wg.Done()
//         }()
//     }
//     wg.Wait()
// }

// ========================================================================
// 实现斐波那契

func fibonacci() func() int {
	i, j := 1, 1
	return func() int {
		i, j = j, j+i
		return i
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
