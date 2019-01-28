package src

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	"github.com/golang/glog"

	"github.com/graphql-services/go-saga/eventstore"
)

// StartPump ...
func StartPump(aggregatorURL string) error {

	u, err := url.Parse(aggregatorURL)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, "/events")

	eventstore.OnEvent(eventstore.OnEventOptions{
		Channel: "test",
		HandlerFunc: func(e eventstore.Event) error {
			data, err := json.Marshal(e)

			if err != nil {
				return err
			}

			glog.Infof("Sending Event %s", e.ID)
			res, err := http.Post(u.String(), "application/json", bytes.NewReader(data))
			if err != nil {
				return err
			}
			defer res.Body.Close()

			if err := checkStatusCode(res, 201, "when forwarding(importing) event"); err != nil {
				return err
			}

			glog.Infof("Event %s processed", e.ID)

			return nil
		},
	})

	return nil
}
