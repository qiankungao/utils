package time

import (
	"fmt"
	"github.com/qiankungao/utils/convert"
	"time"
)
//时间戳转日期
func Unix2time(tmp int64) string {
	return time.Unix(tmp, 0).Format("2006-01-02 15:04:05")
}

//获取后面第几天的零点 0 表示当天 1 表示明天。。。。
func GetAfterDayZero(day int) int64 {
	nextDay := time.Now().AddDate(0, 0, day)
	nt, _ := time.ParseInLocation("2006-01-02", nextDay.Format("2006-01-02"), time.Local)
	return nt.Unix()
}

//获取周几的零点时间，若是大于，则取下周
func GetWeekDayZero(day int) int64 {
	wd := convert.StrTo(fmt.Sprintf("%d", time.Now().Weekday())).MustInt()
	if wd == 0 {
		wd = 7
	}
	if wd >= day { //若是大于，则取下周
		return GetNextWeekDayZero(day)
	}
	nt, _ := time.ParseInLocation("2006-01-02", time.Now().AddDate(0, 0, day-wd).Format("2006-01-02"), time.Local)
	return nt.Unix()
}

//获取下周几的零点时间
func GetNextWeekDayZero(day int) int64 {
	//step:1 获取当天是周几
	wk := convert.StrTo(fmt.Sprintf("%d", time.Now().Weekday())).MustInt()
	if wk == 0 {
		wk = 7
	}

	nt, _ := time.ParseInLocation("2006-01-02", time.Now().AddDate(0, 0, 7-wk+day).Format("2006-01-02"), time.Local)
	return nt.Unix()
}

//获取本月几号的零点时间，若是大于，则取下周
func GetMonthDay(day int) int64 {
	currentYear, currentMonth, currentDay := time.Now().Date()
	if currentDay >= day {
		return GetNextMonthDay(day)
	}
	return time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.Now().Location()).AddDate(0, 0, day-1).Unix()
}

//获取下个月几号零点的时间戳
func GetNextMonthDay(day int) int64 {
	currentYear, currentMonth, _ := time.Now().Date()
	return time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.Now().Location()).AddDate(0, 1, day-1).Unix()
}
