package startinstance

import (
	"log"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

type Message struct {
	Data []byte `json:"data"`
}

func StartInstance(ctx context.Context, m Message) error {
	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}
	project := os.Getenv("PROJECT")
	zone := os.Getenv("ZONE")
	instance := os.Getenv("INSTANCE")

	resp, err := computeService.Instances.Start(project, zone, instance).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", resp)
	return nil
}
