package logger

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */

type Logger struct {
	limitStore map[string]int
}

func Constructor() Logger {
	ls := make(map[string]int)
	return Logger{limitStore: ls}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	ts, ok := this.limitStore[message]
	if !ok {
		this.limitStore[message] = timestamp
		return true
	}
	if timestamp-ts >= 10 {
		this.limitStore[message] = timestamp
		return true
	}
	return false
}
