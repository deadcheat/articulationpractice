package main

import (
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/deadcheat/alexa"
	"github.com/deadcheat/twister/action"
	"github.com/deadcheat/twister/globals"
	"github.com/deadcheat/twister/values"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	globals.TwsitersSize = len(values.Twisters)
	// create lambda handler
	h := alexa.NewLambdaHandler()

	// assign handler to request and intent
	h.HandleLaunch(action.Launch)
	h.HandleEnd(action.End)

	/// assign intent handlers
	h.HandleIntent([]string{alexa.IntentAMAZONMoreIntent, alexa.IntentAMAZONYesIntent, alexa.IntentAMAZONNextIntent}, action.New)

	h.HandleIntent([]string{alexa.IntentAMAZONStopIntent, alexa.IntentAMAZONNoIntent}, action.End)

	lambda.Start(h.Handle)
}
