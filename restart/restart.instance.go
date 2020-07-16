package startinstance

import (
	"log"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

var (
	project  = os.Getenv("PROJECT")
	zone     = os.Getenv("ZONE")
	instance = os.Getenv("INSTANCE")
)

// PubSubMessage pub/sub message type
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// RestartInstance start instance
func RestartInstance(ctx context.Context, m PubSubMessage) error {
	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := computeService.Instances.Reset(project, zone, instance).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", resp)
	return nil
}
