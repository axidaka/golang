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

func test1() {

	// 创建chan切片，主routine与每个routine共享1个
	chs := make([]chan int, 10)

	for i := 0; i < 10; i ++ {
	   chs[i] = make(chan int) // 无缓冲
	   go Count(chs[i])
	}

	for _, ch := range(chs) {
	   // 从channel读取数据通常导致程序阻塞，直到有其他goroutine写入channel
	  fmt.Println(<- ch)
	}

}

func test2() {
	// 创建单个chan，主routine与所有routine共享1个
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		go Count(ch)
	}

	var count int = 0
	for {
		count += <- ch
		fmt.Println(count)
		if count >= 10 {
			break
		}
	}
}

func test3() {
	//创建的是带1个缓冲chan， 只能写入1个就满了，
	ch_test := make(chan int, 1)
	for {
		select { // 两个写都能执行，select随机选择一个执行
		case ch_test <- 0:
		case ch_test <- 1:
		default:
			fmt.Println("can not write")
		}

		i := <- ch_test
		fmt.Println("Value receive:", i)
	}
}

func test4() {
	ch := make(chan int, 10)// 十个缓冲，可以满足同事写入10个不阻塞

	for i := 0; i < 10; i ++ {
		go Count(ch)
	}

	for i := range ch {
		fmt.Println("recive:", i)
	}

	/*
	output
	counting
	counting
	counting
	counting
	counting
	recive: 1
	recive: 1
	recive: 1
	recive: 1
	recive: 1
	counting
	recive: 1
	counting
	recive: 1
	counting
	recive: 1
	counting
	recive: 1
	counting
	recive: 1

	进程阻塞不退出
	*/
}

func read(ch chan int) {
	fmt.Println("read....")
	x, ok := <- ch
	if ok {
		fmt.Println("readed succ ", x)
	}else {
		fmt.Println("readed fail ", x)
	}
}
func test5() {

	// test close
	ch := make(chan int, 1)
	ch <- 1
	close(ch)
	go read(ch)
	go read(ch)
	time.Sleep(5e9)

	/*
	output
	read....
	readed succ  1
	read....
	readed fail  0

	*/
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
	test4()
}