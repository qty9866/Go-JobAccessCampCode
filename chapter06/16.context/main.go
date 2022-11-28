package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//WithCancel()
	//withTimeout()
	//withValue()
	withDeadline()
	time.Sleep(10 * time.Second)
}

func withDeadline() {
	now := time.Now()
	deadline := now.Add(5 * time.Second)
	//func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
	ctx, _ := context.WithDeadline(context.TODO(), deadline)
	go WatchTv(ctx)
	go WatchPhone(ctx)
	go PlayGames(ctx)
	go singing(ctx)
}
func WatchTv(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关电视了，睡觉")
			return
		default:
			fmt.Println("依旧在看电视...")
			time.Sleep(1 * time.Second)
		}
	}
}
func WatchPhone(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关手机了，睡觉")
			return
		default:
			fmt.Println("依旧在玩手机...")
			time.Sleep(1 * time.Second)
		}
	}
}
func PlayGames(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关电脑了，睡觉")
			return
		default:
			fmt.Println("依旧在打游戏...")
			time.Sleep(1 * time.Second)
		}
	}
}
func singing(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("闭嘴了，睡觉")
			return
		default:
			fmt.Println("依旧在唱歌...")
			time.Sleep(1 * time.Second)
		}
	}
}

func withValue() {
	ctx := context.WithValue(context.TODO(), "1", "wallet") // WithValue(parent Context, key any, val any) Context
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("-------------自己拿了钱包--------------")
		fmt.Println("1", ctx.Value("1"))
		fmt.Println("2", ctx.Value("2"))
		fmt.Println("3", ctx.Value("3"))
		fmt.Println("4", ctx.Value("4"))
	}(ctx)
	gotoPapa(ctx)
}

func gotoPapa(ctx context.Context) {
	ctx = context.WithValue(ctx, "2", "charging treasure")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("-------------去了爸爸那里--------------")
		fmt.Println("1", ctx.Value("1"))
		fmt.Println("2", ctx.Value("2"))
		fmt.Println("3", ctx.Value("3"))
		fmt.Println("4", ctx.Value("4"))
	}(ctx)
	gotoMama(ctx)
}
func gotoMama(ctx context.Context) {
	ctx = context.WithValue(ctx, "3", "jacket")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("-------------去了妈妈那里--------------")
		fmt.Println("1", ctx.Value("1"))
		fmt.Println("2", ctx.Value("2"))
		fmt.Println("3", ctx.Value("3"))
		fmt.Println("4", ctx.Value("4"))
	}(ctx)
	gotoGrandma(ctx)
}
func gotoGrandma(ctx context.Context) {
	ctx = context.WithValue(ctx, "4", "Apples")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("-------------去了奶奶那里--------------")
		fmt.Println("1", ctx.Value("1"))
		fmt.Println("2", ctx.Value("2"))
		fmt.Println("3", ctx.Value("3"))
		fmt.Println("4", ctx.Value("4"))
	}(ctx)
	gotoCamping(ctx)
}
func gotoCamping(ctx context.Context) {
	time.Sleep(4 * time.Second)
	go func(ctx context.Context) {
		fmt.Println("-------------去了营地那里--------------")
		fmt.Println("1", ctx.Value("1"))
		fmt.Println("2", ctx.Value("2"))
		fmt.Println("3", ctx.Value("3"))
		fmt.Println("4", ctx.Value("4"))
	}(ctx)
}

// withTimeout 部署望远镜
func withTimeout() {
	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second) // 只给两秒钟的时间去完成这项任务
	fmt.Println("开始部署望远镜，发送信号")
	go DistributeMainFrame(ctx)
	go DistributeMainBody(ctx)
	go DistributeCover(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("任务时间(10s)已到，取消所有未完成的任务")
	}
	time.Sleep(15 * time.Second)
}

func DistributeMainFrame(ctx context.Context) {
	start := time.Now()
	fmt.Println("开始部署MainFrame")
	time.Sleep(9 * time.Second)
	select {
	case <-ctx.Done():
		finish := time.Now()
		fmt.Println("任务取消：DistributeMainFrame，用时：", finish.Sub(start))
		return
	default:
	}
	finish := time.Now()
	fmt.Println("部署MainFrame完成", finish.Sub(start))
}
func DistributeMainBody(ctx context.Context) {
	start := time.Now()
	fmt.Println("开始部署MainBody")
	time.Sleep(6 * time.Second)
	select {
	case <-ctx.Done():
		finish := time.Now()
		fmt.Println("任务取消：DistributeMainBody", finish.Sub(start))
		return
	default:
	}
	finish := time.Now()
	fmt.Println("部署MainBody完成,用时：", finish.Sub(start))
}
func DistributeCover(ctx context.Context) {
	start := time.Now()
	fmt.Println("开始部署Cover")
	time.Sleep(11 * time.Second)
	select {
	case <-ctx.Done():
		finish := time.Now()
		fmt.Println("任务取消：DistributeCover,用时:", finish.Sub(start))
		return
	default:
	}
	finish := time.Now()
	fmt.Println("部署Cover完成，用时:", finish.Sub(start))
}

func WithCancel() {
	ctx := context.TODO()
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("做蛋挞需要买材料：")
	go BuyFlour(ctx)
	go BuyOil(ctx)
	go BuyEgg(ctx)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("没电了，取消购买所有")
	cancel()
}

func BuyFlour(ctx context.Context) {
	fmt.Println("去买面")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买面了")
		return
	default:
	}
	fmt.Println("买面")
}

func BuyOil(ctx context.Context) {
	fmt.Println("去买油")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买油了")
		return
	default:
	}
	fmt.Println("买油")
}

func BuyEgg(ctx1 context.Context) {
	ctx, cancelFunc := context.WithCancel(ctx1)
	defer cancelFunc()
	fmt.Println("去买蛋")
	//time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买蛋了")
		return
	default:
	}
	fmt.Println("买蛋")
	go BuyBEgg(ctx)
	go BuySEgg(ctx)
	time.Sleep(1 * time.Second)
}

func BuyBEgg(ctx context.Context) {
	fmt.Println("去买大鸡蛋")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买大鸡蛋了")
		return
	default:
	}
	fmt.Println("买大鸡蛋")
}

func BuySEgg(ctx context.Context) {
	fmt.Println("去买小鸡蛋")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买小鸡蛋了")
		return
	default:
	}
	fmt.Println("买小鸡蛋")
}
