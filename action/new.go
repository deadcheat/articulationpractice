package action

import (
	"math/rand"
	"time"

	"github.com/deadcheat/alexa"
	"github.com/deadcheat/twister/globals"
	"github.com/deadcheat/twister/types"
	"github.com/deadcheat/twister/values"
	"github.com/rs/xid"
)

// New execute new trial
func New(req alexa.RequestEnvelope) (alexa.ResponseEnvelope, error) {
	guid := xid.New()
	i := rand.Intn(globals.TwsitersSize)
	q := values.Twisters[i]
	mi := req.Session.Attributes[values.SessionAttributeKeyMatch]
	sa := make(map[string]interface{})
	m := globals.ConvertInterfaceToMatch(mi)
	if m == nil {
		// TODO this pattern, error occured...
		m = &types.Match{
			MatchID:   guid.String(),
			Questions: make([]*types.Question, 0),
			StartedAt: time.Now().UnixNano(),
		}
	}
	if m.Questions == nil {
		m.Questions = make([]*types.Question, 0)
	}
	if m.Current != nil {
		m.Questions = append(m.Questions, m.Current)
	}
	m.Current = &types.Question{
		Text: q,
	}
	m.Total++

	sa[values.SessionAttributeKeyMatch] = *m
	return alexa.ResponseEnvelopeV1(
		sa,
		alexa.Response{
			OutputSpeech: &alexa.OutputSpeech{
				Type: alexa.TypePlainText,
				Text: q,
			},
			ShouldEndSession: false,
		}), nil
}
