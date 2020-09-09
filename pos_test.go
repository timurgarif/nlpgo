package nlpgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPosHasForm(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		obj  POSId
		subj POSId
		out  bool
		msg  string
	}{
		{
			msg:  "Adjective superlative is a form of Adjective",
			obj:  PosIdAdj,
			subj: PosIdJjs,
			out:  true,
		},
		{
			msg:  "Adjective is not a form of Adjective superlative",
			obj:  PosIdJjs,
			subj: PosIdAdj,
			out:  false,
		},
		{
			msg:  "Gerund is a form of Verb",
			obj:  PosIdVerb,
			subj: PosIdVbg,
			out:  true,
		},
		{
			msg:  "Plural Noun is not a form of Verb",
			obj:  PosIdVerb,
			subj: PosIdNns,
			out:  false,
		},
	}

	for _, t := range cases {
		assert.Equal(t.out, t.obj.HasForm(t.subj), t.msg)
	}
}
