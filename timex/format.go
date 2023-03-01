package time

import "time"

// Format 今天此刻格式化
func (t Time) Format() string {
	return t.Time.Format(DateTimeFormat)
}

// ToDateFormat 今天此刻日期
func (t Time) ToDateFormat() string {
	return t.Time.Format(DateFormat)
}

// ToTimeFormat 今天此刻时间
func (t Time) ToTimeFormat() string {
	return t.Time.Format(TimeFormat)
}

// Timestamp 今天此刻时间戳
func (t Time) Timestamp() int64 {
	return t.Time.Unix()
}

// TimestampWithSecond 今天此刻时间戳
func (t Time) TimestampWithSecond() int64 {
	return t.Time.Unix()
}

// TimestampWithMillisecond 今天毫秒级时间戳
func (t Time) TimestampWithMillisecond() int64 {
	return t.Time.UnixNano() / int64(time.Millisecond)
}

// TimestampWithMicrosecond 今天微秒级时间戳
func (t Time) TimestampWithMicrosecond() int64 {
	return t.Time.UnixNano() / int64(time.Microsecond)
}

// TimestampWithNanosecond 今天纳秒级时间戳
func (t Time) TimestampWithNanosecond() int64 {
	return t.Time.UnixNano()
}
