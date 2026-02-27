package messaging

import (
    "encoding/json"
    "log"
    "github.com/nats-io/nats.go"
)

type EventPublisher struct {
    nc *nats.Conn
}

func NewEventPublisher(nc *nats.Conn) *EventPublisher {
    return &EventPublisher{nc: nc}
}

func (ep *EventPublisher) Publish(subject string, event interface{}) error {
    data, err := json.Marshal(event)
    if err != nil {
        return err
    }

    if err := ep.nc.Publish(subject, data); err != nil {
        log.Printf("Error publishing event to subject %s: %v", subject, err)
        return err
    }

    log.Printf("Event published to subject %s: %s", subject, string(data))
    return nil
}

func (ep *EventPublisher) Subscribe(subject string, handler func(msg *nats.Msg)) (*nats.Subscription, error) {
    return ep.nc.Subscribe(subject, handler)
}