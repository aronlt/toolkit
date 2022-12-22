package tutils

import (
	"fmt"
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

// demo
func GetDaysDiff(date1 time.Time, date2 time.Time) int {
	diff := date1.Sub(date2)

	//func Since(t Time) Duration
	//Since returns the time elapsed since t.
	//It is shorthand for time.Now().Sub(t).

	fmt.Println(diff.Hours())       // number of Hours
	fmt.Println(diff.Nanoseconds()) // number of Nanoseconds
	fmt.Println(diff.Minutes())     // number of Minutes
	fmt.Println(diff.Seconds())     // number of Seconds

	fmt.Println(int(diff.Hours() / 24)) // number of days
	return int(diff.Hours() / 24)
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
