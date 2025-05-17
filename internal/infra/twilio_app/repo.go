package twilio_app

import (
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type twilioRepo struct {
	twilioClient *twilioClient
}

func NewTwilioRepo(twilioClient *twilioClient) *twilioRepo {
	return &twilioRepo{
		twilioClient: twilioClient,
	}
}

func (repo *twilioRepo) MakeOrderCall(callAnswerApiUrl, callFrom, callTo string) error {
	params := new(api.CreateCallParams)
	params.SetUrl(callAnswerApiUrl)
	params.SetTo(callTo)
	params.SetFrom(callFrom)
	_, err := repo.twilioClient.client.Api.CreateCall(params)
	return err
}
