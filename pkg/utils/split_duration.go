package utils

import (
	"go-ffprope/pkg/types"
	"math"
)

func splitDuration(d float64) types.Duration {
	// we first round the seconds to the nearest integer
	// for example: 1.1 => 1, and 2.5 => 3
	// days := seconds / (60 * 60 * 24)
	// seconds -= days * (60 * 60 * 24)
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
