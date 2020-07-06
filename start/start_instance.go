package startinstance

import (
	"context"
	"log"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func StartInstance(ctx context.Context, m PubSubMessage) error {
	name := string(m.Data)
	if name == "" {
		name = "Daigo"
	}
	log.Printf("Hello, %s!", name)
	return nil
}
