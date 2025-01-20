package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type max_balance struct {
	balance float64
	when    string
}

func check_error(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	file_path := flag.String("input", "", "a file path")
	flag.Parse()

	csv_file, err := os.Open(*file_path)
	check_error(err)

	defer csv_file.Close()

	csv_reader := csv.NewReader(csv_file)

	records, err := csv_reader.ReadAll()

	check_error(err)

	var mb max_balance
	for _, record := range records {
		balance, err := strconv.ParseFloat(record[5], 32)
		if err != nil {
		}
		if balance > mb.balance {
			mb.balance = balance
			mb.when = record[0]
		}
	}
	fmt.Printf("%s: %.2f\n", mb.when, mb.balance)

}
