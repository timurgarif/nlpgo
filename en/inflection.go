package en

import (
	"regexp"

	"github.com/timurgarif/nlpgo"
	"github.com/timurgarif/nlpgo/lm"
)

// Quote from Wikipedia on `Y`:
// "In the English writing system, it sometimes represents a vowel and
// sometimes a consonant, and in other orthographies it may represent a vowel
// or a consonant."
// So it's present neither in the connosant const nor in the vowel const.
const (
	connosant = "[b-df-hj-np-tvwxz]"
	vowel     = "[aeiou]"
	// Golang regexp back referencing (like "[b]\1") not supported, so use explicit
	// Except w x
	doubleConnosantButWX = "(bb|cc|dd|ff|gg|hh|jj|kk|ll|mm|nn|pp|rr|ss|tt|vv|zz)"
)

var MorphRules = []lm.Rule{
	{
		Affix: "ing",
		Pos:   []nlpgo.POSId{nlpgo.PosIdVbg},
		Transforms: []lm.RuleTransform{
			{
				Cutoff:      3,
				Augment:     "e",
				MinValidLen: 5,
			},
			{
				Cutoff:      4,
				ReBefore:    regexp.MustCompile(`.` + vowel + doubleConnosantButWX + `ing$`),
				MinValidLen: 6,
			},
			{
				Cutoff:      4,
				ReBefore:    regexp.MustCompile(`.ying$`),
				Augment:     "ie",
				MinValidLen: 5,
			},
			// fallback
			{
				Cutoff:      3,
				MinValidLen: 5,
			},
		},
	},
	{
		Affix: "ed",
		Pos:   []nlpgo.POSId{nlpgo.PosIdVbn, nlpgo.PosIdVbd},
		Transforms: []lm.RuleTransform{
			{
				// faked -> fake
				Cutoff:   1,
				ReBefore: regexp.MustCompile(`.[^i]ed$`),
			},
			{
				//  played -> play
				Cutoff:   2,
				ReBefore: regexp.MustCompile(`.` + vowel + `yed$`),
			},
			{
				//  mimicked -> mimic
				Cutoff:   3,
				ReBefore: regexp.MustCompile(`..cked$`),
			},
			// tried -> try
			{
				Cutoff:   3,
				ReBefore: regexp.MustCompile(`.ied$`),
				Augment:  "y",
			},
			// zipped -> zip
			{
				Cutoff:      3,
				ReBefore:    regexp.MustCompile(vowel + doubleConnosantButWX + `ed$`),
				MinValidLen: 5,
			},
			// fallback
			{
				Cutoff:      2,
				MinValidLen: 4,
			},
		},
	},
	{
		Affix: "er",
		Pos:   []nlpgo.POSId{nlpgo.PosIdJjr},
		Transforms: []lm.RuleTransform{
			// easier -> easy
			{
				Cutoff:      3,
				ReBefore:    regexp.MustCompile(`ier$`),
				Augment:     "y",
				MinValidLen: 6,
			},
			// hotter -> hot
			{
				Cutoff:      3,
				ReBefore:    regexp.MustCompile(vowel + doubleConnosantButWX + `er$`),
				MinValidLen: 5,
			},
			// larger -> large
			{
				Cutoff:      1,
				MinValidLen: 4,
			},
			// smaller -> small
			{
				Cutoff:      2,
				MinValidLen: 5,
			},
		},
	},
	{
		Affix: "est",
		Pos:   []nlpgo.POSId{nlpgo.PosIdJjs},
		Transforms: []lm.RuleTransform{
			// easiest -> easy
			{
				Cutoff:      4,
				ReBefore:    regexp.MustCompile(`iest$`),
				Augment:     "y",
				MinValidLen: 7,
			},
			// hottest -> hot
			{
				Cutoff:      4,
				ReBefore:    regexp.MustCompile(vowel + doubleConnosantButWX + `est$`),
				MinValidLen: 6,
			},
			// largest -> large
			{
				Cutoff:      2,
				MinValidLen: 5,
			},
			// smallest -> small
			{
				Cutoff:      3,
				MinValidLen: 6,
			},
		},
	},
	{
		Affix: "s",
		Pos:   []nlpgo.POSId{nlpgo.PosIdNns, nlpgo.PosIdVbz},
		Transforms: []lm.RuleTransform{
			{
				Cutoff:   2,
				ReBefore: regexp.MustCompile(`.ches$`),
			},
			{
				Cutoff:   3,
				ReBefore: regexp.MustCompile(`.ves$`),
				Augment:  "f",
			},
			{
				Cutoff:   2,
				ReBefore: regexp.MustCompile(`.ses$`),
			},
			{
				Cutoff:   2,
				ReBefore: regexp.MustCompile(`.oes$`),
			},
			{
				Cutoff:   2,
				ReBefore: regexp.MustCompile(`.shes$`),
			},
			{
				Cutoff:   2,
				ReBefore: regexp.MustCompile(`.xes$`),
			},
			{
				Cutoff:   2,
				ReBefore: regexp.MustCompile(`.zes$`),
			},
			{
				Cutoff:   1,
				ReBefore: regexp.MustCompile(`.[^zs']s$`),
			},
			// countries -> country
			{
				Cutoff:   3,
				ReBefore: regexp.MustCompile(`.` + connosant + `ies$`),
				Augment:  "y",
			},
		},
	},
	{
		Affix: "men",
		Pos:   []nlpgo.POSId{nlpgo.PosIdNns},
		Transforms: []lm.RuleTransform{
			{
				Cutoff:      2,
				ReBefore:    regexp.MustCompile(`men$`),
				Augment:     "an",
				MinValidLen: 5,
			},
		},
	},
}
