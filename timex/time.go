package time

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 时间格式化常量
const (
	RFC3339Format       = time.RFC3339
	Iso8601Format       = "2006-01-02T15:04:05-07:00"
	CookieFormat        = "Monday, 02-Jan-2006 15:04:05 MST"
	RFC1036Format       = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC7231Format       = "Mon, 02 Jan 2006 15:04:05 GMT"
	DayDateTimeFormat   = "Mon, Jan 2, 2006 3:04 PM"
	DateTimeFormat      = "2006-01-02 15:04:05"
	DateFormat          = "2006-01-02"
	TimeFormat          = "15:04:05"
	ShortDateTimeFormat = "20060102150405"
	ShortDateFormat     = "20060102"
	ShortTimeFormat     = "150405"
)

// SetCurrent 设置当前的时间
func SetCurrent(sTime time.Time) Time {
	p := NewTime()
	p.Time = sTime
	return p
}

// SetCurrentParse 设置当前的时间
func SetCurrentParse(str string) Time {
	t := NewTime()
	t.loc, t.Error = time.LoadLocation("Asia/Shanghai")
	layout := DateTimeFormat
	if str == "" || str == "0" || str == "0000-00-00 00:00:00" || str == "0000-00-00" || str == "00:00:00" {
		return t
	}
	if len(str) == 10 && strings.Count(str, "-") == 2 {
		layout = DateFormat
	}
	if strings.Index(str, "T") == 10 {
		layout = RFC3339Format
	}
	if _, err := strconv.ParseInt(str, 10, 64); err == nil {
		switch len(str) {
		case 8:
			layout = ShortDateFormat
		case 14:
			layout = ShortDateTimeFormat
		}
	}
	location, _ := time.ParseInLocation(layout, str, t.loc)
	t.Time = location
	return t
}

// SetCurrentUnix 设置当前的时间 Unix时间戳
func SetCurrentUnix(ts int64) Time {
	p := NewTime()
	p.Time = time.Unix(ts, 0)
	return p
}

type Time struct {
	Time  time.Time
	loc   *time.Location
	Error error
}

func NewTime() Time {
	return Time{
		Time: time.Now(),
		loc:  time.Local,
	}
}
func (t Time) Unix() int64 {
	return t.Time.Unix()
}
func (t Time) Now() time.Time {
	return t.Time
}

// BeforeSeconds 获取n秒前的时间
func (t Time) BeforeSeconds(seconds int64) Time {
	st, _ := time.ParseDuration(fmt.Sprintf("-%ds", seconds))
	t.Time = t.Time.Add(st)
	return t
}

// AfterSeconds 获取n秒后的时间
func (t Time) AfterSeconds(seconds int64) Time {
	st, _ := time.ParseDuration(fmt.Sprintf("+%ds", seconds))
	t.Time = t.Time.Add(st)
	return t
}

// BeforeMinute 获取n分钟前的时间
func (t Time) BeforeMinute(seconds int64) Time {
	st, _ := time.ParseDuration(fmt.Sprintf("-%dm", seconds))
	t.Time = t.Time.Add(st)
	return t
}

// AfterMinute 获取n分钟后的时间
func (t Time) AfterMinute(seconds int64) Time {
	st, _ := time.ParseDuration(fmt.Sprintf("+%dm", seconds))
	t.Time = t.Time.Add(st)
	return t
}

// BeforeHour 获取n小时前的时间
func (t Time) BeforeHour(hour int64) Time {
	st, _ := time.ParseDuration(fmt.Sprintf("-%dh", hour))
	t.Time = t.Time.Add(st)
	return t
}

// AfterHour 获取n小时后的时间
func (t Time) AfterHour(hour int64) Time {
	st, _ := time.ParseDuration(fmt.Sprintf("+%dh", hour))
	t.Time = t.Time.Add(st)
	return t
}

// BeforeDay 获取n天前的时间
func (t Time) BeforeDay(day int) Time {
	t.Time = t.Time.AddDate(0, 0, -day)
	return t
}

// AfterDay 获取n天后的时间
func (t Time) AfterDay(day int) Time {
	t.Time = t.Time.AddDate(0, 0, day)
	return t
}

// BeforeMonth 获取n月前的时间
func (t Time) BeforeMonth(n int) Time {
	t.Time = t.Time.AddDate(0, -n, 0)
	return t
}

// AfterMonth 获取n月后的时间
func (t Time) AfterMonth(n int) Time {
	t.Time = t.Time.AddDate(0, n, 0)
	return t
}

// BeforeYear 获取n年前的时间
func (t Time) BeforeYear(n int) Time {
	t.Time = t.Time.AddDate(-n, 0, 0)
	return t
}

// AfterYear 获取n年后的时间
func (t Time) AfterYear(n int) Time {
	t.Time = t.Time.AddDate(n, 0, 0)
	return t
}

// SetFormat 格式化
func (t Time) SetFormat(layout string) string {
	return t.Time.Format(layout)
}

// Month 获取当前月
func (t Time) Month() int64 {
	return t.MonthOfYear()
}

// MonthOfYear 获取本年的第几月
func (t Time) MonthOfYear() int64 {
	return int64(t.Time.In(t.loc).Month())
}
