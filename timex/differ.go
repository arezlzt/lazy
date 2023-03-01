package time

import "time"

// DiffInHour 相差多少小时
func (t Time) DiffInHour(t2 time.Time) (hour int64) {
	t2.Before(t.Time)
	diff := t.Time.Unix() - t2.Unix()
	hour = diff / 3600
	return hour
}

// DiffInHourWithAbs 相差多少小时(绝对值)
func (t Time) DiffInHourWithAbs(t2 time.Time) (hour int64) {
	t.Time.Before(t2)
	diff := t2.Unix() - t.Time.Unix()
	hour = diff / 3600
	if hour > 0 {
		return hour
	}
	t2.Before(t.Time)
	diff = t.Time.Unix() - t2.Unix()
	hour = diff / 3600
	return hour
}

// DiffInMinutes 相差多少分钟
func (t Time) DiffInMinutes(t2 time.Time) (hour int64) {
	t2.Before(t.Time)
	diff := t.Time.Unix() - t2.Unix()
	hour = diff / 60
	return hour
}

// DiffInMinutesWithAbs 相差多少分钟(绝对值)
func (t Time) DiffInMinutesWithAbs(t2 time.Time) (hour int64) {
	t.Time.Before(t2)
	diff := t2.Unix() - t.Time.Unix()
	hour = diff / 60
	if hour > 0 {
		return hour
	}
	t2.Before(t.Time)
	diff = t.Time.Unix() - t2.Unix()
	hour = diff / 60
	return hour
}

// DiffInSecond 相差多少秒
func (t Time) DiffInSecond(t2 time.Time) (hour int64) {
	t2.Before(t.Time)
	diff := t.Time.Unix() - t2.Unix()
	hour = diff
	return hour
}

// DiffInSecondWithAbs 相差多少秒(绝对值)
func (t Time) DiffInSecondWithAbs(t2 time.Time) (hour int64) {
	t.Time.Before(t2)
	diff := t2.Unix() - t.Time.Unix()
	hour = diff
	if hour > 0 {
		return hour
	}
	t2.Before(t.Time)
	diff = t.Time.Unix() - t2.Unix()
	hour = diff
	return hour
}
