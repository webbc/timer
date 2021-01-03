package timer

type TimerCallback struct {
	CallBack func()
}

func (this *TimerCallback) TimeOut() {
	if this.CallBack != nil {
		this.CallBack()
	}
}
