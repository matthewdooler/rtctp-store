package main

import (
    "net/http"
	"encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
    baseuri := "http://127.0.0.1:8042" // TODO: needs to be in config
    links := Links{
        Link{Rel: "self", Href: baseuri+"/"},
        Link{Rel: "instruments", Href: baseuri+"/instruments"},
        Link{Rel: "status", Href: baseuri+"/status"},
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(links); err != nil {
        panic(err)
    }
}

func StatusIndex(w http.ResponseWriter, r *http.Request) {
    status := Status{Status: "OK"}
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(status); err != nil {
        panic(err)
    }
}

func InstrumentsIndex(w http.ResponseWriter, r *http.Request) {
    instruments := Instruments{
        Instrument{Id: "CS.D.GBPUSD.TODAY.IP"},
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(instruments); err != nil {
        panic(err)
    }
}