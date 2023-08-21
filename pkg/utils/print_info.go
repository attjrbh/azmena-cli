package utils

import (
	"fmt"
	"github.com/fatih/color"
	"go-ffprope/pkg/types"
)

func PrintInfo(info types.DurationInfo) {

	printCalculatedFilesInfo(info.OkPathsCount)
	printFailedFilesInfo(len(info.FailedPaths))

	splitDuration := splitDuration(info.TotalDuration)
	printDuration(splitDuration)
}

func emptyDuration(duration types.Duration) bool {
	if duration.Hours > 0 {
		return false
	}

	if duration.Minutes > 0 {
		return false
	}

	if duration.Hours > 0 {
		return false
	}

	return true
}

func printDuration(duration types.Duration) {
	isEmpty := emptyDuration(duration)
	fmt.Println()
	fmt.Println()
	title := ` TOTAL DURATION `
	var result string

	c := color.New()
	redBackground := c.Add(color.BgRed)
	redBackground.Add(color.FgWhite)

	redBackground.Print(title)

	infoStyle := color.New()
	infoStyle.Add(color.BgBlack)
	infoStyle.Add(color.FgHiWhite)
	fmt.Print("  ")

	if isEmpty {
		infoStyle.Print(" 0 seconds")
	}

	if duration.Hours > 0 {
		result += fmt.Sprintf("%d Hour", duration.Hours)
		if duration.Hours > 1 {
			result += "s"
		}
	}

	if duration.Minutes > 0 {
		if duration.Hours > 0 {
			result += ", "
		}
		result += fmt.Sprintf("%d Minute", duration.Minutes)
		if duration.Minutes > 1 {
			result += "s"
		}
	}

	if duration.Seconds > 0 {
		if duration.Minutes > 0 {
			result += ", "
		}
		result += fmt.Sprintf("%d Second", duration.Seconds)
		if duration.Seconds > 1 {
			result += "s"
		}
	}

	infoStyle.Printf(" %s ", result)

}

func printCalculatedFilesInfo(count int) {
	fmt.Println()
	c := color.New()
	titleStyle := c.Add(color.BgRed)
	titleStyle.Add(color.FgWhite)

	infoStyle := color.New()
	infoStyle.Add(color.BgBlack)
	infoStyle.Add(color.FgHiWhite)

	title := " NUMBER OF FILES "
	titleStyle.Print(title)
	fmt.Print(" ")
	result := fmt.Sprintf(" %d file", count)

	if count != 1 {
		result += "s"
	}
	result += " "

	infoStyle.Print(result)
}

func PrintFailedFilesInfo(count int) {
	if count == 0 {
		return
	}

	fmt.Printf("Couldn't get duration for %d files\n", count)
}

func printFailedFilesInfo(count int) {
	if count == 0 {
		return
	}

	fmt.Printf("Couldn't get duration for %d files\n", count)
}
