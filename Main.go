package main

import (
    "log"
    "net/http"
)

func main() {

    router := NewRouter()

    log.Fatal(http.ListenAndServe(":8042", router)) // TODO: move port into config
}