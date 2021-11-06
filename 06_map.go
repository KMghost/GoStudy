package main

import "fmt"

/* 映射
   1、映射将键映射到值。
   2、映射的零值为 nil 。nil 映射既没有键，也不能添加键。
   3、make 函数会返回给定类型的映射，并将其初始化备用。

   映射的文法
   1、映射的文法与结构体相似，不过必须有键名。
   2、若顶级类型只是一个类型名，你可以在文法的元素中省略它。
   3、顶级类型只是一个类型名，你可以在文法的元素中省略它。

   修改映射：
   1、在映射 m 中插入或修改元素：
       m[key] = elem
   2、获取元素：
       elem = m[key]
   3、删除元素：
       delete(m, key)
   4、通过双赋值检测某个键是否存在：
       elem, ok = m[key]
       若 key 在 m 中，ok 为 true ；否则，ok 为 false。
       若 key 不在映射中，那么 elem 是该映射元素类型的零值。
       同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。
       注 ：若 elem 或 ok 还未声明，你可以使用短变量声明：
           elem, ok := m[key]
*/

// ========================================================================
// 使用 new 创建 map

// type Vertex1 struct {
//     Lat, Long float64
// }
//
// var m map[string]Vertex1
//
// func main() {
//     m = make(map[string]Vertex1)
//     m["Bell Labs"] = Vertex1{
//         40.68433, -74.39967,
//     }
//     fmt.Println(m["Bell Labs"])
// }

// ========================================================================
// map 的文法
// type vertex struct {
//     lat, long float64
// }
//
// var m = map[string]vertex{
//     "Bell Labs": vertex{
//         1,2,
//     },
// }
// func main() {
//     fmt.Println(m)
// }

// ========================================================================
// 映射的文法（续）
// type vertex struct {
//     lat, long float64
// }
// type vertex1 struct {
//     tt, tl int8
// }
//
// var m = map[string]vertex{
//     "Bell Labs":{1,2},
// }
// func main() {
//     fmt.Println(m)
// }

// ========================================================================
// 修改映射

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
