package handlers

import (
	"github.com/labstack/echo"
	"github.com/twilio/twilio-go/twiml"
)

type AlertHandler struct {
}

func NewAlertHandler() *AlertHandler {
	return &AlertHandler{}
}

func (h *AlertHandler) MoistureHighAlertHandler(ctx echo.Context) error {
	level := ctx.QueryParam("level")
	say := twiml.VoiceSay{
		Message: "Soil moisture is high, please turn the motor off, the current soil moisture is " + level + ", Thank you for your time",
	}

	twimlResult, err := twiml.Voice([]twiml.Element{say})

	if err != nil {
		return ctx.String(500, err.Error())
	}

	ctx.Response().Header().Set(echo.HeaderContentType, "text/xml")

	return ctx.String(200, twimlResult)

}

func (h *AlertHandler) MoistureLowAlertHandler(ctx echo.Context) error {
	level := ctx.QueryParam("level")
	say := twiml.VoiceSay{
		Message: "Soil moisture is low, please turn the motor on, the current soil moisture is " + level + ", Thank you for your time",
	}

	twimlResult, err := twiml.Voice([]twiml.Element{say})

	if err != nil {
		return ctx.String(500, err.Error())
	}

	ctx.Response().Header().Set(echo.HeaderContentType, "text/xml")

	return ctx.String(200, twimlResult)
}
