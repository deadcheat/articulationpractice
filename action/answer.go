package action

import (
	"fmt"

	"github.com/deadcheat/alexa"
	"github.com/deadcheat/twister/globals"
	"github.com/deadcheat/twister/values"
)

// Answer invoke when answer returns
func Answer(req alexa.RequestEnvelope) (alexa.ResponseEnvelope, error) {
	mi := req.Session.Attributes[values.SessionAttributeKeyMatch]
	m := globals.ConvertInterfaceToMatch(mi)
	if m == nil {
		return alexa.ResponseEnvelopeV1(
			alexa.EmptySessionAttributes,
			alexa.Response{
				OutputSpeech: &alexa.OutputSpeech{
					Type: alexa.TypePlainText,
					Text: "申し訳ありません。問題が発生しました。もう一度最初からやり直してください。",
				},
				ShouldEndSession: true,
			}), nil
	}
	message := "残念、正しく聞き取れませんでした。もう一度挑戦される場合は、はい、や、次、などのように話しかけてください。"
	if m.Current.Text == req.Intent.Slots[values.AnswerSlot].Value {
		m.Current.Success = true
		m.Score++
		message = `<say-as interpret-as="interjection">正解です、おめでとうございます</say-as><audio src='https://s3.amazonaws.com/ask-soundlibrary/human/amzn_sfx_crowd_applause_01.mp3'/>
		もう一度挑戦される場合は、はい、や、次、などのように話しかけてください。`
	}
	ssml := fmt.Sprintf(`<speak>%s</speak>`, message)
	sa := make(map[string]interface{})
	sa[values.SessionAttributeKeyMatch] = *m
	return alexa.ResponseEnvelopeV1(
		sa,
		alexa.Response{
			OutputSpeech: &alexa.OutputSpeech{
				Type: alexa.TypeSSML,
				SSML: ssml,
			},
			Reprompt: &alexa.Reprompt{
				OutputSpeech: &alexa.OutputSpeech{
					Type: alexa.TypePlainText,
					Text: "もう一度挑戦される場合は、はい、や、次、などのように話しかけてください。",
				},
			},
			ShouldEndSession: false,
		}), nil
}
