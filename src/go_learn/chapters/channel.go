package chapters

import "fmt"
import "time"

// 无缓冲channel 定义 make(chan int) ,读/写操作会导致程序阻塞，直到有其他的goroutine执行写/读操作
// 有缓冲channle 定义 make(chan int, 1) ， 1表示缓冲区大小为1，满就无法再写入，除非读走


func Count(ch chan int) {
    fmt.Println("counting")
    // 向channel写入数据通常导致程序阻塞，直到有其他goroutine从这个channel读取数据
    ch <- 1
}

func chan_timeout() {

	// 有无缓冲都可以
	timeout := make(chan bool)

	// 无缓冲
	ch := make(chan int)

	go func() {
		time.Sleep(5e9)
		timeout <- true
	}()

	select {
	case <- ch:
		fmt.Println("receive:", <- ch)
	case <- timeout:
		fmt.Println("timeout")
	}
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
	
	// 创建的是无缓冲chan
	// ch_test := make(chan int, 1)
	// for {
	// 	select { // 两个写都能执行，select随机选择一个执行
	// 	case ch_test <- 0:
	// 	case ch_test <- 1:
	// 	}

	// 	i := <- ch_test
	// 	fmt.Println("Value receive:", i)
	// }
	chan_timeout()
}