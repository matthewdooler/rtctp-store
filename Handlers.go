package main

import (
    "fmt"
    "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "rtctp store") // TODO: links to collections etc
}
