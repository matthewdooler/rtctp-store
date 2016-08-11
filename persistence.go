package main

import (
    "log"
    "time"
    "errors"
    // "database/sql"
    // _ "github.com/couchbase/go_n1ql"
    // couchbase "github.com/couchbase/go-couchbase"
    "gopkg.in/couchbase/gocb.v1"
)

type DBContext struct {
	Cluster *gocb.Cluster
	Bucket *gocb.Bucket
}

// TODO: use standard go err handling for all these, so that we can display errors from the api
// TODO: reconnect on error (easily test it by restarting couchbase during a run)

func dbConnect() (DBContext, error) {
	cluster, err := gocb.Connect("couchbase://ezra.heart")
	if err != nil {
		return DBContext{}, errors.New("unable to connect: " + err.Error())
	}
	bucket, err := cluster.OpenBucket("rtctp-store", "")
	if err != nil {
		return DBContext{}, errors.New("unable to open bucket: " + err.Error())
	}
	return DBContext{
		Cluster: cluster,
		Bucket: bucket,
	}, nil
}

// Wrapper around Candle which includes resolution and instrument, to support indexing and searching
type DBCandle struct {
	Type        string  `json:"type"`
	Candle      Candle  `json:"candle"`
	Instrument  string  `json:"instrument"`
	Resolution  string  `json:"resolution"`
}

func persistCandle(dbContext DBContext, candle Candle, instrument string, resolution string) bool {
	candle.Time = candle.Time.UTC()
	path := "candle:" + instrument + ":" + resolution + ":" + candle.Time.Format(time.RFC3339)
	dbCandle := DBCandle{
		Type: "candle",
		Candle: candle,
		Instrument: instrument,
		Resolution: resolution,
	}
	dbContext.Bucket.Upsert(path, dbCandle, 0)
	return true // TODO: return true or false depending on whether or not it worked
}

func getCandles(dbContext DBContext, instrumentId string, resolution string, startDate time.Time, endDate time.Time) Candles {
	
	query := gocb.NewN1qlQuery("SELECT * FROM `rtctp-store`")
	rows, err := dbContext.Bucket.ExecuteN1qlQuery(query, []interface{}{})
	if err != nil {
		// TODO: handle error
		log.Printf("Error running N1QL query: %s", err)
		return Candles{}
	}

	var row interface{}
	for rows.Next(&row) {
		log.Printf("Row: %v", row)
	}

	return Candles{

	}
}

