package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func writeOutput(crashes []crash, output string) error {
	data := [][]string{{"image", "latitude", "longitude", "crash_severity", "holiday", "light", "weather", "speed_limit", "region"}}

	for _, crash := range crashes {
		slice := []string{
			crash.image,
			fmt.Sprintf("%f", crash.latitude),
			fmt.Sprintf("%f", crash.longitude),
			fmt.Sprintf("%f", crash.severity),
			fmt.Sprintf("%d", crash.holiday),
			crash.light,
			crash.weather,
			crash.speed_limit,
			crash.region,
		}
		data = append(data, slice)
	}

	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			return err
		}
	}

	return nil
}
