package Lib

import (
	"github.com/infobip/infobip-api-go-client/v2"
	"laundry/Config"
	"os"
)

func SendSMS(phoneNumber string, text string) {
	request := infobip.NewSmsAdvancedTextualRequest()

	destination := infobip.NewSmsDestination(phoneNumber)

	from := os.Getenv("SMS_FROM")
	message := infobip.NewSmsTextualMessage()
	message.From = &from
	message.Destinations = &[]infobip.SmsDestination{*destination}
	message.Text = &text

	request.Messages = &[]infobip.SmsTextualMessage{*message}

	_, _, err := Config.SmsMessageConfig().
		SmsAdvancedTextualRequest(*request).
		Execute()

	if err != nil {
		panic(err)
		apiErr, isApiErr := err.(infobip.GenericOpenAPIError)
		if isApiErr {
			ibErr, isIbErr := apiErr.Model().(infobip.SmsApiException)
			if isIbErr {
				ibErr.RequestError.ServiceException.GetMessageId()
				ibErr.RequestError.ServiceException.GetText()
			}
		}
	}
}
