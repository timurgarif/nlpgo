package lm

import "github.com/timurgarif/nlpgo"

// LemmaIndex is a default implementation of LmChecker
type LemmaIndex struct {
	idx map[string][]nlpgo.POSId
}

func NewLemmaIndex(data map[string][]nlpgo.POSId) LemmaIndex {
	return LemmaIndex{idx: data}
}

func (l LemmaIndex) lookup(text string) Lemma {
	if v, ok := l.idx[text]; ok {
		return Lemma{Val: text, Pos: v}
	}

	return Lemma{}
}
