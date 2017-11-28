package utils

import (
	"time"
	"math"
)

/**
 * 获取今天的日期
 * yyyymmdd
 */
func Today() string {
	return time.Now().Format("20060102")
}

/**
 * 将指定日期转换为字符串格式
 * yyyymmdd
 */
func DateToStr(item time.Time) string {
	return item.Format("20060102")
}

/**
 * 将字符串格式的时间转换为Time类型
 * date: yyyymmdd
 */
func StrToDate(date string) time.Time {
	tm2, _ := time.Parse("20060102", date)

	return tm2
}

/**
 * 基于某个日期甲减几天后的日期
 */
func AddDays(current string, delta int) string {
	d, _ := time.ParseDuration("24h")
	cur := StrToDate(current)

	return DateToStr(cur.Add(time.Duration(delta) * d))
}

/**
 * 加减几天的日期(自动跳过周末)
 */
func AddDaysExceptWeekend(current string, delta int) string {
	var target time.Time
	var d time.Duration
	cur := StrToDate(current)

	if delta >= 0 {
		d, _ = time.ParseDuration("24h")
	} else {
		d, _ = time.ParseDuration("-24h")
	}

	dtIdx := 0
	absDelta := int(math.Abs(float64(delta)))

	for idx:=0; idx < absDelta*2; idx++ {
		target = cur.Add(time.Duration(idx) * d)
		if time.Saturday != target.Weekday() && time.Sunday != target.Weekday() {
			dtIdx++
		}

		if dtIdx >= absDelta {
			break
		}

	}

	return DateToStr(target)
}














