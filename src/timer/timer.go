package timer

type Timer struct {
	TimerOuter      // 执行方法
	id       uint32 // 定时器id
	endTime  uint32 // 执行时间
	interval uint32 // 间隔时间
	index    int    // 在堆数组中的索引
}