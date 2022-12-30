package tutils

import (
	"fmt"
	"time"
)

// UnixToString 时间戳转字符串
func UnixToString(tm int64, format ...string) string {
	t := time.Unix(tm, 0)
	layout := getLayout(format...)
	return t.Format(layout)
}

// UnixToTime 时间戳转Time
func UnixToTime(tm int64) time.Time {
	return time.Unix(tm, 0)
}

// StringToUnix 字符串转换为时间戳
func StringToUnix(t string, format ...string) int64 {
	tm := StringToTime(t, format...)
	return tm.Unix()
}

// StringToTime 字符串转时间
func StringToTime(t string, format ...string) time.Time {
	tm, err := time.ParseInLocation(getLayout(format...), t, time.Local)
	if err != nil {
		return time.Now()
	}
	return tm
}

// TimeToString 时间转字符串
func TimeToString(t time.Time, format ...string) string {
	layout := getLayout(format...)
	return t.Format(layout)
}

// TimeToUnix 时间转时间戳
func TimeToUnix(t time.Time) int64 {
	return t.Unix()
}
func getLayout(format ...string) string {
	layout := "2006-01-02 15:04:05"
	if len(format) > 0 {
		layout = format[0]
	}
	return layout
}

// demo
func GetDaysDiff(date1 time.Time, date2 time.Time) int {
	diff := date1.Sub(date2)
	fmt.Println(diff.Hours())           // number of Hours
	fmt.Println(diff.Nanoseconds())     // number of Nanoseconds
	fmt.Println(diff.Minutes())         // number of Minutes
	fmt.Println(diff.Seconds())         // number of Seconds
	fmt.Println(int(diff.Hours() / 24)) // number of days
	return int(diff.Hours() / 24)
}

// GetDayOffset 获取n天前或者后的时间
func GetDayOffset(tm time.Time, ago int64, format ...string) string {
	t := tm.Add(time.Duration(ago) * time.Hour * 24)
	return TimeToString(t, format...)
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

// GetThisMonthFirstZeros 获取本月初时间戳
func GetThisMonthFirstZeros() int64 {
	year, month, _ := time.Now().Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, time.Local).Unix()
}

// GetWeekByIntDay 获取每周返回为int值,1-7
func GetWeekByIntDay(weekDay string) int32 {
	switch weekDay {
	case "Monday":
		return 1
	case "Tuesday":
		return 2
	case "Wednesday":
		return 3
	case "Thursday":
		return 4
	case "Friday":
		return 5
	case "Saturday":
		return 6
	case "Sunday":
		return 7
	}
	return 0
}

// GetMonthByInt 获取每月返回为int值,1-12
func GetMonthByInt(weekDay string) int32 {
	switch weekDay {
	case "January":
		return 1
	case "February":
		return 2
	case "March":
		return 3
	case "April":
		return 4
	case "May":
		return 5
	case "June":
		return 6
	case "July":
		return 7
	case "August":
		return 8
	case "September":
		return 9
	case "October":
		return 10
	case "November":
		return 11
	case "December":
		return 12
	}
	return 0
}
