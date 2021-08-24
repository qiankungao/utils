package time

import (
	"fmt"
	"github.com/1975210542/utils/convert"
	"time"
)

//获取后面第几天的零点
func GetAfterDayZero(day int) int64 {
	nextDay := time.Now().AddDate(0, 0, day)
	nt, _ := time.ParseInLocation("2006-01-02", nextDay.Format("2006-01-02"), time.Local)
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
