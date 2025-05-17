package handlers

import (
	"github.com/labstack/echo"
	"github.com/twilio/twilio-go/twiml"
)

type OrderNotifyHandler struct {
}

func NewOrderNotifyHandler() *OrderNotifyHandler {
	return &OrderNotifyHandler{}
}

func (h *OrderNotifyHandler) OrderNotifyHandler(ctx echo.Context) error {
	say := twiml.VoiceSay{
		Message: "an order has been placed successfully please checkout your bag in the agriconnect app, thank you for your time",
	}

	twimlResult, err := twiml.Voice([]twiml.Element{say})

	if err != nil {
		return ctx.String(500, err.Error())
	}

	ctx.Response().Header().Set(echo.HeaderContentType, "text/xml")

	return ctx.String(200, twimlResult)

}
