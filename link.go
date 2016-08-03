package main

type Link struct {
    Rel      string    `json:"rel"`
    Href     string    `json:"href"`
}

type Links []Link