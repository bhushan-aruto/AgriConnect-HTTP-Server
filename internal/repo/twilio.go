package repo

type TwilioRepo interface {
	MakeOrderCall(callAnswerApiUrl, callFrom, callTo string) error
}
