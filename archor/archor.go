package archor

import (
	"fmt"
	"time"
)

func NewAnchor() *Anchor {
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
		return fmt.Sprintf("%fms", float64(duration.Nanoseconds())/1000000.0)
	case "s":
		return fmt.Sprintf("%fs", duration.Seconds())
	case "ns":
		return fmt.Sprintf("%dns", duration.Nanoseconds())
	default:
		return duration.String()
	}
}
