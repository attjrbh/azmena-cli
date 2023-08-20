package utils

import (
	"context"

	"gopkg.in/vansante/go-ffprobe.v2"
)

func getFileDuration(src string) (float64, error) {

	data, err := ffprobe.ProbeURL(context.Background(), src)
	if err != nil {
		return 0, err
	}

	return data.Format.DurationSeconds, nil

}
