package archor

import (
	"fmt"
	"time"
)

func New() *Anchor {
	return &Anchor{StartTime: time.Now()}
}

func (a *Anchor) Since(format string) string {
	duration := time.Since(a.StartTime)

	return FormatDuration(duration, format)
}

func (a *Anchor) Duration() time.Duration {
	return time.Since(a.StartTime)
}

func FormatDuration(duration time.Duration, format string) string {

	switch format {
	case "ms":
		return fmt.Sprintf("%.2fms", float64(duration.Nanoseconds())/float64(time.Millisecond)) // 1,000,000(1e6)
	case "s":
		return fmt.Sprintf("%fs", duration.Seconds())
	case "ns":
		return fmt.Sprintf("%dns", duration.Nanoseconds())
	default:
		return duration.String()
	}
}
