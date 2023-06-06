package tools

import "time"

// ISO 时间转化的格式
func ISO(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}
