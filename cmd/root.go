/*
Copyright © 2023 MUHAMMAD YASSER <mddysr@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"fmt"
	"math"
	"time"

	"sync"

	"github.com/mdyssr/azmena/pkg/types"

	"github.com/mdyssr/azmena/pkg/utils"

	"github.com/briandowns/spinner"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "azmena",
	Short: "A CLI tool to calculate video files total duration in directories",
	Long:  `Azmena helps you determine the total duration of your video fiels using only a single command.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		isFlat, err := cmd.Flags().GetBool("flat")
		if err != nil {
			fmt.Println("Error")
			os.Exit(1)
		}
		isMinimal, err := cmd.Flags().GetBool("minimal")
		if err != nil {
			fmt.Println("Error")
			os.Exit(0)
		}
		extensions, err := cmd.Flags().GetStringSlice("extensions")
		if err != nil {
			fmt.Println("Error")
			os.Exit(0)
		}

		validExtensions := utils.ValidateExtensions(extensions)
		if !validExtensions {
			fmt.Println("Extension(s) not supported.")
			os.Exit(0)
		}

		options := types.RunOptions{
			IsFlat:     isFlat,
			IsMinimal:  isMinimal,
			Extensions: extensions,
		}
		// fmt.Println(extensions)
		run(options)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.azmena.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolP("flat", "f", false, "Only calculate duration inside the root directory")
	// rootCmd.PersistentFlags().StringArrayP("extensions", "x", []string{}, "Only calculate duration for the specified file extensions")
	rootCmd.PersistentFlags().StringSliceP("extensions", "x", []string{}, "Only calculate duration for the specified file extensions")
	rootCmd.PersistentFlags().BoolP("minimal", "m", false, "Output duration in seconds with no output styling. It's useful if you want to pipe the result to another program")
}

func run(options types.RunOptions) {
	processStartTime := time.Now()

	s := spinner.New(spinner.CharSets[35], 150*time.Millisecond) // Build our new spinner
	s.Start()

	var totalDuration float64
	failedPaths := []string{}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: couldn't get working directory!")
		os.Exit(1)
	}

	paths, err := utils.GetFilePaths(wd, options)

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

	if options.IsMinimal {
		fmt.Println(math.Round(info.TotalDuration))
		return
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
