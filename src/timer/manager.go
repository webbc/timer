package timer

import (
	"container/heap"
	"time"
)

type Manager struct {
	id        uint32        // 自增的timerId
	tq        TimerQueue    // 定时器
	execTimer []interface{} // 将要执行的timer
}

func NewTimerManager() *Manager {
	return &Manager{
		tq: make(TimerQueue, 0, 1024),
	}
}

// 添加定时器
func (this *Manager) AddTimer(timerOuter TimerOuter, endTime uint32, interval uint32) uint32 {

	// id自增
	this.id++

	timer := &Timer{
		id:         this.id,
		TimerOuter: timerOuter,
		endTime:    endTime,
		interval:   interval,
	}

	// 添加到堆中
	heap.Push(&this.tq, timer)

	// 将定时器id返回，调用层保存定时器id，通过定时器id来删除定时器
	return this.id
}

// 删除定时器
func (this *Manager) RemoveTimer(timerId uint32) {
	for _, timer := range this.tq {
		if timer.id == timerId {
			heap.Remove(&this.tq, timer.index)
			return
		}
	}
}

// 执行定时器
func (this *Manager) RunTimer() {

	// 没有需要执行的任务
	if this.tq.Len() <= 0 {
		return
	}

	for this.tq.Len() > 0 {

		// 从堆顶取一个
		tmp := this.tq[0]

		// 时间未到
		if uint32(time.Now().Unix()) < tmp.endTime {
			break
		}

		timer := heap.Pop(&this.tq).(*Timer)
		this.execTimer = append(this.execTimer, timer)

		// 可重复执行
		if timer.interval > 0 {
			timer.endTime += timer.interval
			heap.Push(&this.tq, timer)
		}
	}

	// 执行定时器
	if len(this.execTimer) > 0 {
		for _, timer := range this.execTimer {
			timer.(TimerOuter).TimeOut()
		}
	}

	// 清空
	this.execTimer = this.execTimer[:0]
}
