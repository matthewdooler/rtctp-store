package main

import (
    "net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

type Link struct {
    Rel      string    `json:"rel"`
    Href     string    `json:"href"`
}

type Links []Link

func IndexController(w http.ResponseWriter, r *http.Request) {
    var links = Links{
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

type Status struct {
    Status      string    `json:"status"`
}

func StatusController(w http.ResponseWriter, r *http.Request) {
    var status = Status{Status: "OK"}
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(status); err != nil {
        panic(err)
    }
}

func InstrumentsController(w http.ResponseWriter, r *http.Request) {
    var instruments = Instruments{
        getInstrument("CS.D.GBPUSD.TODAY.IP"),
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(instruments); err != nil {
        panic(err)
    }
}

func getInstrument(id string) Instrument {
    return Instrument{Id: id, Links: instrumentLinks(id)}
}

func instrumentLinks(id string) Links {
	var links = Links{
        Link{Rel: "self", Href: config.BaseURI+"/instruments/"+id},
    }
	for _,resolution := range resolutions {
		links = append(links, Link{Rel: resolution, Href: config.BaseURI+"/instruments/"+id+"/"+resolution})
	}
	return links
}

func InstrumentController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var instrumentId string = vars["instrumentId"]
	var instrument = getInstrument(instrumentId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(instrument); err != nil {
        panic(err)
    }
}

func ResolutionController(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var instrumentId string = vars["instrumentId"]
    var resolution string = vars["resolution"]
    var instrument = getInstrument(instrumentId)
    var response = Resolution{Instrument: instrument, Resolution: resolution}

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        panic(err)
    }
}

