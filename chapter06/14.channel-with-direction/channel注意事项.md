## channel注意事项
- 没有缓冲区的channel在没有可用的接收者时，程序等待
- Channel不能重复关闭
- 关闭后的channel不能再向里面装数据，但是可以取数据
- 非多线程的程序中，谨慎使用channel
- Select在选择channel时，如果多个channel都准备好了，他会随机选择一个，而不是从上到下