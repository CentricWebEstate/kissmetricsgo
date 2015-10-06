# Kissmetrics Go
Kissmetrics Go is a KissMetrics library written in Go to help simplify the task of sending events to Kiss from your go project.

## Get Started
It's really simple to get started, first we need to make sure that we have downloaded the library:

```sh
$ go get github.com/centricwebestate/kissmetricsgo
```
Now that we've got the library we can initialise the `KM` type and start sending events. 

**Note:** The library exposes it's types and functions under the `kmgo` package.

```go
package main

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
```

That's it, now just run the script and you'll see your event in kissmetrics live view :)

You'll find this script in example.go just in case you need it.

## Licence
This library is Copyright 2015 Centric Web Estate and licenced under MIT.