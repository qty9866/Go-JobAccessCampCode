### withTimeout()实例
  
部署哈勃望远镜，如果部署时间超过了10秒钟，那么就证明部署结束(失败了)

- 部署主框架 DistributeMainFrame() 9s
- 部署主体 DistributeMainBody() 6s
- 部署覆盖板 DistributeCover() 11s  --> 这个会失败

> ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
> 
> 这里表示给10s时间，10s到了就发送信号停止goroutine 

运行结果
>开始部署望远镜，发送信号 
>
>开始部署Cover 
>
>开始部署MainFrame 
> 
>开始部署MainBody 
> 
>部署MainBody完成,用时： 6.009446s 
>
>部署MainFrame完成 9.0138931s
>
>任务时间(10s)已到，取消所有未完成的任务
>
>任务取消：DistributeCover,用时: 11.0142248s



### WithValue()实例

要出门野营了，就拿了个小钱包,钱包里有钱:

- 到爸爸哪里，换了个拎包，小钱包在拎包里，给包里装了充电宝
- 到妈妈那里换了个书包，拎包装书包里，给包里装了个小夹克
- 到奶奶那里，换了个旅行箱，书包装在旅行箱里，给箱子里装了几个大苹果

到营地了，需要什么，拿到什么。


### WithDeadline()

withDeadline()获得一个带有截止时间的Context。到截止时间时，Context会自动取消，后续生成的也会自动取消，和withTimeout()差不多，主要时间格式不同。

**实例**

定个闹钟，到晚上11点，所有人都停下在做的事情，睡觉：

- 看电视的关电视
- 玩手机的关手机
- 玩游戏的关电脑
- 唱歌的闭嘴