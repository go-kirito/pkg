/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/9/14 15:19
 */
package period

import (
	"fmt"
	"strconv"
	"time"
)

func PeriodToYM(period int32) string {
	return fmt.Sprintf("%d年%02d期", period/100, period%100)
}

//根据开始和结束时间，返回时间列表
func GetDateRange(start time.Time, end time.Time) []string {
	if start.Unix() > end.Unix() {
		return nil
	}

	if start.Unix() == end.Unix() {
		return []string{start.Format("2006-01-02")}
	}

	dateList := make([]string, 0)
	for start.Unix() <= end.Unix() {
		dateList = append(dateList, start.Format("2006-01-02"))
		start = start.Add(24 * time.Hour)
	}

	return dateList
}

func NextPeriod(period int32) int32 {
	year := period / 100
	month := period % 100

	if month == 12 {
		year = year + 1
		month = 1
	} else {
		month = month + 1
	}

	return year*100 + month
}

func PrePeriod(period int32) int32 {
	year := period / 100
	month := period % 100

	if month == 1 {
		year = year - 1
		month = 12
	} else {
		month = month - 1
	}

	return year*100 + month
}

func NowPeriod() int32 {
	now := time.Now().Format("200601")
	period, _ := strconv.ParseInt(now, 10, 64)
	return int32(period)
}
