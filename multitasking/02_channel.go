package main

/*
    Channel
    这种方式的优点是通过提供原子的通信原语，避免了竞态情形(race condition)下复杂的锁机制。
    golang 的哲学是通过 channel 进行协程(goroutine)之间的通信来实现数据共享
    1、信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。
        ch <- v    // 将 v 发送至信道 ch。
        v := <-ch  // 从 ch 接收值并赋予 v。
        （“箭头”就是数据流的方向。）
    2、和映射与切片一样，信道在使用前必须创建：
        ch := make(chan int)
    3、默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。
    4、以下示例对切片中的数进行求和，将任务分配给两个 Go 程。一旦两个 Go 程完成了它们的计算，它就能算出最终的结果。

    带缓冲的Channel
    1、Channel 可以是 带缓冲的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的 Channel：
        ch := make(chan int, 100)
    2、仅当 Channel 的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。


    range 和 close
    1、发送者可通过 close 关闭一个信道来表示没有需要发送的值了。
        接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭：
        若没有值可以接收且信道已被关闭，那么在执行完
        v, ok := <-ch
            之后 ok 会被设置为 false。

   2、循环 for i := range c 会不断从信道接收值，直到它被关闭。
    *注意：* 只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发程序报错（panic）。

    *还要注意：* 信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如终止一个 range 循环。

    select 语句
    1、select 语句使一个 Go 程可以等待多个通信操作。
    2、select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。

    select 语句
    1、select 语句使一个 Go 程可以等待多个通信操作。
    2、select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。

    默认选择
    1、当 select 中的其它分支都没有准备好时，default 分支就会执行。
    2、为了在尝试发送或者接收时不发生阻塞，可使用 default 分支：
        select {
            case i := <-c:
                // 使用 i
            default:
                // 从 c 中接收会阻塞时执行
            }

    sync.Mutex
    1、如果我们并不需要通信，保证每次只有一个 Go 程能够访问一个共享的变量，从而避免冲突的概念叫做 *互斥（mutual*exclusion）*
    2、我们通常使用 *互斥锁（Mutex）* 这一数据结构来提供这种机制。
    3、Go 标准库中提供了 sync.Mutex 互斥锁类型及其两个方法：
        Lock
        Unlock
    4、我们可以通过在代码前调用 Lock 方法，在代码后调用 Unlock 方法来保证一段代码的互斥执行。参见 Inc 方法。
    5、你可以用一个 map 来缓存已经获取的 URL，但是要注意 map 本身并不是并发安全的！

*/

// ========================================================================
// channel

// func sum(s []int, c chan int) {
//     sum := 0
//     for _, v := range s {
//         sum += v
//     }
//     c <- sum // 将和送入 c
// }
//
// func main() {
//     s := []int{7, 2, 8, -9, 4, 0}
//
//     c := make(chan int)
//     go sum(s[:len(s)/2], c)
//     go sum(s[len(s)/2:], c)
//     // 从相同信道接收到的内容是无序的，FIFO
//     x, y := <-c, <-c // 从 c 中接收
//
//     fmt.Println(x, y, x+y)
// }

// ========================================================================
// 信道缓冲
// func main() {
//     // 缓冲区溢出会造成死锁
//     ch := make(chan int, 10)
//     ch <- 1
//     ch <- 2
//     ch <- 2
//     ch <- 2
//     ch <- 2
//     ch <- 2
//     ch <- 2
//     ch <- 2
//     ch <- 2
//     ch <- 2
//     fmt.Println(<-ch)
//     fmt.Println(<-ch)
//     fmt.Println(<-ch)
//     fmt.Println(<-ch)
//     fmt.Println(<-ch)
//     fmt.Println(<-ch)
//     fmt.Println(<-ch)
//     fmt.Println(<-ch)
//     fmt.Println(<-ch)
//     fmt.Println(<-ch)
// }

// ========================================================================
// range 和 close

// func fibonacci(n int, c chan int) {
//     x, y := 0, 1
//     for i := 0; i < n; i++ {
//         c <- x
//         x, y = y, x+y
//     }
//     close(c)
// }
//
// func main() {
//     c := make(chan int, 10)
//     go fibonacci(cap(c), c)
//     for i := range c {
//         fmt.Println(i)
//     }
// }

// ========================================================================
// select 语句

// func fibonacci(c, quit, test chan int) {
//     x, y := 0, 1
//     for {
//         select {
//         case c <- x:
//             x, y = y, x+y
//         case <- quit:
//             fmt.Println("quit")
//             return
//         case <- test:
//             fmt.Println("test")
//         }
//     }
// }
//
// func main() {
//     c := make(chan int , 10)
//     quit := make(chan int)
//     test := make(chan int)
//     go func() {
//         test <- 6
//         for i := 0; i < 20; i++ {
//             fmt.Println(<-c)
//         }
//         quit <- 0
//     }()
//     fibonacci(c, quit, test)
// }

// ========================================================================
// 默认选择

// func main() {
//     tick := time.Tick(100 * time.Millisecond)
//     boom := time.After(500 * time.Millisecond)
//     for {
//         select {
//         case <-tick:
//             fmt.Println("tick.")
//         case <-boom:
//             fmt.Println("BOOM!")
//             return
//         default:
//             fmt.Println("    .")
//             time.Sleep(50 * time.Millisecond)
//         }
//     }
// }

// ========================================================================
// sync.Mutex（互斥）
// SafeCounter 的并发使用是安全的。

// type SafeCounter struct {
//     v   map[string]int
//     mux sync.Mutex
// }
//
// // Inc 增加给定 key 的计数器的值。
// func (c *SafeCounter) Inc(key string) {
//     c.mux.Lock()
//     // Lock 之后同一时刻只有一个 goroutine 能访问 c.v
//     c.v[key]++
//     c.mux.Unlock()
// }
//
// // Value 返回给定 key 的计数器的当前值。
// func (c *SafeCounter) Value(key string) int {
//     c.mux.Lock()
//     // Lock 之后同一时刻只有一个 goroutine 能访问 c.v
//     defer c.mux.Unlock()
//     return c.v[key]
// }
//
// func main() {
//     c := SafeCounter{v: make(map[string]int)}
//     for i := 0; i < 1000; i++ {
//         go c.Inc("somekey")
//     }
//
//     time.Sleep(time.Second)
//     fmt.Println(c.Value("somekey"))
// }
