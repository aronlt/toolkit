package demo

type LeaveType = string

const (
	AnnualLeave LeaveType = "AnnualLeave"
	Sick        LeaveType = "Sick"
	BankHoliday LeaveType = "BankHoliday"
	Other       LeaveType = "Other"
)

type Weekday int //自定义一个星期类型，作为枚举类型

// 这里必须要+1，不然会导致有默认值错误
const (
	Sun Weekday = iota + 1
	Mon
	Tues
	Wed
	Thur
	Fri
	Sat
)

func (w Weekday) String() string {
	switch w {
	case Sun:
		return "Sun"
	case Mon:
		return "Mon"
	case Tues:
		return "Tues"
	case Wed:
		return "Wed"
	case Thur:
		return "Thur"
	case Fri:
		return "Fri"
	case Sat:
		return "Sat"
	}
	//不存在的枚举类型就返回"N/A"
	return "N/A"
}
