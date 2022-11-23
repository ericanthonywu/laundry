package Config

import (
	"context"
	"github.com/infobip/infobip-api-go-client/v2"
	"os"
)

func infoBipClient() infobip.ApiSendSmsMessageRequest {
	configuration := infobip.NewConfiguration()
	configuration.Host = os.Getenv("INFOBIP_BASEURL")
	auth := context.WithValue(context.Background(), infobip.ContextAPIKey, os.Getenv("INFOBIP_APIKEY"))
	return infobip.NewAPIClient(configuration).SendSmsApi.SendSmsMessage(auth)
}

func SendSMS(phoneNumber string, text string) {
	request := infobip.NewSmsAdvancedTextualRequest()

	destination := infobip.NewSmsDestination(phoneNumber)

	from := "Laundry Tech"
	message := infobip.NewSmsTextualMessage()
	message.From = &from
	message.Destinations = &[]infobip.SmsDestination{*destination}
	message.Text = &text

	request.Messages = &[]infobip.SmsTextualMessage{*message}

	_, _, err := infoBipClient().
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
