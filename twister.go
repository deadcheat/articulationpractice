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
	h.HandleIntent([]string{values.TwisterEndIntent, alexa.IntentAMAZONStopIntent, alexa.IntentAMAZONNoIntent, alexa.IntentAMAZONCancelIntent}, action.End)

	/// assign intent handlers
	h.HandleIntent([]string{values.TwisterContinueIntent, alexa.IntentAMAZONMoreIntent, alexa.IntentAMAZONYesIntent, alexa.IntentAMAZONNextIntent}, action.New)
	h.HandleIntent([]string{values.TwisterAnswerIntent}, action.Answer)
	h.HandleIntent([]string{alexa.IntentAMAZONHelpIntent}}, action.Help)

	lambda.Start(h.Handle)
}
