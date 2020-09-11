package lm

import (
	"regexp"

	"github.com/timurgarif/nlpgo"
)

// Rule defines an affix-related attributes to restore a lemma from a lexeme.
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
	// Minimal lexeme length used as a threshold trigger to apply the rule.
	// [optional] (ignored if zero value).
	MinValidLen int
	// A regexp to validate the lexeme before applying the affix detaching.
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

func (rr *ruleResolver) Resolve(lexeme string, acc LemmaAccumulator, max int) {
	lxmRuneLen := len([]rune(lexeme))
	lxmLen := len(lexeme)

	// If no rules are specified yield no lemma candidates
	if rr.rs == nil {
		return
	}

	// If no lemma cheker, consider the lexeme is lemma
	if rr.lc == nil {
		acc.Set(lexeme, nil)
		return
	}

	for _, r := range rr.rs {
		// Match the lexeme ending to suffix
		sfxStartIndex := lxmLen - len(r.Affix)
		if sfxStartIndex <= 0 ||
			r.Affix != lexeme[sfxStartIndex:] {
			continue
		}

		// Apply transforms until successful match or end
		for _, rt := range r.Transforms {
			c := rt.transform(lexeme, lxmRuneLen)

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
			// then a proper lexeme -> lemma match found
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

func (rt *RuleTransform) transform(lexeme string, lxmRuneLen int) string {
	if lxmRuneLen < rt.MinValidLen {
		return ""
	}
	detachIndex := lxmRuneLen - rt.Cutoff
	if detachIndex <= 0 {
		return ""
	}

	// Apply pre-op regexp
	if rt.ReBefore != nil && !rt.ReBefore.MatchString(lexeme) {
		return ""
	}

	c := string([]rune(lexeme)[:detachIndex]) + rt.Augment

	// Apply post-op regexp
	if rt.ReAfter != nil && !rt.ReAfter.MatchString(c) {
		return ""
	}

	return c
}
