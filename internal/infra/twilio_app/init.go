package twilio_app

import (
	"github.com/twilio/twilio-go"
)

type twilioClient struct {
	client *twilio.RestClient
}

func NewTwilioClient() *twilioClient {
	client := twilio.NewRestClient()

	return &twilioClient{
		client: client,
	}
}
