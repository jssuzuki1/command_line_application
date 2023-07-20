package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// Sum function calculates the sum of the values in the data slice.
func sum(data []float64) float64 {
	sum := 0.0

	for _, v := range data {
		sum += v
	}

	return sum
}

// Min function calculates the minimum of the values in the data slice.
func min(data []float64) float64 {
	min := data[0]

	for _, v := range data {
		if v < min {
			min = v
		}
	}

	return min
}

// Avg function calculates the average of the values in the data slice.
func mean(data []float64) float64 {
	return sum(data) / float64(len(data))
}

// Max function calculates the maximum of the values in the data slice.
func max(data []float64) float64 {
	max := data[0]

	for _, v := range data {
		if v > max {
			max = v
		}
	}

	return max
}

// Uniform signature
type statsFunc func(data []float64) float64

// csv2float reads a CSV file and returns a slice of float64 values.
func csv2float(r io.Reader, column int) ([]float64, error) {
	cr := csv.NewReader(r)

	// Adjust for 0-based index because columns start from 1.
	column--

	// Read CSV data
	allData, err := cr.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("cannot read data from file: %w", err)
	}

	var data []float64

	for i, row := range allData {
		// skip first row of column names
		if i == 0 {
			continue
		}

		if len(row) <= column {
			// Files does not have that many columns
			return nil,
				fmt.Errorf("%w: File has only %d columns",
					ErrInvalidColumn,
					len(row))
		}

		// Convert data to float number
		// If any of the data is not a number, it will return an error that it is not a number.
		v, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}

		data = append(data, v)
	}

	return data, nil
}

// func main() {
// 	// Open the CSV file
// 	file, err := os.Open("housesInput.csv")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	// Create a new CSV reader
// 	reader := csv.NewReader(file)

// 	// Read all the records from the CSV file
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		fmt.Println("Error reading CSV file:", err)
// 		return
// 	}

// 	// Set an initial minimum value
// 	minValue := 0.0

// 	for _, record := range records {
// 		// Parse the value from the record
// 		value, err := strconv.ParseFloat(record[1], 64)
// 		if err != nil {
// 			fmt.Println("Error parsing value:", err)
// 			continue
// 		}

// 		// Update the minimum value if necessary
// 		if value < minValue || minValue == 0.0 {
// 			minValue = value
// 		}
// 	}

// 	// Print the minimum value
// 	fmt.Println("Minimum value:", minValue)
// }
