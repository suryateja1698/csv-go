package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type Players struct {
	FirstName string
	LastName  string
	Country   string
	Role      string
}

func main() {
	file, err := os.Open("sample.csv")
	if err != nil {
		return
	}
	players, err := csvParse(file)
	if err != nil {
		log.Println("err in csv parse", err)
		return
	}
	for _, s := range players {
		fmt.Println("First name:", s.FirstName)
		fmt.Println("LastName:", s.LastName)
		fmt.Println("Country:", s.Country)
		fmt.Println("Role:", s.Role)
	}

}

func csvParse(file *os.File) ([]Players, error) {
	// Read CSV file using csv.NewReader
	csvReader := csv.NewReader(file)

	csvReader.LazyQuotes = true
	csvReader.Comma = ';'
	csvReader.FieldsPerRecord = -1
	csvReader.Read()
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	playersList := make([]Players, 0, len(data)-1)
	var rec Players
	for _, line := range data {
		for _, s := range line {
			fields := strings.Split(s, ",")
			for j, f := range fields {
				switch j {
				case 0:
					sanitizedValue := Sanitize(f)
					rec.FirstName = sanitizedValue
				case 1:
					sanitizedValue := Sanitize(f)
					rec.LastName = sanitizedValue
				case 2:
					sanitizedValue := Sanitize(f)
					rec.Country = sanitizedValue
				case 3:
					sanitizedValue := Sanitize(f)
					rec.Role = sanitizedValue
				}
			}
		}
		playersList = append(playersList, Players{
			FirstName: rec.FirstName,
			LastName:  rec.LastName,
			Country:   rec.Country,
			Role:      rec.Role,
		})
	}
	return playersList, nil

}

// Sanitize is used to sanitize the string from special characters and unreadable characters
func Sanitize(s string) string {
	sanitizedString := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, s)
	return sanitizedString
}
