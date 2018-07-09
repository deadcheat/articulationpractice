package action

import (
	"fmt"
	"time"

	"github.com/deadcheat/alexa"
	"github.com/deadcheat/twister/globals"
	"github.com/deadcheat/twister/types"
	"github.com/deadcheat/twister/values"
	"github.com/rs/xid"
)

// Help execute when user call help
func Help(alexa.RequestEnvelope) (alexa.ResponseEnvelope, error) {
	return alexa.ResponseEnvelopeV1(
		alexa.EmptySessionAttributes,
		alexa.Response{
			OutputSpeech: &alexa.OutputSpeech{
				Type: alexa.TypePlainText,
				Text: "このスキルでは、早口言葉の練習をすることができます。準備ができたら、OK、や、次、スタートのように話しかけてください。早口言葉のお題を読み上げますので、読み上げたとおりに繰り返してください。正確に聞き取ることができたら、正解になります。それでは、早口言葉をはじめます。心と口の準備はよろしいですか？",
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
func End(req alexa.RequestEnvelope) (alexa.ResponseEnvelope, error) {
	mi := req.Session.Attributes[values.SessionAttributeKeyMatch]
	m := globals.ConvertInterfaceToMatch(mi)
	if m == nil {
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
	return alexa.ResponseEnvelopeV1(
		alexa.EmptySessionAttributes,
		alexa.Response{
			OutputSpeech: &alexa.OutputSpeech{
				Type: alexa.TypePlainText,
				Text: fmt.Sprintf("今回は%d回正解しました。また練習したくなったら呼んでくださいね、さようなら", m.Score),
			},
			ShouldEndSession: true,
		}), nil
}
