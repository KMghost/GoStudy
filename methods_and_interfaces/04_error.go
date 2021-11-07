package main

/*
   错误
   1、Go 程序使用 error 值来表示错误状态。
   2、与 fmt.Stringer 类似，error 类型是一个内建接口：
       type error interface {
          Error() string
       }
       （与 fmt.Stringer 类似，fmt 包在打印值时也会满足 error。）
   3、通常函数会返回一个 error 值，调用的它的代码应当判断这个错误是否等于 nil 来进行错误处理。
       i, err := strconv.Atoi("42")
       if err != nil {
           fmt.Printf("couldn't convert number: %v\n", err)
           return
       }
       fmt.Println("Converted integer:", i)

   Reader
   1、io 包指定了 io.Reader 接口，它表示从数据流的末尾进行读取。
   2、Go 标准库包含了该接口的许多实现，包括文件、网络连接、压缩和加密等等。
   3、io.Reader 接口有一个 Read 方法：
       func (T) Read(b []byte) (n int, err error)
   4、Read 用数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 io.EOF 错误。

*/

// ========================================================================
// 格式化错误打印
// type MyError struct {
//     When time.Time
//     What string
// }
//
// func (e *MyError) Error() string {
//     return fmt.Sprintf("at %v, %s",
//         e.When, e.What)
// }
//
// func run() error {
//     return &MyError{
//         time.Now(),
//         "it didn't work",
//     }
// }
//
// func main() {
//     if err := run(); err != nil {
//         fmt.Println(err)
//     }
// }

// ========================================================================
// Reader

// func main() {
//     r := strings.NewReader("Hello, Reader!")
//
//     b := make([]byte, 8)
//     for {
//         n, err := r.Read(b)
//         fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
//         fmt.Printf("b[:n] = %q\n", b[:n])
//         if err == io.EOF {
//             break
//         }
//     }
// }
