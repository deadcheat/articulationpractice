package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/deadcheat/alexa"
)

func main() {
	// create lambda handler
	h := alexa.NewLambdaHandler()

	// assign handler to request and intent
	h.HandleLaunch(func(alexa.RequestEnvelope) (alexa.ResponseEnvelope, error) {
		return alexa.EmptyResponse, alexa.ErrNoHandler
	})
	h.HandleEnd(func(alexa.RequestEnvelope) (alexa.ResponseEnvelope, error) {
		return alexa.EmptyResponse, alexa.ErrNoHandler
	})

	/// assign intent handlers
	h.HandleIntent("", func(alexa.RequestEnvelope) (alexa.ResponseEnvelope, error) {
		return alexa.EmptyResponse, alexa.ErrNoHandler
	})

	lambda.Start(h.Handle)
}
