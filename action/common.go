package action

import (
	"time"

	"github.com/deadcheat/alexa"
	"github.com/deadcheat/twister/types"
	"github.com/deadcheat/twister/values"
	"github.com/rs/xid"
)

// Launch execute when alexa launch skills
func Launch(alexa.RequestEnvelope) (alexa.ResponseEnvelope, error) {
	sa := make(map[string]interface{})
	guid := xid.New()
	sa[values.SessionAttributeKeyMatch] = types.Match{
		MatchID:   guid.String(),
		Questions: make([]*types.Question, 0),
		StartedAt: time.Now().UnixNano(),
	}
	return alexa.ResponseEnvelopeV1(
		sa,
		alexa.Response{
			OutputSpeech: &alexa.OutputSpeech{
				Type: alexa.TypePlainText,
				Text: "早口言葉をはじめます。心と口の準備はよろしいですか？",
			},
			Reprompt: &alexa.Reprompt{
				OutputSpeech: &alexa.OutputSpeech{
					Type: alexa.TypePlainText,
					Text: "準備ができたら、OK、や、次、スタートのように話しかけてください。",
				},
			},
			ShouldEndSession: false,
		}), nil
}

// End execute when alexa end skills
func End(alexa.RequestEnvelope) (alexa.ResponseEnvelope, error) {
	return alexa.ResponseEnvelopeV1(
		alexa.EmptySessionAttributes,
		alexa.Response{
			OutputSpeech: &alexa.OutputSpeech{
				Type: alexa.TypePlainText,
				Text: "また練習したくなったら呼んでくださいね、さようなら",
			},
			ShouldEndSession: true,
		}), nil
}