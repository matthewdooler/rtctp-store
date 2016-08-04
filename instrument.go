package main

import (
    "time"
)

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
    Id    string `json:"id"`
    Links Links  `json:"links"`
}

type Instruments []Instrument

type Candle struct {
	Time       time.Time `json:"time"`
	OpenPrice  Quote     `json:"openPrice"`
	ClosePrice Quote     `json:"closePrice"`
	LowPrice   Quote     `json:"lowPrice"`
	HighPrice  Quote     `json:"highPrice"`
}

type Candles []Candle

type Quote struct {
	Ask float32	`json:"ask"`
	Bid float32	`json:"bid"`
}

type CandlesResponse struct {
	Instrument Instrument `json:"instrument"`
	Resolution string     `json:"resolution"`
	StartDate  time.Time  `json:"startDate"`
	EndDate    time.Time  `json:"endDate"`
	Candles    Candles    `json:"candles"`
}