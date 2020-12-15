package lm

import (
	"github.com/timurgarif/nlpgo"
)

type Lemma struct {
	// Lemma value
	Val string
	// Optional list of possible Parts Of Speech for the word parsed.
	Pos []nlpgo.POSId
}

// A map type to accumulate lemma candidates
type LemmaAccumulator map[string]map[nlpgo.POSId]struct{}

// LmChecker provides an interface to check if a lemma exist. What is considered
// to be lemma is implementation specific.
type LmChecker interface {
	// Lemma with the zero value shall be returned if `text` is not found
	lookup(text string) Lemma
}

// LmResolver is a way to apply a strategy to the
// lemmatization process
type LmResolver interface {
	// Resolves lemmata and adds them to acc. It should stop resolving if total
	// acc size >= max
	Resolve(word string, acc LemmaAccumulator, max int)
}

// Lemmatizer type implements lemmatization
type Lemmatizer struct {
	lc  LmChecker
	rs  []LmResolver
	acc LemmaAccumulator
}

// LmOption defines a functional option type for the Lemmatizer
type LmOption func(*Lemmatizer)

func NewLemmatizer(lkpr LmChecker, resolvers []LmResolver, opts ...LmOption) *Lemmatizer {
	return &Lemmatizer{lc: lkpr, rs: resolvers, acc: make(LemmaAccumulator)}
}

// Lemmatize returns the first resolved lemma
func (l Lemmatizer) Lemmatize(word string) Lemma {
	cc := l.LemmaCandidates(word, 1)
	if len(cc) > 0 {
		return cc[0]
	}

	return Lemma{}
}

// LemmaCandidates returns up to `max` lemma candidates for
// the given word.
// The order of the candidates in the returning array depends on:
//	- the order of the resolvers passed in to the Lemmatizer constructor
//	- the internal policy of each resolver
func (l Lemmatizer) LemmaCandidates(word string, max int) (candidates []Lemma) {
	if max <= 0 {
		max = 5
	}

	if word == "" {
		return
	}

	l.acc.clear()

	// First check if input is already a lemma
	if lm := l.lc.lookup(word); lm.Val != "" {
		l.acc.Set(lm.Val, lm.Pos)
	}

	if len(l.acc) >= max {
		candidates = l.acc.lemmata(max)
		return
	}

	// Apply resolvers
	for _, r := range l.rs {
		r.Resolve(word, l.acc, max)
		if len(l.acc) >= max {
			candidates = l.acc.lemmata(max)
			return
		}
	}

	candidates = l.acc.lemmata(max)
	return
}

func (acc LemmaAccumulator) clear() {
	// Optimized by the compiler since Go 1.11
	for k := range acc {
		delete(acc, k)
	}
}

func (acc LemmaAccumulator) lemmata(max int) (ll []Lemma) {
	i := 0

	for k, v := range acc {
		var pos []nlpgo.POSId
		for p := range v {
			pos = append(pos, p)
		}
		ll = append(ll, Lemma{Val: k, Pos: pos})

		i++
		if i >= max {
			return
		}
	}

	return
}

func (acc LemmaAccumulator) Set(lemma string, pp []nlpgo.POSId) {
	if _, ok := acc[lemma]; !ok {
		acc[lemma] = map[nlpgo.POSId]struct{}{}
	}

	for _, p := range pp {
		acc[lemma][p] = struct{}{}
	}
}
