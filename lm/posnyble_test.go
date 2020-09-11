package lm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/timurgarif/nlpgo"
)

type assertOp func(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool

func TestPosNyble(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		in  []POSId
		out []POSId
		op  assertOp
	}{
		{
			in:  []POSId{PosIdNoun, PosIdVerb},
			out: []POSId{PosIdNoun, PosIdVerb},
			op:  assert.Equal,
		},
		{
			in:  []POSId{PosIdNoun, PosIdVerb, PosIdAdj, PosIdPron, PosIdAdv},
			out: []POSId{PosIdNoun, PosIdVerb, PosIdAdj, PosIdPron},
			op:  assert.Equal,
		},
		{
			in:  []POSId{PosIdNoun},
			out: []POSId{PosIdNoun},
			op:  assert.Equal,
		},
		{
			in:  []POSId{PosIdNoun, PosIdVerb},
			out: []POSId{PosIdVerb, PosIdNoun},
			op:  assert.NotEqual,
		},
		{
			in:  []POSId{},
			out: nil,
			op:  assert.Equal,
		},
		{
			in:  nil,
			out: nil,
			op:  assert.Equal,
		},
	}

	for _, tt := range cases {
		tt.op(tt.out, UnpackPosNyble(PackPosNyble(tt.in)))
	}
}
