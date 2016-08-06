package main

import (
    "net/http"
    "encoding/json"
    "math/rand"
    "time"

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
    // TODO: Query database to get all resolutions (don't have to be specific to this instrument - could just be a globally collected list)
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

func getCandles(instrumentId string, resolution string, startDate time.Time, endDate time.Time) Candles {
    // TODO: Retrieve candles from database
    var quote = Quote{Ask: rand.Float32()*1000, Bid: rand.Float32()*1000}
    return Candles{
        Candle{Time: time.Now(), OpenPrice: quote, ClosePrice: quote, LowPrice: quote, HighPrice: quote},
        Candle{Time: time.Now(), OpenPrice: quote, ClosePrice: quote, LowPrice: quote, HighPrice: quote},
        Candle{Time: time.Now(), OpenPrice: quote, ClosePrice: quote, LowPrice: quote, HighPrice: quote},
        Candle{Time: time.Now(), OpenPrice: quote, ClosePrice: quote, LowPrice: quote, HighPrice: quote},
        Candle{Time: time.Now(), OpenPrice: quote, ClosePrice: quote, LowPrice: quote, HighPrice: quote},
    }
}

func AllCandlesController(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var instrumentId string = vars["instrumentId"]
    var resolution string = vars["resolution"]
    var instrument = getInstrument(instrumentId)
    startDate, endDate := getDateRangeNCandlesAgo(time.Now(), 10, resolution)
    var candles = getCandles(instrumentId, resolution, startDate, endDate)
    var response = CandlesResponse{Instrument: instrument, Resolution: resolution, StartDate: startDate, EndDate: endDate, Candles: candles}

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        panic(err)
    }
}

func getDateRangeNCandlesAgo(now time.Time, candles int, resolution string) (time.Time, time.Time) {
    // TODO: needs a real impl + test
    return time.Now(), now
}

func RangeCandlesController(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    instrumentId := vars["instrumentId"]
    resolution := vars["resolution"]

    startDate, err := time.Parse(time.RFC3339, vars["startDate"])
    if err != nil {
        setHttpError(w, http.StatusBadRequest, "INVALID_START_DATE", "Invalid start date. Must conform to RFC3339.")
        return
    }

    endDate, err := time.Parse(time.RFC3339, vars["endDate"])
    if err != nil {
        setHttpError(w, http.StatusBadRequest, "INVALID_END_DATE", "Invalid end date. Must conform to RFC3339.")
        return
    }

    instrument := getInstrument(instrumentId)
    candles := getCandles(instrumentId, resolution, startDate, endDate)
    response := CandlesResponse{Instrument: instrument, Resolution: resolution, StartDate: startDate, EndDate: endDate, Candles: candles}

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        panic(err)
    }
}

type ErrorResponse struct {
    Error Error `json:"error"`
}

type Error struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

func setHttpError(w http.ResponseWriter, statusCode int, errorCode string, errorMessage string) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(statusCode)
    response := ErrorResponse{
        Error: Error{Code: errorCode, Message: errorMessage},
    }
    if err := json.NewEncoder(w).Encode(response); err != nil {
        panic(err)
    }
}
