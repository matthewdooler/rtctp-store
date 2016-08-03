package main

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

// TODO: Map legacy urls (e.g., /prices/get?startDate=2015-09-13T17%3A15%3A00%2B00%3A00&epic=CS.D.GBPUSD.TODAY.IP&resolution=MINUTE_5&endDate=2015-09-14T21%3A00%3A00%2B00%3A00)

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "Instruments",
        "GET",
        "/instruments",
        InstrumentsIndex,
    },
}