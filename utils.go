package main

import (
	"fmt"
	"strings"
	"time"
)

func formatDuration(duration time.Duration) string {
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func FormatHours(hours []time.Time) string {
	var formattedHours string
	var timeDifference time.Duration
	for index, hour := range hours {
		if index > 0 && index%2 != 0 {
			timeDifference += hours[index].Sub(hours[index-1])
		}
		if index%2 == 0 {
			formattedHours += " > "
		} else {
			formattedHours += " < "
		}

		formattedHours += hour.Format("15:04")
	}

	formattedHours += " = " + formatDuration(timeDifference)
	return strings.TrimLeft(formattedHours, " ")
}

func PrintWorkHours(date time.Time, hours string) {
	fmt.Printf("Work hours for date %s:\n", date.Format("02-01-2006"))
	fmt.Println(hours)
}
