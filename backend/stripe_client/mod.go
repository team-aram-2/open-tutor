package stripe_client

import (
	"fmt"
	"os"

	"github.com/stripe/stripe-go/v81/client"
)

var stripeClient *client.API

func GetClient() *client.API {
	if stripeClient == nil {
		stripeClient = &client.API{}
		stripeClient.Init(os.Getenv("STRIPE_SK"), nil)
		fmt.Printf("Stripe client initialized\n")
	}
	return stripeClient
}
