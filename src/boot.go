package src

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	"github.com/golang/glog"
	"github.com/graphql-services/go-saga/eventstore"
)

// PerformBootupOptions ..
type PerformBootupOptions struct {
	AggregatorURL string
}

// PerformBootup ...
func PerformBootup(options PerformBootupOptions) error {
	glog.Info("Initializing bootup")

	u, err := url.Parse(options.AggregatorURL)
	if err != nil {
		return err
	}

	u.Path = path.Join(u.Path, "events/latest")
	res, err := http.Get(u.String())
	if err != nil {
		return err
	}

	if res.StatusCode == 200 || res.StatusCode == 204 {
		glog.Info("Bootup not required")
		return nil
	}

	glog.Info("Booting up aggregator")

	if err := checkStatusCode(res, 404, "checking last event"); err != nil {
		return err
	}

	ctx := context.Background()

	glog.Info("Fetching all events started")

	ch := fetchAllEvents(ctx)

	for resp := range ch {
		err := forwardResponse(ctx, resp, options.AggregatorURL)
		if err != nil {
			return err
		}
	}

	return nil
}

func forwardResponse(ctx context.Context, res eventstore.FetchEventsResponse, aggregatorURL string) error {
	glog.Info("Forwarding events ", len(res.Events))
	importURL, err := url.Parse(aggregatorURL)
	if err != nil {
		return err
	}
	importURL.Path = path.Join(importURL.Path, "events/import")

	data, err := json.Marshal(res.Events)

	resp, err := http.Post(importURL.String(), "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}

	if err := checkStatusCode(resp, 201, "importing events"); err != nil {
		return err
	}

	glog.Info("Events forwarded ", resp.StatusCode)

	return nil
}

func fetchAllEvents(ctx context.Context) <-chan eventstore.FetchEventsResponse {
	var cursor *string
	ch := make(chan eventstore.FetchEventsResponse)

	go func() {
		for {
			glog.Info("Fetching events from cursor ", cursor)
			options := eventstore.FetchEventsOptions{Cursor: cursor}

			var data eventstore.FetchEventsResponse
			if err := eventstore.FetchEvents(ctx, options, &data); err != nil {
				panic(err)
			}

			ch <- data
			if len(data.Events) == 0 {
				close(ch)
				break
			}
			cursor = &data.Events[len(data.Events)-1].Cursor
		}
	}()

	return ch
}
