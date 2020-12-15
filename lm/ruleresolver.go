package lm

import (
	"regexp"

	"github.com/timurgarif/nlpgo"
)

// Rule defines an affix-related attributes to restore a lemma from a word.
// Currently only suffix-based rules are supported.
type Rule struct {
	Affix string
	// Tag(s) to identify the inflection form
	Pos        []nlpgo.POSId
	Transforms []RuleTransform
}

type RuleTransform struct {
	// Number of chars to discard
	Cutoff int
	// A compensative affix to augment to a word after Affix detaching
	// [optional]
	Augment string
	// Minimal word length used as a threshold trigger to apply the rule.
	// [optional] (ignored if zero value).
	MinValidLen int
	// A regexp to validate the word before applying the affix detaching.
	// [optional]
	ReBefore *regexp.Regexp
	// A regexp to validate the lemma candidate after Affix detaching
	// [optional]
	ReAfter *regexp.Regexp
}

type ruleResolver struct {
	rs []Rule
	lc LmChecker
}

func NewSuffixRuleResolver(rules []Rule, lc LmChecker) LmResolver {
	return &ruleResolver{rs: rules, lc: lc}
}

func (rr *ruleResolver) Resolve(word string, acc LemmaAccumulator, max int) {
	wdRuneLen := len([]rune(word))
	wdLen := len(word)

	// If no rules are specified yield no lemma candidates
	if rr.rs == nil {
		return
	}

	// If no lemma cheker, consider the word is lemma
	if rr.lc == nil {
		acc.Set(word, nil)
		return
	}

	for _, r := range rr.rs {
		// Match the word ending to suffix
		sfxStartIndex := wdLen - len(r.Affix)
		if sfxStartIndex <= 0 ||
			r.Affix != word[sfxStartIndex:] {
			continue
		}

		// Apply transforms until successful match or end
		for _, rt := range r.Transforms {
			c := rt.transform(word, wdRuneLen)

			if c == "" {
				continue
			}

			l := rr.lc.lookup(c)
			if l.Val == "" {
				continue
			}

			// Match lemma candidate POS'es to the rule form POS'es.
			var pp []nlpgo.POSId
			for _, lemmaPos := range l.Pos {
				for _, rulePos := range r.Pos {
					if lemmaPos.HasForm(rulePos) {
						pp = append(pp, rulePos)
					}
				}
			}
			// If any rule POS forms correspond to the cheker lemma POS'es
			// then a proper word -> lemma match found
			if len(pp) > 0 {
				acc.Set(l.Val, pp)

				return
				// TODO: Currently we stop after the first rule match
				// 	Consider to parameterize the logic to allow all the rules to be applied
				// in order to look for more matches.
			}
		}
	}
}

func (rt *RuleTransform) transform(word string, wdRuneLen int) string {
	if wdRuneLen < rt.MinValidLen {
		return ""
	}
	detachIndex := wdRuneLen - rt.Cutoff
	if detachIndex <= 0 {
		return ""
	}

	// Apply pre-op regexp
	if rt.ReBefore != nil && !rt.ReBefore.MatchString(word) {
		return ""
	}

	c := string([]rune(word)[:detachIndex]) + rt.Augment

	// Apply post-op regexp
	if rt.ReAfter != nil && !rt.ReAfter.MatchString(c) {
		return ""
	}

	return c
}
