# csv-transform

Transform CSV by migrating anz_crash.csv and image.csv

## How does it work?

Input is two csv and does computation to pre process our data. The csv transformer uses a viewport of two coordinates to find all crashes within these two coords. The transformer links images number with the count of crashes and severity of the crash.

The average severity of the crash is calculated by averaging the severity. Severity is converted from enum values of `fatality`, `serious_injury`, `property_damage`, `minor_injury` in order of severity. If there is no crash in the image the severity is zero. The team has chosen to give each severity a arbitrary value to indicate average.

```go
const (
	fatality       crashSeverity = 1
	seriousInjury  crashSeverity = 0.75
	propertyDamage crashSeverity = 0.5
	minorInjury    crashSeverity = 0.25
	none           crashSeverity = 0.0
)
```

## Inputs

```bash
csv-transform [input-crashes] [input-coords] [output]
```

The inputs are
- `input-crashes`
    - csv with `latitude,longitude,,,severity`
    - location of crashes using latitude and longitude
    - only the order of the csv header matters, empty headers shown above are optional
    - 1st row is ignored
- `input-coords`
    - csv with `,,lat0,lng0,lat1,lng1`
    - alternate corners of each image in latitude and longitude
    - only the order of the csv header matters, empty headers shown above are optional
    - 1st row is ignored
- `output`
    - output csv filename
    - csv headers `index,count,avg_severity`
    - index increments from zero to last entry
    - `count` number of crashes 
    - `avg_severity` average severity of crashes


## Install

- Have Go installed or use compiled binary
- `git clone` repo

### Run only
```bash
go run . ./anz_crash.csv ./image.csv ./output.csv
```

### Build and Run
```bash
go get
go build
```

Running built binary:
```bash
csv-transform ./anz_crash.csv ./image.csv ./output.csv
```


