package utils

import (
	"sync"

	"github.com/mdyssr/azmena/pkg/types"
)

func GetFileStatus(path string, wg *sync.WaitGroup, ch chan types.FileStatus) {
	defer wg.Done()

	status := types.FileStatus{}

	duration, err := getFileDuration(path)
	if err == nil {
		status.Ok = true
		status.Duration = duration
	}

	ch <- status

}
