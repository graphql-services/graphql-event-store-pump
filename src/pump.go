package src

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/golang/glog"

	"github.com/graphql-services/go-saga/eventstore"
	"github.com/graphql-services/go-saga/healthcheck"
)

// StartPumpWithHealthCheck ...
func StartPumpWithHealthCheck(aggregatorURL string) error {
	eventstore.OnEvent(eventstore.OnEventOptions{
		HandlerFunc: func(e eventstore.Event) error {
			data, err := json.Marshal(e)

			if err != nil {
				return err
			}

			glog.Infof("Sending Event %s", e.ID)
			res, err := http.Post(aggregatorURL, "application/json", bytes.NewReader(data))
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

	return healthcheck.StartHealthcheckServer()
}
