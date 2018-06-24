package main

import (
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/deadcheat/alexa"
	"github.com/deadcheat/twister/action"
	"github.com/deadcheat/twister/globals"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	globals.TwsitersSize = len(values.Twsiters)
	// create lambda handler
	h := alexa.NewLambdaHandler()

	// assign handler to request and intent
	h.HandleLaunch(action.Launch)
	h.HandleEnd(action.End)

	/// assign intent handlers
	h.HandleIntent("", func(alexa.RequestEnvelope) (alexa.ResponseEnvelope, error) {
		return alexa.EmptyResponse, alexa.ErrNoHandler
	})

	lambda.Start(h.Handle)
}
