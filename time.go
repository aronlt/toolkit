package toolkit

import (
	"time"
)

// GetTimeFromTimestamp 根据时间戳打印对应的日期
func GetTimeFromTimestamp(tm int64, format ...string) string {
	t := time.Unix(tm, 0)
	layout := "2006-01-02 15:04:05"
	if len(format) > 0 {
		layout = format[0]
	}
	return t.Format(layout)
}

// GetDateAgo 获取n天前的时间
func GetDateAgo(ago int64, format ...string) string {
	t := time.Now().Add(-1 * time.Duration(ago) * time.Hour * 24)
	return GetTimeFromTimestamp(t.Unix(), format...)
}

type Printer func(time.Duration)

// FuncCost 获取函数的运行时间
func FuncCost(start time.Time, printer Printer) {
	elapsed := time.Since(start)
	printer(elapsed)
}

// RunPeriod 周期运行函数
func RunPeriod(interval time.Duration, handler func(), closeCh ...chan struct{}) {
	ticker := time.NewTicker(interval * time.Second)
	defer ticker.Stop()

	if len(closeCh) > 0 {
		for {
			select {
			case <-ticker.C:
				handler()
			case <-closeCh[0]:
				return
			}
		}
	} else {
		for {
			select {
			case <-ticker.C:
				handler()
			}
		}
	}

}
