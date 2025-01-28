package model

import (
    "time"
)

type Link struct {
    ID        string    `json:"id"`
    LongURL   string    `json:"long_url"`
    ShortURL  string    `json:"short_url"`
    CreatedAt time.Time `json:"created_at"`
}