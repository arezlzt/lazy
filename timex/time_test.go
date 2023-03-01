package time

import (
	"fmt"
	"testing"
	"time"
)

func TestNow2(t *testing.T) {
	fmt.Println(NewTime().Unix())
}
func TestTime_BeforeSeconds(t *testing.T) {
	fmt.Println(NewTime().Unix())
	fmt.Println(NewTime().BeforeSeconds(10).Unix())
}
func TestTime_BeforeMonth(t *testing.T) {
	fmt.Println(NewTime().BeforeMonth(1).Unix())
}
func TestTime_EndOfDay(t *testing.T) {
	fmt.Println(NewTime().EndOfDay().Unix())
}
func TestTime_StartOfDay(t *testing.T) {
	fmt.Println(NewTime().BeforeYear(1).BeforeMonth(1).EndOfMonth().Format())
	fmt.Println(NewTime().BeforeYear(1).BeforeMonth(1).StartOfMonth().Format())

	fmt.Println(NewTime().BeforeYear(1).BeforeMonth(1).EndOfQuarter().Format())
	fmt.Println(NewTime().BeforeYear(1).BeforeMonth(1).StartOfQuarter().Format())

	fmt.Println(NewTime().BeforeYear(1).BeforeMonth(1).EndOfYear().Format())
	fmt.Println(NewTime().BeforeYear(1).BeforeMonth(1).StartOfYear().Format())
}
func TestSetCurrentParse(t *testing.T) {
	fmt.Println(SetCurrentParse("2023-01-01 12:12:12").Unix())
	fmt.Println(SetCurrent(time.Now()).Unix())
	fmt.Println(SetCurrentUnix(1555555555).Format())
}
