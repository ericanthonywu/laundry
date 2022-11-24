package Config

import (
	"context"
	"github.com/infobip/infobip-api-go-client/v2"
	"os"
)

func infoBipClient() (*infobip.APIClient, context.Context) {
	configuration := infobip.NewConfiguration()
	configuration.Host = os.Getenv("INFOBIP_BASEURL")
	return infobip.NewAPIClient(configuration),
		context.WithValue(context.Background(), infobip.ContextAPIKey, os.Getenv("INFOBIP_APIKEY"))
}

func SmsMessageConfig() infobip.ApiSendSmsMessageRequest {
	client, ctx := infoBipClient()
	return client.SendSmsApi.SendSmsMessage(ctx)
}
