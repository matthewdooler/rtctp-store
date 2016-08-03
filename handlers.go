package main

import (
    "net/http"
	"encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
    links := Links{
        Link{Rel: "self", Href: config.BaseURI+"/"},
        Link{Rel: "status", Href: config.BaseURI+"/status"},
        Link{Rel: "instruments", Href: config.BaseURI+"/instruments"},
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
        Instrument{Id: "CS.D.GBPUSD.TODAY.IP", Links: instrumentLinks("CS.D.GBPUSD.TODAY.IP")},
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(instruments); err != nil {
        panic(err)
    }
}

func instrumentLinks(id string) Links {
	links := Links{
        Link{Rel: "self", Href: config.BaseURI+"/instruments/"+id},
    }
	for _,resolution := range resolutions {
		links = append(links, Link{Rel: resolution, Href: config.BaseURI+"/instruments/"+id+"/"+resolution})
	}
	return links
}