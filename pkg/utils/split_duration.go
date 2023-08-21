package utils

import (
	"go-ffprope/pkg/types"
	"math"
)

func splitDuration(d float64) types.Duration {
	seconds := int(math.Round(d))
	hours := seconds / (60 * 60)
	seconds -= hours * (60 * 60)
	minutes := seconds / 60
	seconds -= minutes * 60

	return types.Duration{
		Seconds: seconds,
		Minutes: minutes,
		Hours:   hours,
	}
}
