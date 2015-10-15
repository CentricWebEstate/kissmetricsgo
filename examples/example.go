package main

import "github.com/centricwebestate/kissmetricsgo"
import "net/url"

func main() {
	// Initialise the Go Library
	km := kmgo.NewKM("yourapikeyhere")

	// Create a list of properties to send to kiss (can be empty)
	properties := url.Values{}
	properties.Set("test property", "true")

	// Finally create the event and send it
	event := km.NewEvent("testing go", "someguy@somesite.com", properties)
	event.Send()
}
