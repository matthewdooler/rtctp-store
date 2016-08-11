package main

import (
    "log"
    "time"
    // "database/sql"
    // _ "github.com/couchbase/go_n1ql"
    // couchbase "github.com/couchbase/go-couchbase"
    "gopkg.in/couchbase/gocb.v1"
)

func testdb() {
	// n1ql, err := sql.Open("n1ql", "http://ezra.heart:8093") // TODO: use config
	// if err != nil {
	//     log.Printf("Unable to connect to database: %s", err)
	// }

	// rows, err := n1ql.Query("select * from rtctp-store")
	// if err != nil {
	//     log.Printf("Unable to run query: %s", err)
	// } else {
	// 	log.Printf("Got some rows: %s", rows)
	// }

	// defer rows.Close()
	// for rows.Next() {
	//     var contacts string
	//     if err := rows.Scan(&contacts); err != nil {
	//         log.Printf("Unable to scan rows: %s", err)
	//     }
	//     log.Printf("Row returned %s : \n", contacts)
	// }




	// c, err := couchbase.Connect("http://ezra.heart:8091/")
	// if err != nil {
	//     log.Printf("Error connecting:  %v", err)
	//     return
	// }

	// pool, err := c.GetPool("default")
	// if err != nil {
	//     log.Printf("Error getting pool:  %v", err)
	//     return
	// }

	// bucket, err := pool.GetBucket("rtctp-store")
	// if err != nil {
	//     log.Printf("Error getting bucket:  %v", err)
	// } else {
	// 	log.Printf("Got bucket: %s", bucket)
	// }

	// TODO: http://developer.couchbase.com/documentation/server/4.5/sdk/go/start-using-sdk.html
	cluster, _ := gocb.Connect("couchbase://ezra.heart")
	bucket, _ := cluster.OpenBucket("rtctp-store", "")
	log.Printf("Got bucket: %s", bucket)


	// query := gocb.NewN1qlQuery("SELECT * FROM rtctp-store")
	// rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{})
	// var row interface{}
	// for rows.Next(&row) {
	// 	log.Printf("Row: %v", row)
	// }

}

type DBContext struct {
	Cluster *gocb.Cluster
	Bucket *gocb.Bucket
}

func dbConnect() DBContext {
	// TODO: error handling! :O
	cluster, _ := gocb.Connect("couchbase://ezra.heart")
	bucket, _ := cluster.OpenBucket("rtctp-store", "")
	return DBContext{
		Cluster: cluster,
		Bucket: bucket,
	}
}

// Wrapper around Candle which includes resolution and instrument, to support indexing and searching
type DBCandle struct {
	Type          string  `json:"type"`
	Candle        Candle  `json:"candle"`
	Instrument  string  `json:"instrument"`
	Resolution    string  `json:"resolution"`
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


