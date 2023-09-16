package demo

type LeaveType = string

const (
	AnnualLeave LeaveType = "AnnualLeave"
	Sick        LeaveType = "Sick"
	BankHoliday LeaveType = "BankHoliday"
	Other       LeaveType = "Other"
)

type Weekday int //自定义一个星期类型，作为枚举类型
type WeekdayDesc string

// 这里必须要+1，不然会导致有默认值错误
const (
	Unknown Weekday = iota
	Sun
	Mon
	Tues
	Wed
	Thur
	Fri
	Sat
	SunDesc  WeekdayDesc = "sun"
	MonDesc  WeekdayDesc = "mon"
	TuesDesc WeekdayDesc = "Tues"
	WedDesc  WeekdayDesc = "Wed"
	ThurDesc WeekdayDesc = "Thur"
	FriDesc  WeekdayDesc = "Fri"
	SatDesc  WeekdayDesc = "Sat"
)

func NewWeekday(value string) Weekday {
	desc := WeekdayDesc(value)
	switch desc {
	case SunDesc:
		return Sun
	case MonDesc:
		return Mon
	case TuesDesc:
		return Tues
	case WedDesc:
		return Wed
	case ThurDesc:
		return Thur
	case FriDesc:
		return Fri
	case SatDesc:
		return Sat
	default:
		return Unknown
	}
}

func (w Weekday) String() string {
	var value WeekdayDesc
	switch w {
	case Sun:
		value = SunDesc
	case Mon:
		value = MonDesc
	case Tues:
		value = TuesDesc
	case Wed:
		value = WedDesc
	case Thur:
		value = ThurDesc
	case Fri:
		value = FriDesc
	case Sat:
		value = SatDesc
	default:
		//不存在的枚举类型就返回"N/A"
		return "N/A"
	}
	return string(value)
}
