package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("CAS.geojson")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvfile, err := os.Create("CAS.csv")
	defer csvfile.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(csvfile)
	defer w.Flush()
	if err := w.Write([]string{
		"longitude",
		"latitude",
		"crash_severity",
		"holiday",
		"light",
		"weather",
		"speed_limit",
		"region",
	}); err != nil {
		log.Fatalln("error writing record to file", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var c Crash
		err := json.Unmarshal([]byte(strings.TrimSuffix(scanner.Text(), ",")), &c)
		if err != nil {
			panic(err)
		}

		holiday := "1"
		if c.P.Holiday == "" {
			holiday = "0"
		}

		weather := Weather[c.P.Weather]
		if weather == "" {
			continue
		}

		light := Light[c.P.Light]
		if light == "" {
			continue
		}
		row := []string{
			fmt.Sprintf("%f", c.G.Coordinates[0]),
			fmt.Sprintf("%f", c.G.Coordinates[1]),
			Severity[c.P.Crashseverity],
			holiday,
			light,
			weather,
			fmt.Sprintf("%d", c.P.SpeedLimit),
			c.P.Region,
		}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
