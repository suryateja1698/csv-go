package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
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
	csvReader.FieldsPerRecord = -1
	csvReader.Read()
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	playersList := make([]Players, 0, len(data)-1)
	var rec Players
	for _, line := range data {
		rec.FirstName = line[0]
		rec.LastName = line[1]
		rec.Country = line[2]
		rec.Role = line[3]
		playersList = append(playersList, rec)
	}
	return playersList, nil

}
