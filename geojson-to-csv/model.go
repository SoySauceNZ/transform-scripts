package main

// Crash - crash line
type Crash struct {
	P struct {
		Crashseverity string `json:"crashSeverity"`
		Holiday       string `json:"holiday"`
		Light         string `json:"light"`
		Region        string `json:"region"`
		Weather       string `json:"weatherA"`
		SpeedLimit    int    `json:"speedLimit"`
	} `json:"properties"`
	G struct {
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
}

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
