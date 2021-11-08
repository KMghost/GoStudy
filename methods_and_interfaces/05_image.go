package main

// import "golang.org/x/tour"
/*
   图像
   1、image 包定义了 Image 接口：
       package image

       type Image interface {
           ColorModel() color.Model
           Bounds() Rectangle
           At(x, y int) color.Color
       }
   2、注意: Bounds 方法的返回值 Rectangle 实际上是一个 image.Rectangle，它在 image 包中声明。
*/

// ========================================================================
// image

// func main() {
//     m := image.NewRGBA(image.Rect(0, 0, 100, 100))
//     fmt.Println(m.Bounds())
//     fmt.Println(m.At(0, 0).RGBA())
// }

// type Image struct{}
//
// func main() {
//     m := Image{}
//     pic.ShowImage(m)
// }
