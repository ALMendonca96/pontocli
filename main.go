package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func enableLogging(cmd *cobra.Command) *log.Logger {
	var logger *log.Logger
	enableLogging, _ := cmd.Flags().GetBool("log")
	if enableLogging {
		logger = log.New(os.Stdout, "[pontocli] ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		logger = log.New(os.Stdout, "", 0)
		logger.SetOutput(io.Discard)
	}

	return logger
}

func main() {

	var rootCmd = &cobra.Command{Use: "pontocli",
		Short: "An app to register work hours",
	}

	var viewCmd = &cobra.Command{
		Use:   "view",
		Short: "View hours registered for a date",
		Run: func(cmd *cobra.Command, args []string) {

			logger := enableLogging(cmd)

			dateArg, _ := cmd.Flags().GetString("date")

			logger.Println("Setting up date...")
			dateArg = ResolveDate(dateArg)

			logger.Println("Validating date...")
			date, err := time.Parse("2006-01-02", dateArg)
			if err != nil {
				logger.Println("Wrong date format:", err)
				return
			}

			logger.Println("Loading hours...")
			hours, err := GetHours(logger, date)
			if err != nil {
				logger.Println("Error:", err)
			}

			logger.Println("Printing hours...")
			PrintWorkHours(date, hours)
		},
	}

	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Register a date and hour(s)",
		Run: func(cmd *cobra.Command, args []string) {
			logger := enableLogging(cmd)

			logger.Println("Getting arguments...")
			dateArg, _ := cmd.Flags().GetString("date")
			hourArg, _ := cmd.Flags().GetStringSlice("hour")

			logger.Println("Setting up date...")
			dateArg = ResolveDate(dateArg)

			logger.Printf("Setting up hours [%s]...\n", hourArg)
			if len(hourArg) == 0 {
				hourArg = append(hourArg, time.Now().Format("15:04"))
			}

			logger.Println("Validating date...")
			date, err := time.Parse("2006-01-02", dateArg)
			if err != nil {
				fmt.Printf("Error at date %s: %s", dateArg, err)
				return
			}

			logger.Println("Validating hours...")
			var hours []time.Time
			for _, h := range hourArg {
				hour, err := time.Parse("15:04", h)
				if err != nil {
					fmt.Printf("Error at hour %s: %s\n", h, err)
				}
				hours = append(hours, hour)
			}

			logger.Println("Saving date and hours...")
			err = SaveHours(logger, date, hours)
			if err != nil {
				logger.Fatal(err)
			}

			logger.Println("Retrieving date and hours...")
			savedHours, err := GetHours(logger, date)
			if err != nil {
				logger.Println("Error:", err)
			}

			logger.Println("Printing hours...")
			PrintWorkHours(date, savedHours)
		},
	}

	var deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete hours(s) from a date",
		Run: func(cmd *cobra.Command, args []string) {
			logger := enableLogging(cmd)

			logger.Println("Getting arguments...")
			dateArg, _ := cmd.Flags().GetString("date")
			hourArg, _ := cmd.Flags().GetStringSlice("hour")

			logger.Println("Setting up date...")
			dateArg = ResolveDate(dateArg)

			logger.Println("Setting up hour...")
			if len(hourArg) == 0 {
				return
			}

			logger.Println("Validating date...")
			date, err := time.Parse("2006-01-02", dateArg)
			if err != nil {
				logger.Printf("Error at date %s: %s", dateArg, err)
				return
			}

			logger.Println("Validating hours...")
			var hours []time.Time
			for _, h := range hourArg {
				hour, err := time.Parse("15:04", h)
				if err != nil {
					logger.Printf("Error at hour %s: %s\n", h, err)
				}
				hours = append(hours, hour)
			}

			logger.Println("Removing hours...")
			err = DeleteHours(logger, date, hours)
			if err != nil {
				logger.Fatal(err)
			}

			logger.Println("Retrieving date and hours...")
			savedHours, err := GetHours(logger, date)
			if err != nil {
				logger.Println("Error:", err)
			}

			logger.Println("Printing hours...")
			PrintWorkHours(date, savedHours)
		},
	}

	addCmd.Flags().String("date", "", "Specify the date in the format 'YYYY-MM-DD' or omit for the current date")
	addCmd.Flags().StringSliceP("hour", "", []string{}, "Specify the time in the format 'HH:mm' or omit for the current time")
	addCmd.Flags().Bool("log", false, "Enable logging")

	viewCmd.Flags().String("date", "", "Specify the date in the format 'YYYY-MM-DD' or omit for the current date")
	viewCmd.Flags().Bool("log", false, "Enable logging")

	deleteCmd.Flags().String("date", "", "Specify the date in the format 'YYYY-MM-DD' or omit for the current date")
	deleteCmd.Flags().StringSliceP("hour", "", []string{}, "Specify the time in the format 'HH:mm' or omit for the current time")
	deleteCmd.Flags().Bool("log", false, "Enable logging")

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(viewCmd)
	rootCmd.AddCommand(deleteCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
