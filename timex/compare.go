package time

import "time"

// Gt 是否大于
func (t Time) Gt(t2 time.Time) bool {
	return t.Time.After(t2)
}

// Lt 是否小于
func (t Time) Lt(t2 time.Time) bool {
	return t.Time.Before(t2)
}

// Eq 是否等于
func (t Time) Eq(t2 time.Time) bool {
	return t.Time.Equal(t2)
}

// Ne 是否不等于
func (t Time) Ne(t2 time.Time) bool {
	return !t.Eq(t2)
}

// Gte 是否大于等于
func (t Time) Gte(t2 time.Time) bool {
	return t.Gt(t2) || t.Eq(t2)
}

// Lte 是否小于等于
func (t Time) Lte(t2 time.Time) bool {
	return t.Lt(t2) || t.Eq(t2)
}

// Between 是否在两个时间之间(不包括这两个时间)
func (t Time) Between(start time.Time, end time.Time) bool {
	if t.Gt(start) && t.Lt(end) {
		return true
	}
	return false
}

// BetweenIncludedStart 是否在两个时间之间(包括开始时间)
func (t Time) BetweenIncludedStart(start time.Time, end time.Time) bool {
	if t.Gte(start) && t.Lt(end) {
		return true
	}
	return false
}

// BetweenIncludedEnd 是否在两个时间之间(包括结束时间)
func (t Time) BetweenIncludedEnd(start time.Time, end time.Time) bool {
	if t.Gt(start) && t.Lte(end) {
		return true
	}
	return false
}

// BetweenIncludedBoth 是否在两个时间之间(包括这两个时间)
func (t Time) BetweenIncludedBoth(start time.Time, end time.Time) bool {
	if t.Gte(start) && t.Lte(end) {
		return true
	}
	return false
}
