package time

import (
	"fmt"
	"strings"
	"time"
)

//获取两个时间相差多少小时
func GetHourDiffer(start_time, end_time string) int64 {
	var hour int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", start_time, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", end_time, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 3600
		return hour
	} else {
		return hour
	}
}

func GetTimeStamp() (string, string) {
	timenow := time.Now()
	ts := timenow.Unix()
	format := "20060102150405.000"
	ss := timenow.Format(format)
	ss = strings.Replace(ss, ".", "", -1)
	return fmt.Sprintf("%v", ts), ss
}

func BeginOfTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
}

func EndOfTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
}

func BeginOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

func BeginOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func EndOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, 999999999, t.Location())
}

func BeginOfDayBefore(num int) time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day()+num, 0, 0, 0, 0, now.Location())
}

func BeforeOrAfterDaysOfTheTime(day *time.Time, num int) time.Time {
	return time.Date(day.Year(), day.Month(), day.Day()+num, day.Hour(), day.Minute(), day.Second(), 0, day.Location())
}
