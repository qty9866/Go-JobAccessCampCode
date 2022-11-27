package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDefChannel(t *testing.T) {
	var simpleMap = map[string]int{}
	fmt.Println(simpleMap)
	// 声明一个int类型的channel
	var intChan chan int
	// channel智能通过make关键字进行实例化
	fmt.Println(intChan)
	intChan = make(chan int, 3)
	fmt.Println(intChan)

	fmt.Println("装入数据：")
	intChan <- 6
	intChan <- 9
	out := <-intChan
	fmt.Println("读取数据：")
	fmt.Println(out) //6
	out = <-intChan
	fmt.Println(out) //9  channel是先进先出的
}

func TestChanPutGet(t *testing.T) {
	workerCount := 10
	intCHAN := make(chan int) //创建一个不带buffer的channel
	for i := 0; i < workerCount; i++ {
		go func(item int) {
			fmt.Println("放进", item)
			intCHAN <- item
		}(i)
	}

	for j := 0; j < workerCount; j++ {
		go func(jj int) {
			out := <-intCHAN
			fmt.Printf("出%d拿到%d\n", jj, out)
		}(j)
	}
	time.Sleep(time.Second)
}

// 这是一个让out部分等待一段时间再去取数据，来观察i的行为
func TestChanPutGet2_owait(t *testing.T) {
	workerCount := 10
	intCHAN := make(chan int) //创建一个不带buffer的channel
	for i := 0; i < workerCount; i++ {
		go func(item int) {
			fmt.Println("开始工作", time.Now())
			intCHAN <- item
			fmt.Println("结束工作", time.Now())
		}(i)
	}
	fmt.Println(time.Now())
	time.Sleep(5 * time.Second)
	fmt.Println(time.Now())

	for j := 0; j < workerCount; j++ {
		go func(jj int) {
			out := <-intCHAN
			fmt.Printf("%s 出%d拿到%d\n", time.Now(), jj, out)
		}(j)
	}
	time.Sleep(time.Second)
}

/*
运行结果: 如果没有out，in部分会等待直到out开始
=== RUN   TestChanPutGet2_owait
开始工作 2022-11-25 22:40:23.2366101 +0800 CST m=+0.00265    4001
开始工作 2022-11-25 22:40:23.2371289 +0800 CST m=+0.00317    2801
开始工作 2022-11-25 22:40:23.2371289 +0800 CST m=+0.00317    2801
开始工作 2022-11-25 22:40:23.2371289 +0800 CST m=+0.00317    2801
开始工作 2022-11-25 22:40:23.2371289 +0800 CST m=+0.00317    2801
开始工作 2022-11-25 22:40:23.2371289 +0800 CST m=+0.00317    2801
开始工作 2022-11-25 22:40:23.2366101 +0800 CST m=+0.00265    4001
2022-11-25 22:40:23.2366101 +0800 CST m=+0.002654001
开始工作 2022-11-25 22:40:23.2371289 +0800 CST m=+0.00317    2801
开始工作 2022-11-25 22:40:23.2371289 +0800 CST m=+0.00317    2801
开始工作 2022-11-25 22:40:23.2371289 +0800 CST m=+0.00317    2801
2022-11-25 22:40:28.2533188 +0800 CST m=+5.019362701
2022-11-25 22:40:28.2537038 +0800 CST m=+5.019747701 出9  拿到0
结束工作 2022-11-25 22:40:28.2537038 +0800 CST m=+5.01974    7701
2022-11-25 22:40:28.2537038 +0800 CST m=+5.019747701 出5拿   到6
结束工作 2022-11-25 22:40:28.2537038 +0800 CST m=+5.01974    7701
2022-11-25 22:40:28.2537038 +0800 CST m=+5.019747701 出6拿   到1
结束工作 2022-11-25 22:40:28.2537038 +0800 CST m=+5.0197477    01
2022-11-25 22:40:28.2537038 +0800 CST m=+5.019747701 出7  拿到8
结束工作 2022-11-25 22:40:28.2537038 +0800 CST m=+5.01974    7701
2022-11-25 22:40:28.2537038 +0800 CST m=+5.019747701 出4  拿到2
结束工作 2022-11-25 22:40:28.2537038 +0800 CST m=+5.01974    7701
结束工作 2022-11-25 22:40:28.2537038 +0800 CST m=+5.01974    7701
2022-11-25 22:40:28.2537038 +0800 CST m=+5.019747701 出3  拿到3
2022-11-25 22:40:28.2537038 +0800 CST m=+5.019747701 出8拿  到5
结束工作 2022-11-25 22:40:28.2537038 +0800 CST m=+5.019747    701
2022-11-25 22:40:28.2537038 +0800 CST m=+5.019747701 出2拿  到9
结束工作 2022-11-25 22:40:28.2537038 +0800 CST m=+5.019747    701
2022-11-25 22:40:28.2537038 +0800 CST m=+5.019747701 出1拿  到4
2022-11-25 22:40:28.2537038 +0800 CST m=+5.019747701 出0拿  到7
结束工作 2022-11-25 22:40:28.2537038 +0800 CST m=+5.01974    7701
结束工作 2022-11-25 22:40:28.2537038 +0800 CST m=+5.01974    7701
--- PASS: TestChanPutGet2_owait (6.02s)
PASS
进程 已完成，退出代码为 0*/

func TestChanPutGet2_owait_withBuffer(t *testing.T) {
	workerCount := 10
	intCHAN := make(chan int, 10) //创建一个带buffer的channel
	for i := 0; i < workerCount; i++ {
		go func(item int) {
			fmt.Println("开始工作了", time.Now())
			intCHAN <- item
			fmt.Println("结束工作了", time.Now())
		}(i)
	}
	fmt.Println(time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now())

	for j := 0; j < workerCount; j++ {
		go func(jj int) {
			out := <-intCHAN
			fmt.Printf("%s 出%d拿到%d\n", time.Now(), jj, out)
		}(j)
	}
	time.Sleep(time.Second)
}

/*输出:
=== RUN   TestChanPutGet2_owait_withBuffer

2022-11-25 23:46:59.4414654 +0800 CST m=+0.002681601
开始工作了 2022-11-25 23:46:59.4414654 +0800 C     ST m=+0.002681601
结束工作了 2022-11-25 23:46:59.4534061 +0800 C     ST m=+0.014622301
开始工作了 2022-11-25 23:46:59.4414654 +0800 C     ST m=+0.002681601
结束工作了 2022-11-25 23:46:59.4534061 +0800 C     ST m=+0.014622301
开始工作了 2022-11-25 23:46:59.4414654 +0800 C     ST m=+0.002681601
结束工作了 2022-11-25 23:46:59.4534061 +0800 C     ST m=+0.014622301
开始工作了 2022-11-25 23:46:59.4414654 +0800 C     ST m=+0.002681601
结束工作了 2022-11-25 23:46:59.4534061 +0800 C     ST m=+0.014622301
开始工作了 2022-11-25 23:46:59.4414654 +0800 C     ST m=+0.002681601
结束工作了 2022-11-25 23:46:59.4534061 +0800 C     ST m=+0.014622301
开始工作了 2022-11-25 23:46:59.4414654 +0800 C     ST m=+0.002681601
结束工作了 2022-11-25 23:46:59.4534061 +0800 C     ST m=+0.014622301
开始工作了 2022-11-25 23:46:59.4414654 +0800 C     ST m=+0.002681601
结束工作了 2022-11-25 23:46:59.4534061 +0800 C     ST m=+0.014622301
开始工作了 2022-11-25 23:46:59.4414654 +0800 C     ST m=+0.002681601
结束工作了 2022-11-25 23:46:59.4534061 +0800 C     ST m=+0.014622301
开始工作了 2022-11-25 23:46:59.4414654 +0800 C     ST m=+0.002681601
结束工作了 2022-11-25 23:46:59.4534061 +0800 C     ST m=+0.014622301
开始工作了 2022-11-25 23:46:59.4420153 +0800 C     ST m=+0.003231501
结束工作了 2022-11-25 23:46:59.4534061 +0800 C     ST m=+0.014622301
2022-11-25 23:47:00.4631638 +0800 CST m=+1.024380001
2022-11-25 23:47:00.4631638 +0800 CST m=+1.024380001 出9拿到2
2022-11-25 23:47:00.4631638 +0800 CST m=+1.024380001 出0拿到3
2022-11-25 23:47:00.4631638 +0800 CST m=+1.024380001 出5拿到0
2022-11-25 23:47:00.4631638 +0800 CST m=+1.024380001 出6拿到5
2022-11-25 23:47:00.4631638 +0800 CST m=+1.024380001 出7拿到7
2022-11-25 23:47:00.4631638 +0800 CST m=+1.024380001 出8拿到1
2022-11-25 23:47:00.4631638 +0800 CST m=+1.024380001 出2拿到8
2022-11-25 23:47:00.4631638 +0800 CST m=+1.024380001 出4拿到4
2022-11-25 23:47:00.4631638 +0800 CST m=+1.024380001 出3拿到6
2022-11-25 23:47:00.4631638 +0800 CST m=+1.024380001 出1拿到9
--- PASS: TestChanPutGet2_owait_withBuffer (2.02s)
*/

// 遍历channel 结束不关闭channel
func TestRanngeChannelWithoutClose(t *testing.T) {
	intChan := make(chan int, 10)
	intChan <- 1
	intChan <- 2
	intChan <- 3
	intChan <- 4
	intChan <- 5
	intChan <- 6

	for i := range intChan {
		fmt.Println(i) // 可以遍历管道，但是因为没有关闭管道最后会报错fatal error: all goroutines are asleep - deadlock!
	}
}

func TestRanngeChannelWithChannelClose(t *testing.T) {
	intChan := make(chan int, 10)
	intChan <- 1
	intChan <- 2
	intChan <- 3
	intChan <- 4
	intChan <- 5
	intChan <- 6
	close(intChan)
	{
		o1, ok := <-intChan
		fmt.Println("直接取出", o1, ok)
	}

	for i := range intChan {
		fmt.Println("遍历取出", i)
	}
	{
		o1, ok := <-intChan
		fmt.Println("直接取出", o1, ok)
	}
}

// 直接运行select
func TestSelectChannel(t *testing.T) {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		chan1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "Golang"
	}()

	select {
	case o1 := <-chan1:
		fmt.Println("chan1已经ready,go:", o1)
	case o2 := <-chan2:
		fmt.Println("chan1已经ready,go", o2)
	default:
		fmt.Println(time.Now(), "所有的channel都不ready")
	}
}

// case的优先级高于default。只要有case中的channel ready，default不会走的
func TestSelectChannelWithDefaultAndChannelReady(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan string)
	fmt.Println("start:", time.Now())

	ch1 <- 1
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Golang"
	}()

	fmt.Println("select:", time.Now())
	select {
	case o := <-ch1:
		fmt.Println(time.Now(), "ch1 ready, go", o)
	case o := <-ch2:
		fmt.Println(time.Now(), "ch2 ready, go", o)
	default:
		fmt.Println(time.Now(), "所有的channel都不ready，直接走default")
	}
	fmt.Println("DONE")
}

// 关闭的(close)的channel在select的时候总是可用的
func TestSelectChannelWithDefaultAndClosedChannel(t *testing.T) {
	chan1 := make(chan int, 1)
	chan2 := make(chan string)
	fmt.Println("关闭chan1")
	close(chan1)

	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "Golang"
	}()

	fmt.Println("select:", time.Now())
	select {
	case o := <-chan1:
		fmt.Println(time.Now(), "ch1 ready, go", o)
	case o := <-chan2:
		fmt.Println(time.Now(), "ch2 ready, go", o)
	default:
		fmt.Println(time.Now(), "所有的channel都不ready，直接走default")
	}
	fmt.Println("DONE")
}

func TestMultipleSelect2(t *testing.T) {
	channel := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			select {
			case <-channel:
				fmt.Println(i, time.Now())
			}
		}(i)
	}

	fmt.Println("关闭channel", time.Now())
	close(channel)

	time.Sleep(1 * time.Second)
}

// 不能重复关一个channel
func TestDualCLoseChannel(t *testing.T) {
	c := make(chan struct{})
	close(c)
	close(c)
	/*=== RUN   TestDualCLoseChannel
	--- FAIL: TestDualCLoseChannel (0.00s)
	panic: close of closed channel [recovered]
		panic: close of closed channel*/
}

func TestMultipleChannelReadySelect(t *testing.T) {
	ch1, ch2 := make(chan int), make(chan int)
	close(ch1)
	close(ch2)
	ch1Count := 0
	ch2Count := 0
	for i := 0; i < 10000; i++ {
		select {
		case <-ch1:
			ch1Count++
		case <-ch2:
			ch2Count++
		}
	}
	fmt.Println("ch1count:", ch1Count, "ch2count:", ch2Count)
}
