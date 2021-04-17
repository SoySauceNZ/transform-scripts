# CAS Parser

Parser for the dataset located at https://catalogue.data.govt.nz/dataset/crash-analysis-system-cas-data1


## How to use

Set Input and output file
default input: `CAS.geojson`
default output: `CAS.csv`

Change output below
```go
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
    ...
```

Run using:
```
go run .
```

## Mapping
```go
// Severity - Severity to value
var Severity = map[string]string{
	"Non-Injury Crash": "0.0",
	"Minor Crash":      "0.33",
	"Serious Crash":    "0.66",
	"Fatal Crash":      "1.0",
}

// Weather - Order of how bad
var Weather = map[string]string{
	"Null":          "",
	"Fine":          "F",
	"Hail or Sleet": "HS",
	"Light rain":    "LR",
	"Heavy rain":    "HR",
	"Mist or Fog":   "MF",
	"Snow":          "S",
}

// Light - Order how Lighting
var Light = map[string]string{
	"Bright sun": "B",
	"Dark":       "D",
	"Overcast":   "O",
	"Twilight":   "T",
	"Unknown":    "",
}
```

## Example output
```
longitude,latitude,crash_severity,holiday,light,region,weather
174.756815,-36.899430,0.0,1,O,LR,Auckland Region
174.835880,-36.936018,0.0,0,T,F,Auckland Region
174.810287,-36.902124,0.0,0,B,LR,Auckland Region
174.789391,-36.936286,0.0,0,D,F,Auckland Region
174.312742,-35.730428,0.0,0,O,F,Northland Region
175.525505,-36.753408,0.66,0,B,F,Waikato Region
174.793909,-36.950342,0.0,0,D,F,Auckland Region
174.967990,-37.140820,0.33,0,T,F,Auckland Region
176.072647,-38.686122,0.0,0,B,F,Waikato Region
174.077823,-35.731221,0.0,0,D,F,Northland Region
174.778146,-41.283205,0.33,0,O,F,Wellington Region
```