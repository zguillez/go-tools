package system

import (
	"fmt"
	"math"
	"time"

	"github.com/fatih/color"
)

func ByteCount(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%dB", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	bytes := float64(b) / float64(div)
	if s := fmt.Sprintf("%c", "KMGTPE"[exp]); s == "K" {
		return fmt.Sprintf("%v%c", int(math.Round(bytes)), "KMGTPE"[exp])
	}
	return fmt.Sprintf("%.1f%c", bytes, "KMGTPE"[exp])
}

var timeProcessChan chan time.Time

func TimeProcessIn() {
	timeProcessChan = make(chan time.Time, 2)
	timeProcessChan <- time.Now()
}

func TimeProcessOut() {
	timeProcessChan <- time.Now()
	close(timeProcessChan)

	t1 := <-timeProcessChan
	t2 := <-timeProcessChan

	diff := t2.Sub(t1)
	out := time.Time{}.Add(diff)
	color.Yellow("[processed in %v]", out.Format("15:04:05"))
}
