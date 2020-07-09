package utils

import (
    "time"
)

func Reverseint64(list []int64) []int64 {
    for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
        list[i], list[j] = list[j], list[i]
    }
    return list
}

func ReverseString(list []string) []string {
    for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
        list[i], list[j] = list[j], list[i]
    }
    return list
}

func GetBetweenDates(startDate, endDate string) []string {
    d := []string{}
    timeFormatTpl := "2006-01-02 15:04:05"
    if len(timeFormatTpl) != len(startDate) {
        timeFormatTpl = timeFormatTpl[0:len(startDate)]
    }
    date, err := time.Parse(timeFormatTpl, startDate)
    if err != nil {
        // 时间解析，异常
        return d
    }
    date2, err := time.Parse(timeFormatTpl, endDate)
    if err != nil {
        // 时间解析，异常
        return d
    }
    if date2.Before(date) {
        // 如果结束时间小于开始时间，异常
        return d
    }
    // 输出日期格式固定
    timeFormatTpl = "2006-01-02"
    date2Str := date2.Format(timeFormatTpl)
    d = append(d, date.Format(timeFormatTpl))
    for {
        date = date.AddDate(0, 0, 1)
        dateStr := date.Format(timeFormatTpl)
        d = append(d, dateStr)
        if dateStr == date2Str {
            break
        }
    }
    return d
}

func GetYearAndMonth(dd time.Time) (start time.Time, end time.Time) {
    year, month, _ := dd.Date()
    loc := dd.Location()

    startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, loc)
    endOfMonth := startOfMonth.AddDate(0, 1, -1)
    return startOfMonth, endOfMonth
}
