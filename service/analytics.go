package service

type Analytics struct {
	ShortLink string `json:"shortLink"`
	IP 	  string `json:"ip"`
	Timestamp 	  int64 `json:"timestamp"`
	Method string `json:"method"`
}