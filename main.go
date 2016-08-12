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
    DBHost      string
    DBBucket    string
    DBPassword  string
}

var config Configuration
var dbContext DBContext

func main() {

    config = loadConfig("config.json")
    configOverrides := loadConfig("config-overrides.json")
    config = applyConfigOverrides(config, configOverrides)

    db, err := dbConnect(config.DBHost, config.DBBucket, config.DBPassword)
    dbContext = db
    if err != nil {
        log.Printf("error connecting to database: %s", err)
        return
    }

    router := NewRouter()

    log.Printf("Starting server on 127.0.0.1:%s", config.Port)
    log.Fatal(http.ListenAndServe(":"+config.Port, router))
}

func loadConfig(filename string) Configuration {
    file, _ := os.Open(filename)
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
        log.Fatal("error loading config: ", err)
    }
    return configuration
}

func applyConfigOverrides(config Configuration, configOverrides Configuration) Configuration {
    // TODO: use reflection
    if configOverrides.Port != "" { config.Port = configOverrides.Port }
    if configOverrides.BaseURI != "" { config.BaseURI = configOverrides.BaseURI }
    if configOverrides.DBHost != "" { config.DBHost = configOverrides.DBHost }
    if configOverrides.DBBucket != "" { config.DBBucket = configOverrides.DBBucket }
    if configOverrides.DBPassword != "" { config.DBPassword = configOverrides.DBPassword }
    return config
}