package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancelFunc()
	successFlag := make(chan bool) // 验证成功，返回true，否则返回false

	go account(ctx)
	go distributeService(ctx)
	go config(ctx)
	go verifyService(ctx, successFlag)

	select {
	case <-ctx.Done():
		fmt.Println("超时，验证失败")
	case v := <-successFlag:
		if v {
			fmt.Println("验证成功")
		} else {
			fmt.Println("验证失败")
		}
	}
}

// 账号处理
func account(ctx context.Context) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	doneChan := make(chan string, 2)
	go accountRegister(ctx, doneChan)
	go accountAuth(ctx, doneChan)
	successCount := 0
	for v := range doneChan {
		successCount++
		fmt.Println("job", v)
		if successCount == 2 {
			close(doneChan)
		}
	}
	fmt.Println("账号处理完成")
}

// 账号处理：用户注册
func accountRegister(ctx context.Context, doneChan chan string) {
	fmt.Println("注册账号")
	// TODO：调用一些接口
	defer fmt.Println("注册账号完成")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context结束,不再注册")
			return
		default:
		}
		doneChan <- "AccountRegister Done"
		fmt.Println("accountRegister成功")
		break
	}
}

// 账号处理：用户授权
func accountAuth(ctx context.Context, doneChan chan string) {
	fmt.Println("授权账号")
	// TODO：调用一些接口
	defer fmt.Println("授权账号完成")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context结束,不再授权")
			return
		default:
		}
		doneChan <- "AccountAuth Done"
		fmt.Println("accountAuth成功")
		break
	}
}

// 部署服务
func distributeService(ctx context.Context) {
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second*7)
	defer cancelFunc()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go distributeLB(ctx, &wg)
	go distributeInstance(ctx, &wg)
	fmt.Println("distributeService Done")
}

// 部署服务：部署负载均衡器
func distributeLB(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context结束，删除已经创建的负载均衡器")
			return
		default:
		}
		fmt.Println("distributeLB完成")
		break
	}
}

// 部署服务：部署实例
func distributeInstance(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context结束，删除已经创建的实例")
			return
		default:
		}
		fmt.Println("distributeInstance完成")
		break
	}
}

// 配置
func config(_ context.Context) {
	fmt.Println("注入新服务账号")
}

// 验证
func verifyService(ctx context.Context, flag chan bool) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go verifyFunc(ctx, &wg)
	wg.Wait()
	fmt.Println("验证服务完成")
	flag <- true
}

func verifyFunc(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("context结束，不需要再验证了")
			return
		default:
		}
		fmt.Println("开始验证...")
		time.Sleep(300 * time.Millisecond) // 用来替换验证部分的环节，比如：服务调用、服务模拟等
		if i <= 1 {
			fmt.Println("服务尚未完成，重试")
			continue
		}
		fmt.Println("服务验证完成")
		break
	}
}
