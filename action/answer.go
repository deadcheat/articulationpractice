package action

import (
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
	sa := make(map[string]interface{})
	message := "残念、正しく聞き取れませんでした"
	if m.Current.Text == req.Intent.Slots[values.AnswerSlot].Value {
		m.Current.Success = true
		m.Score++
		message = "正解です、おめでとうございます"
	}
	sa[values.SessionAttributeKeyMatch] = *m
	return alexa.ResponseEnvelopeV1(
		sa,
		alexa.Response{
			OutputSpeech: &alexa.OutputSpeech{
				Type: alexa.TypePlainText,
				Text: message,
			},
			ShouldEndSession: false,
		}), nil
}
