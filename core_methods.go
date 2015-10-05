package kmgo

import "net/http"
import "net/url"

func NewKM(key string) *KM {
	return &KM{api_key: key}
}

func (km KM) NewEvent(event string, person string, properties url.Values) *KMEvent {
	event_struct := KMEvent{km, person, event, properties, 0}
	return &event_struct
}

func (ev *KMEvent) SetTimestamp(time int) *KMEvent {
	ev.timestamp = time
	return ev
}

func (ev *KMEvent) Send() (*KMEvent, error) {
	address := "https://trk.kissmetrics.com/e?"
	ev.properties.Add("_k", ev.api_key)
	ev.properties.Add("_p", ev.person)
	ev.properties.Add("_n", ev.name)

	// If we're specifying a custom timestamp
	if ev.timestamp != 0 {
		ev.properties.Add("_t", string(ev.timestamp))
		ev.properties.Add("_d", "1")
	}

	url_params := ev.properties.Encode()

	address += url_params

	request, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return ev, nil
}
