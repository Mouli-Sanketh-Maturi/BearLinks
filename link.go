package main

type Link struct {
	ShortLink string `json:"shortLink"`
	LongLink  string `json:"longLink"`
	Enabled   bool `json:"enabled"`
}
