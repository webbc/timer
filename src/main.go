package main

import (
	"time"
	"timer"
	"fmt"
)

func main() {

	closeChan := make(chan struct{})

	// 开启调度器
	timerManager := timer.NewTimerManager()
	go scheduler(timerManager)

	time.Sleep(time.Second * 1)

	// 添加定时器
	now := uint32(time.Now().Unix())
	timerManager.AddTimer(&timer.TimerCallback{CallBack: A}, now, 0)
	timerId := timerManager.AddTimer(&timer.TimerCallback{CallBack: B}, now, 5)
	timerManager.AddTimer(&timer.TimerCallback{CallBack: C}, now, 10)

	// 删除定时器
	time.Sleep(time.Second * 20)
	timerManager.RemoveTimer(timerId)

	<-closeChan
}

// 启动调度器
func scheduler(timerManager *timer.Manager) {
	ticker := time.NewTicker(1000 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			timerManager.RunTimer()
		}
	}
}

func A() {
	now := uint32(time.Now().Unix())
	fmt.Printf("%v => aaa\n", now)
}

func B() {
	now := uint32(time.Now().Unix())
	fmt.Printf("%v => bbb\n", now)
}

func C() {
	now := uint32(time.Now().Unix())
	fmt.Printf("%v => ccc\n", now)
}
