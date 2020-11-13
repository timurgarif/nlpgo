// +build !dataset

package en

import (
	"github.com/timurgarif/nlpgo"
	"github.com/timurgarif/nlpgo/lm"
)

var (
	ExceptionsIdx = map[string][]lm.Lemma{}
	LemmaIdx      = map[string][]nlpgo.POSId{}
)
