# graphql-event-store-pump

[![Build Status](https://travis-ci.org/graphql-services/graphql-event-store-pump.svg?branch=master)](https://travis-ci.org/graphql-services/graphql-event-store-pump)

Data pump for services aggregating EventStore events.

This service can be used for 2 things:

1. initial bootstrap of aggregator by loading all events from EventStore and sending them to aggregator
1. continual updates of aggregator by consuming events from NSQ and forwarding them to aggregator

NOTE: batch importing should be sequential, so this worker should not be running more than once for each aggregator.

## Diagram

![EventStore Pump diagram](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/graphql-services/graphql-event-store-pump/master/resources/diagram.puml?v1 'EvemtStore Pump diagram')

# Aggregator requirements

Aggregator should have these 3 methods:

```
GET /events/latest
- used for initial check if aggregator needs bootup
- it's aggregator's responsibility to store latest event
- status code 200 with event payload: no bootup needed, skipping
- status code 204: no bootup needed, skipping
- status code 404: bootup required, proceed to

POST /events/batch
- endpoint for forwarding array of events, response should be 201 to indicate success
- payload (list of events) is sorted and should be applied in same order
- it is expected that last event is available in GET /events/latest call after success
- batch import is sequential (if scaled to 1 running instance)
- last batch is indicated by empty array and aggregator should switch to "booted" state

POST /events
- endpoint for forwarding single event. Response code should be 201 to indicate success
- it is expected that this payload is available in GET /events/latest call after success
- this method is used only in continual updates

```

# Healtchecks

It's good practise to provide service healthcheck endpoint. Aggregator should not return successful healthcheck until fully booted up (be aware of readiness probes timeout in kubernetes)
