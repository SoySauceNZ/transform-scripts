package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

var crashesArray []crash

func read(args []string) error {
	crashesFile, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer crashesFile.Close()

	crashes := csv.NewReader(crashesFile)
	if _, err := crashes.Read(); err != nil {
		return err
	}

	crashesLines, err := crashes.ReadAll()
	if err != nil {
		return err
	}

	for _, line := range crashesLines {
		latitude, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			return err
		}

		longitude, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			return err
		}

		severity, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			return err
		}

		holiday, err := strconv.Atoi(line[3])
		if err != nil {
			return err
		}

		crashesArray = append(crashesArray, crash{
			latitude:    latitude,
			longitude:   longitude,
			severity:    severity,
			holiday:     holiday,
			light:       line[4],
			weather:     line[5],
			speed_limit: line[6],
			region:      line[7],
		})
	}

	imagesFile, err := os.Open(args[1])
	if err != nil {
		return err
	}
	defer imagesFile.Close()

	images := csv.NewReader(imagesFile)
	if _, err := images.Read(); err != nil {
		return err
	}

	imagesLines, err := images.ReadAll()
	if err != nil {
		return err
	}

	var output []crash

	for _, line := range imagesLines {

		latitude0, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			return err
		}

		longitude0, err := strconv.ParseFloat(line[3], 64)
		if err != nil {
			return err
		}

		latitude1, err := strconv.ParseFloat(line[4], 64)
		if err != nil {
			return err
		}

		longitude1, err := strconv.ParseFloat(line[5], 64)
		if err != nil {
			return err
		}

		crashes := getCrash(image{
			image:     line[0],
			lattop:    latitude0,
			lngleft:   longitude0,
			latbottom: latitude1,
			lngright:  longitude1,
		})

		output = append(output, crashes...)
	}

	if err := writeOutput(output, args[2]); err != nil {
		return err
	}

	return nil
}
