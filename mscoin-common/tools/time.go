package tools

import "time"

// ISO 时间转化的格式
func ISO(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

func ToTimeString(mill int64) string {
	milli := time.UnixMilli(mill)
	return milli.Format("2006-01-02 15:04:05")
}

// ZeroTime 方便地获取当前日期的零时刻对应的UNIX时间戳，可以用于进行时间比较、计算时间间隔等相关操作
func ZeroTime() int64 {
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return date.UnixMilli()
}
