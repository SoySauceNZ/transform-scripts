package main

type crash struct {
	image       string
	latitude    float64
	longitude   float64
	severity    float64
	holiday     int
	light       string
	weather     string
	speed_limit string
	region      string
}

type image struct {
	image     string
	lattop    float64
	lngleft   float64
	latbottom float64
	lngright  float64
}
