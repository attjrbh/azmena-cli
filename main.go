package main

import (
	"fmt"
	"math"
	"time"

	"go-ffprope/pkg/types"
	"go-ffprope/pkg/utils"
	"os"
	"sync"

	"github.com/briandowns/spinner"
)

func main() {

	run()

}

func run() {
	processStartTime := time.Now()

	s := spinner.New(spinner.CharSets[26], 150*time.Millisecond) // Build our new spinner
	s.Start()

	var totalDuration float64
	failedPaths := []string{}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: couldn't get working directory!")
		os.Exit(1)
	}

	paths, err := utils.GetFilePaths(wd)

	if err != nil {
		fmt.Println("Error: getting file paths!")
		os.Exit(1)
	}

	// define a WaitGroup
	var wg sync.WaitGroup
	statusChan := make(chan types.FileStatus, len(paths))

	for _, path := range paths {
		wg.Add(1)
		go func(path string, wg *sync.WaitGroup, ch chan types.FileStatus) {
			utils.GetFileStatus(path, wg, ch)

		}(path, &wg, statusChan)

	}
	wg.Wait()
	close(statusChan)
	s.Stop()

	for status := range statusChan {
		if status.Ok {
			totalDuration += status.Duration
		} else {
			failedPaths = append(failedPaths, status.Path)
		}
	}

	processDuration := time.Since(processStartTime)

	info := types.DurationInfo{
		FailedPaths:   failedPaths,
		TotalDuration: totalDuration,
		OkPathsCount:  len(paths) - len(failedPaths),
	}
	var processDurationMsg string = "Process finished in ~ "
	processDurationRounded := roundFloat(processDuration.Seconds(), 2)

	if int(processDurationRounded) == int(math.Round(processDurationRounded)) {
		processDurationMsg += fmt.Sprintf("%d", int(processDuration.Seconds()))
	} else {
		processDurationMsg += fmt.Sprintf("%.2f", processDuration.Seconds())
	}

	if int(processDuration.Seconds()) == 1 {
		processDurationMsg += " second"
	} else {
		processDurationMsg += " seconds"
	}

	utils.PrintInfo(info)
	fmt.Println()
	fmt.Println()
	fmt.Println(processDurationMsg)
}

func roundFloat(value float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Trunc(value*ratio) / ratio
}
