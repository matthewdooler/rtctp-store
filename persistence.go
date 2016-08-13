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

func dbConnect(host string, bucketName string, password string) (DBContext, error) {
	log.Printf("Connecting to database: %s@%s", bucketName, host)
	cluster, err := gocb.Connect(host)
	if err != nil {
		return DBContext{}, errors.New("unable to connect: " + err.Error())
	}
	bucket, err := cluster.OpenBucket(bucketName, password)
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

// TODO: test me
func getPath(instrument string, resolution string, candleTime time.Time) string {
	return "candle:" + instrument + ":" + resolution + ":" + candleTime.Format(time.RFC3339)
}

// TODO: can we bulk write?
func persistCandle(dbContext DBContext, candle Candle, instrument string, resolution string) error {
	candle.Time = candle.Time.UTC()
	path := getPath(instrument, resolution, candle.Time)
	dbCandle := DBCandle{
		Type: "candle",
		Candle: candle,
		Instrument: instrument,
		Resolution: resolution,
	}
	_, err := dbContext.Bucket.Upsert(path, dbCandle, 0)
	return err
}

func getCandles(dbContext DBContext, instrument string, resolution string, resolutionDuration time.Duration, startDate time.Time, endDate time.Time) (Candles, error) {
	startDate = startDate.UTC()
	endDate = endDate.UTC()

	// TODO: are key-value loookups faster?

	// query := gocb.NewN1qlQuery("SELECT * FROM `rtctp-store`")
	// rows, err := dbContext.Bucket.ExecuteN1qlQuery(query, []interface{}{})
	// if err != nil {
	// 	return Candles{}, errors.New("unable to run query: " + err.Error())
	// }

	// var row interface{}
	// for rows.Next(&row) {
	// 	log.Printf("Row: %v", row)
	// }

	var candles = Candles{}

	var candleTime time.Time = startDate
	var path string
	for !candleTime.After(endDate) {
		path = getPath(instrument, resolution, candleTime)
		candleTime = candleTime.Add(resolutionDuration)

		var candle Candle
		_, err := dbContext.Bucket.Get(path, &candle)
		if(err != nil) {
			errString := err.Error()
			if errString != "Key not found." {
				return candles, errors.New("unable to retrieve candle: " + errString)
			}
		} else {
			candles = append(candles, candle)
		}
	}

	return candles, nil
}

