package system

import (
	"time"

	"github.com/fatih/color"
)

func Now() time.Time {
	return time.Now()
}
func Nowf() string {
	return Now().Format("2006-01-02 15:04:05")
}

func Between(init time.Time) time.Time {
	end := time.Now()
	diff := end.Sub(init)
	return time.Time{}.Add(diff)
}

func End(init time.Time, verbose bool) string {
	end := Between(init).Format("15:04:05")
	if verbose {
		color.Magenta("[complete: %v]", end)
	}
	return end
}
func Duration(init time.Time, verbose bool) int64 {
	duration := time.Now().Unix() - init.Unix()
	if verbose {
		color.Magenta("[duration: %v]", duration)
	}
	return duration
}
