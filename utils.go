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
	fmt.Printf("Work hours for date %s:\n", date.Format("2006-01-02"))
	fmt.Println(hours)
}

func ResolveDate(dateArg string, lastLoggedDate time.Time) string {
	if dateArg == "" {
		return time.Now().Format("2006-01-02")
	}

	if dateArg == "today" {
		return time.Now().Format("2006-01-02")
	}

	if dateArg == "yesterday" {
		return time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	}

	if dateArg == "last" {
		return lastLoggedDate.Format("2006-01-02")
	}

	parsedDate, err := time.Parse("2006-01-02", dateArg)
	if err != nil {
		return ""
	}

	if parsedDate.IsZero() {
		return ""
	}

	if parsedDate.After(time.Now()) {
		return ""
	}

	return parsedDate.Format("2006-01-02")
}
