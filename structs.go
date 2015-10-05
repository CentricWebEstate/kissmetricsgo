package kmgo

import "net/url"

type KM struct {
	api_key string
}

type KMEvent struct {
	KM
	person     string // 255 Char Max
	name       string // Name of the event
	properties url.Values
	timestamp  int // Time of the event [optional]
}
