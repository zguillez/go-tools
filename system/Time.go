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

func End(init time.Time, verbose bool) {
	end := Between(init)
	if verbose {
		color.Magenta("[complete: %v]", end.Format("15:04:05"))
	}
}
