package main

import (
    "log"
    "net/http"
    "encoding/json"
    "os"
)

type Configuration struct {
    Port        string
    BaseURI     string
}

var config Configuration
var dbContext DBContext

func main() {

    dbContext = dbConnect()
    config = loadConfig()
    router := NewRouter()

    log.Printf("Starting server on 127.0.0.1:%s", config.Port)
    log.Fatal(http.ListenAndServe(":"+config.Port, router))
}

func loadConfig() Configuration {
    file, _ := os.Open("config.json")
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
        log.Fatal("error loading config: ", err)
    }
    return configuration
}