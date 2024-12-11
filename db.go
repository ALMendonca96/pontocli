package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/boltdb/bolt"
)

func InitDB(logger *log.Logger) (*bolt.DB, error) {
	logger.Println("Initializing database...")
	return bolt.Open("pontocli.db", 0600, nil)
}

func GetHours(logger *log.Logger, date time.Time) (string, error) {
	logger.Printf("Getting hours for date %s...\n", date.Format("2006-01-02"))

	db, err := InitDB(logger)
	if err != nil {
		return "", err
	}
	defer db.Close()

	var formattedHours string
	err = db.View(func(tx *bolt.Tx) error {

		logger.Println("Retrieving bucket...")
		bucket := tx.Bucket([]byte("pontocli"))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}

		logger.Println("Retrieving data...")
		workDayData := bucket.Get([]byte(date.Format("2006-01-02")))
		if workDayData == nil {
			return fmt.Errorf("workday not found for date")
		}

		logger.Println("Decoding data...")
		var retrievedWorkDay WorkDay
		err := json.Unmarshal(workDayData, &retrievedWorkDay)
		if err != nil {
			return err
		}

		logger.Println("Formatting hours...")
		formattedHours = FormatHours(retrievedWorkDay.Hours)

		return nil
	})

	if err != nil {
		return "", err
	}

	return formattedHours, nil
}

func GetLastLoggedDate(logger *log.Logger) (time.Time, error) {
	logger.Println("Getting last logged date...")

	// Inicializa o banco de dados
	db, err := InitDB(logger)
	if err != nil {
		return time.Time{}, err // Retorna time zero e o erro
	}
	defer db.Close()

	var lastLoggedDate time.Time

	// Usando transação de leitura para acessar dados no bucket "logins" (ajuste conforme seu banco)
	err = db.View(func(tx *bolt.Tx) error {
		logger.Println("Retrieving bucket 'logins'...")

		bucket := tx.Bucket([]byte("pontocli"))
		if bucket == nil {
			return fmt.Errorf("bucket 'pontocli' not found")
		}

		// Buscando a última data de login. O exemplo assume que as chaves são as datas (formato YYYY-MM-DD)
		logger.Println("Retrieving last login data...")

		var maxDate time.Time
		bucket.ForEach(func(k, v []byte) error {
			// Cada chave (k) é uma data de login, e o valor (v) é o objeto de login (pode ser um JSON com detalhes)
			date, err := time.Parse("2006-01-02", string(k)) // Assuming the date is stored in "YYYY-MM-DD"
			if err != nil {
				return err
			}

			// Compara a data para encontrar a mais recente
			if date.After(maxDate) {
				maxDate = date
			}

			return nil
		})

		if maxDate.IsZero() {
			return fmt.Errorf("no date found")
		}

		lastLoggedDate = maxDate
		return nil
	})

	if err != nil {
		return time.Time{}, err
	}

	logger.Printf("Last login date: %s\n", lastLoggedDate.Format("2006-01-02"))
	return lastLoggedDate, nil
}

func SaveHours(logger *log.Logger, date time.Time, hours []time.Time) error {
	db, err := InitDB(logger)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		logger.Println("Retrieving bucket...")
		bucket, err := tx.CreateBucketIfNotExists([]byte("pontocli"))
		if err != nil {
			return err
		}

		logger.Println("Retrieving data...")
		existingData := bucket.Get([]byte(date.Format("2006-01-02")))

		var workDay WorkDay
		if existingData == nil {
			workDay = WorkDay{
				Date:  date,
				Hours: hours,
			}
		} else {
			logger.Println("Decoding data...")
			err := json.Unmarshal(existingData, &workDay)
			if err != nil {
				return err
			}

			logger.Println("Updating hours...")
			workDay.Hours = append(workDay.Hours, hours...)
		}

		logger.Println("Sorting hours...")
		sort.Slice(workDay.Hours, func(i, j int) bool {
			return workDay.Hours[i].Before(workDay.Hours[j])
		})

		logger.Println("Encoding data...")
		workDayJSON, err := json.Marshal(workDay)
		if err != nil {
			log.Fatal(err)
		}

		logger.Println("Storing data...")
		err = bucket.Put([]byte(date.Format("2006-01-02")), workDayJSON)
		if err != nil {
			return err
		}

		return nil
	})
}

func DeleteHours(logger *log.Logger, date time.Time, hours []time.Time) error {
	logger.Printf("Deleting %s for date %s...\n", date.Format("2006-01-02"), FormatHours(hours))

	db, err := InitDB(logger)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		logger.Println("Retrieving bucket...")
		bucket, err := tx.CreateBucketIfNotExists([]byte("pontocli"))
		if err != nil {
			return err
		}

		logger.Println("Retrieving data...")
		existingData := bucket.Get([]byte(date.Format("2006-01-02")))
		if existingData == nil {
			return nil
		}

		logger.Println("Decoding data...")
		var workDay WorkDay
		err = json.Unmarshal(existingData, &workDay)
		if err != nil {
			return err
		}

		logger.Println("Updating work hours...")
		var updatedHours []time.Time
		for _, hour := range hours {
			logger.Printf("Updating work hours [%s]...\n", hour.Format("15:04"))
			for _, savedHour := range workDay.Hours {
				logger.Printf("Comparing work hour with saved hour [%s][%s]...\n", hour.String(), savedHour.String())

				if savedHour.Hour() != hour.Hour() ||
					savedHour.Minute() != hour.Minute() ||
					savedHour.Second() != hour.Second() {
					logger.Println("Ignoring work hour...")
					updatedHours = append(updatedHours, savedHour)
				}
			}
		}

		workDay.Hours = updatedHours

		logger.Println("Sorting work hours...")
		sort.Slice(workDay.Hours, func(i, j int) bool {
			return workDay.Hours[i].Before(workDay.Hours[j])
		})

		logger.Println("Encoding data...")
		workDayJSON, err := json.Marshal(workDay)
		if err != nil {
			log.Fatal(err)
		}

		logger.Println("Applying changes...")
		err = bucket.Put([]byte(date.Format("2006-01-02")), workDayJSON)
		if err != nil {
			return err
		}

		return nil
	})
}
