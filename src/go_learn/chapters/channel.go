package chapters

import "fmt"

func Count(ch chan int) {
    fmt.Println("counting")
    // 向channel写入数据通常导致程序阻塞，直到有其他goroutine从这个channel读取数据
    ch <- 1
}

func Channel_test() {

	chs := make([]chan int, 10)

    for i := 0; i < 10; i ++ {
        chs[i] = make(chan int)
        go Count(chs[i])
    }

    for _, ch := range(chs) {
        // 从channel读取数据通常导致程序阻塞，直到有其他goroutine写入channel
       fmt.Println(<- ch)
    }
}