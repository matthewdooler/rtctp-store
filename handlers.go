package main

import (
    "fmt"
    "net/http"
	"encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "rtctp store") // TODO: links to collections etc
}

func InstrumentsIndex(w http.ResponseWriter, r *http.Request) {
    instruments := Instruments{
        Instrument{Epic: "CS.D.GBPUSD.TODAY.IP"},
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(instruments); err != nil {
        panic(err)
    }
}