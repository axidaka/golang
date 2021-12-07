package main

//导入fmt包
import (
	"fmt"
	"golang/src/go_learn/chapters"
)

type Rect struct {
	x, y          float64
	width, height float64
}

func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

//main函数定义
func main() {
	fmt.Println("----------------------")
	chapters.Channel_test()
}
