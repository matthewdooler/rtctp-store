package main

var resolutions = []string{
    "SECOND",
    "MINUTE",
    "MINUTE_2",
    "MINUTE_5",
    "MINUTE_10",
    "MINUTE_15",
    "MINUTE_30",
    "HOUR",
    "HOUR_2",
    "HOUR_3",
    "HOUR_4",
    "DAY"}

type Instrument struct {
    Id      string    `json:"id"`
    Links   Links     `json:"links"`
}

type Instruments []Instrument