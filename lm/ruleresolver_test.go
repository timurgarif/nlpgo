package lm

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timurgarif/nlpgo"
)

func TestResolve(t *testing.T) {
	assert := assert.New(t)

	const (
		// Except w x y
		connosant = "[b-df-hj-np-tvz]"
		vowel     = "[aeiou]"
		// Golang regexp back referencing (like "[b]\1") not supported, so use explicit
		// Except w x y
		doubleConnosant = "(bb|cc|dd|ff|gg|hh|jj|kk|ll|mm|nn|pp|rr|ss|tt|vv|zz)"
	)

	rules := []Rule{
		{
			Affix: "ing",
			Pos:   []nlpgo.POSId{nlpgo.PosIdVbz},
			Transforms: []RuleTransform{
				{
					Cutoff:      3,
					Augment:     "e",
					MinValidLen: 5,
				},
				{
					Cutoff:      4,
					ReBefore:    regexp.MustCompile(`.` + vowel + doubleConnosant + `ing$`),
					MinValidLen: 6,
				},
				{
					Cutoff:      3,
					MinValidLen: 5,
				},
				{
					Cutoff:      4,
					ReBefore:    regexp.MustCompile(`.ying$`),
					Augment:     "ie",
					MinValidLen: 5,
				},
			},
		},
		{
			Affix: "л",
			Pos:   []nlpgo.POSId{nlpgo.PosIdVbd},
			Transforms: []RuleTransform{
				{
					Cutoff:  1,
					Augment: "ть",
				},
			},
		},
	}
	lmChecker := NewLemmaIndex(map[string][]nlpgo.POSId{
		"brainstorm":    {2, 4},
		"brainstorming": {2},
		"build":         {4, 2},
		"builder":       {2},
		"building":      {2},
		"take":          {2, 4},
		"string":        {4, 2, 3},
		"shoestring":    {2},
		"strip":         {4, 2},
		"stripe":        {2, 4},
		"tie":           {2, 4},
		"white-tie":     {3},
		"слушать":       {4},
	})

	emptyResolver := NewSuffixRuleResolver(nil, nil)
	noChkResolver := NewSuffixRuleResolver(rules, nil)
	resolver := NewSuffixRuleResolver(rules, lmChecker)

	cases := []struct {
		in  string
		out []Lemma
		msg string
		r   LmResolver
	}{
		{
			in:  "build",
			out: nil,
			msg: "Expect no lemma for resolver without any rules",
			r:   emptyResolver,
		},
		{
			in:  "building",
			out: []Lemma{{Val: "building"}},
			msg: "Expect input not modified for resolver without lemma checker",
			r:   noChkResolver,
		},
		{
			in:  "building",
			out: []Lemma{{Val: "build", Pos: []nlpgo.POSId{nlpgo.PosIdVbz}}},
			msg: "Expect building - ing -> build",
		},
		{
			in:  "taking",
			out: []Lemma{{Val: "take", Pos: []nlpgo.POSId{nlpgo.PosIdVbz}}},
			msg: "Expect taking - ing + e -> take",
		},
		{
			in:  "build",
			out: nil,
			msg: "Expect no rule is matched and applied",
		},
		{
			in:  "stripping",
			out: []Lemma{{Val: "strip", Pos: []nlpgo.POSId{nlpgo.PosIdVbz}}},
			msg: "Expect stripping -> strip",
		},
		{
			in:  "striping",
			out: []Lemma{{Val: "stripe", Pos: []nlpgo.POSId{nlpgo.PosIdVbz}}},
			msg: "Expect striping -> stripe",
		},
		{
			in:  "string",
			out: nil,
			msg: "Expect 'string' not matched by 'ing' rule",
		},
		{
			in:  "shoestring",
			out: nil,
			msg: "Expect 'shoestring' not matched by 'ing' rule",
		},
		{
			in:  "tying",
			out: []Lemma{{Val: "tie", Pos: []nlpgo.POSId{nlpgo.PosIdVbz}}},
			msg: "Expect tying -> tie",
		},
		{
			in:  "white-tying",
			out: nil,
			msg: "Expect white-tying not matched, because it has no VERB POS in lemma index",
		},
		{
			in:  "слушал",
			out: []Lemma{{Val: "слушать", Pos: []nlpgo.POSId{nlpgo.PosIdVbd}}},
			msg: "Expect rune words are handled correctly",
		},
	}

	for _, tt := range cases {
		const max = 10
		rs := tt.r

		// Default resolver if not specified in tt.r
		if rs == nil {
			rs = resolver
		}

		acc := make(LemmaAccumulator)
		rs.Resolve(tt.in, acc, max)
		assert.Equal(tt.out, acc.lemmata(max), tt.msg)
	}
}
