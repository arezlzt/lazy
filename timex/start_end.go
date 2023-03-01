package time

import "time"

// 数字常量
const (
	YearsPerMillennium         = 1000    // 每千年1000年
	YearsPerCentury            = 100     // 每世纪100年
	YearsPerDecade             = 10      // 每十年10年
	QuartersPerYear            = 4       // 每年4季度
	MonthsPerYear              = 12      // 每年12月
	MonthsPerQuarter           = 3       // 每季度3月
	WeeksPerNormalYear         = 52      // 每常规年52周
	weeksPerLongYear           = 53      // 每长年53周
	WeeksPerMonth              = 4       // 每月4周
	DaysPerLeapYear            = 366     // 每闰年366天
	DaysPerNormalYear          = 365     // 每常规年365天
	DaysPerWeek                = 7       // 每周7天
	HoursPerWeek               = 168     // 每周168小时
	HoursPerDay                = 24      // 每天24小时
	MinutesPerDay              = 1440    // 每天1440分钟
	MinutesPerHour             = 60      // 每小时60分钟
	SecondsPerWeek             = 604800  // 每周604800秒
	SecondsPerDay              = 86400   // 每天86400秒
	SecondsPerHour             = 3600    // 每小时3600秒
	SecondsPerMinute           = 60      // 每分钟60秒
	MillisecondsPerSecond      = 1000    // 每秒1000毫秒
	MicrosecondsPerMillisecond = 1000    // 每毫秒1000微秒
	MicrosecondsPerSecond      = 1000000 // 每秒1000000微秒
)

// StartOfCentury 本世纪开始时间
func (t Time) StartOfCentury() Time {
	t.Time = time.Date(t.Time.Year()/YearsPerCentury*YearsPerCentury, 1, 1, 0, 0, 0, 0, t.Time.Location())
	return t
}

// EndOfCentury 本世纪结束时间
func (t Time) EndOfCentury() Time {
	t.Time = time.Date(t.Time.Year()/YearsPerCentury*YearsPerCentury+99, 12, 31, 23, 59, 59, 999999999, t.Time.Location())
	return t
}

// StartOfDecade 本年代开始时间
func (t Time) StartOfDecade() Time {
	t.Time = time.Date(t.Time.Year()/YearsPerDecade*YearsPerDecade, 1, 1, 0, 0, 0, 0, t.Time.Location())
	return t
}

// EndOfDecade 本年代结束时间
func (t Time) EndOfDecade() Time {
	t.Time = time.Date(t.Time.Year()/YearsPerDecade*YearsPerDecade+9, 12, 31, 23, 59, 59, 999999999, t.Time.Location())
	return t
}

// StartOfYear 本年开始时间
func (t Time) StartOfYear() Time {
	t.Time = time.Date(t.Time.Year(), 1, 1, 0, 0, 0, 0, t.Time.Location())
	return t
}

// EndOfYear 本年结束时间
func (t Time) EndOfYear() Time {
	t.Time = time.Date(t.Time.Year(), 12, 31, 23, 59, 59, 999999999, t.Time.Location())
	return t
}

// Quarter 获取当前季度
func (t Time) Quarter() (quarter int) {
	switch {
	case t.Time.Month() >= 10:
		quarter = 4
	case t.Time.Month() >= 7:
		quarter = 3
	case t.Time.Month() >= 4:
		quarter = 2
	case t.Time.Month() >= 1:
		quarter = 1
	}
	return
}

// StartOfQuarter 本季度开始时间
func (t Time) StartOfQuarter() Time {
	t.Time = time.Date(t.Time.Year(), time.Month(3*t.Quarter()-2), 1, 0, 0, 0, 0, t.Time.Location())
	return t
}

// EndOfQuarter 本季度结束时间
func (t Time) EndOfQuarter() Time {
	quarter, day := t.Quarter(), 30
	switch quarter {
	case 1, 4:
		day = 31
	case 2, 3:
		day = 30
	}
	t.Time = time.Date(t.Time.Year(), time.Month(3*quarter), day, 23, 59, 59, 999999999, t.Time.Location())
	return t
}

// StartOfMonth 本月开始时间
func (t Time) StartOfMonth() Time {
	t.Time = time.Date(t.Time.Year(), time.Month(t.Month()), 1, 0, 0, 0, 0, t.Time.Location())
	return t
}

// EndOfMonth 本月结束时间
func (t Time) EndOfMonth() Time {
	t.Time = time.Date(t.Time.Year(), time.Month(t.Month()), t.CurrentMothDays(), 23, 59, 59, 999999999, t.Time.Location())
	return t
}

// StartOfDay 本日开始时间
func (t Time) StartOfDay() Time {
	t.Time = time.Date(t.Time.Year(), t.Time.Month(), t.Time.Day(), 0, 0, 0, 0, t.Time.Location())
	return t
}

// EndOfDay 本日结束时间
func (t Time) EndOfDay() Time {
	t.Time = time.Date(t.Time.Year(), t.Time.Month(), t.Time.Day(), 23, 59, 59, 0, t.Time.Location())
	return t
}

// CurrentMothDays 获取当前月份总天数
func (t Time) CurrentMothDays() int {
	switch t.Time.Month() {
	case 4, 6, 9, 11:
		return 30
	case 2:
		year := t.Time.Year()
		if (year%4 == 0 && year%100 != 0) || year%100 == 0 {
			return 29
		}
		return 28
	default:
		return 31
	}
}
