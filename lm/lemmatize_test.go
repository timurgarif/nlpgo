package lm

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timurgarif/nlpgo"
)

// func BenchmarkMap100(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		foo(i)
// 	}
// }

type testCases = []struct {
	in      string
	lzOut   Lemma
	lzCcOut []Lemma
	l       *Lemmatizer
}

func lsort(l Lemma) Lemma {
	sort.Slice(l.Pos, func(i, j int) bool {
		return l.Pos[i] < l.Pos[j]
	})
	return l
}

func lmsort(l []Lemma) []Lemma {
	sort.Slice(l, func(i, j int) bool {
		return l[i].Val < l[j].Val
	})

	for _, lm := range l {
		sort.Slice(lm.Pos, func(i, j int) bool {
			return lm.Pos[i] < lm.Pos[j]
		})
	}

	return l
}

func getCases() testCases {
	emptyLemmatizer := NewLemmatizer(NewLemmaIndex(nil), nil)

	lmIdx := make(map[string][]nlpgo.POSId)
	var pos1 []nlpgo.POSId
	pos2 := []nlpgo.POSId{nlpgo.PosIdVerb, nlpgo.PosIdNoun}
	lmIdx["some"] = pos1
	lmIdx["another"] = nil
	lmIdx["pass"] = pos2
	withLmidxLemmatizer := NewLemmatizer(NewLemmaIndex(lmIdx), nil)

	excpIdx := make(map[string][]Lemma)
	excprlvr := NewExceptionResolver(excpIdx)

	lm1 := []Lemma{{Val: "mouse", Pos: []nlpgo.POSId{nlpgo.PosIdNns}}}
	lm2 := []Lemma{
		{Val: "leaf", Pos: []nlpgo.POSId{nlpgo.PosIdNns}},
		{Val: "leave", Pos: []nlpgo.POSId{nlpgo.PosIdNns, nlpgo.PosIdVbz}},
	}
	excpIdx["mice"] = lm1
	excpIdx["leaves"] = lm2
	withLmidxExcpidxLemmatizer := NewLemmatizer(NewLemmaIndex(lmIdx), []LmResolver{excprlvr})

	cases := testCases{
		{
			in:    "",
			lzOut: Lemma{},
			l:     emptyLemmatizer,
		},
		{
			in:    "sample",
			lzOut: Lemma{},
			l:     emptyLemmatizer,
		},
		{
			in:    "sample",
			lzOut: Lemma{},
			l:     emptyLemmatizer,
		},
		{
			in:    "sample",
			lzOut: Lemma{},
			l:     withLmidxLemmatizer,
		},
		{
			in:      "some",
			lzOut:   Lemma{Val: "some", Pos: pos1},
			lzCcOut: []Lemma{{Val: "some", Pos: pos1}},
			l:       withLmidxLemmatizer,
		},
		{
			in:      "another",
			lzOut:   Lemma{Val: "another"},
			lzCcOut: []Lemma{{Val: "another"}},
			l:       withLmidxLemmatizer,
		},
		{
			in:      "pass",
			lzOut:   Lemma{Val: "pass", Pos: pos2},
			lzCcOut: []Lemma{{Val: "pass", Pos: pos2}},
			l:       withLmidxLemmatizer,
		},
		{
			in:    "",
			lzOut: Lemma{},
			l:     withLmidxLemmatizer,
		},
		{
			in:    "not existing text",
			lzOut: Lemma{},
			l:     withLmidxLemmatizer,
		},
		{
			in:      "mice",
			lzOut:   lm1[0],
			lzCcOut: []Lemma{lm1[0]},
			l:       withLmidxExcpidxLemmatizer,
		},
		{
			in:      "leaves",
			lzOut:   lm2[0],
			lzCcOut: []Lemma{lm2[0], lm2[1]},
			l:       withLmidxExcpidxLemmatizer,
		},
	}

	return cases
}

func TestLemmatize(t *testing.T) {
	assert := assert.New(t)

	cases := getCases()

	for _, tt := range cases {
		actual := lsort(tt.l.Lemmatize(tt.in))
		expected := lsort(tt.lzOut)
		assert.Equal(expected, actual)
	}
}

func TestLemmaCandidates(t *testing.T) {
	assert := assert.New(t)

	cases := getCases()

	for _, tt := range cases {
		actual := lmsort(tt.l.LemmaCandidates(tt.in, 10))
		expected := lmsort(tt.lzCcOut)
		assert.Equal(expected, actual)
	}
}
