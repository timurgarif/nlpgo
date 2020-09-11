package en

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timurgarif/nlpgo"
	"github.com/timurgarif/nlpgo/lm"
)

var testLemmaIdx = map[string][]nlpgo.POSId{
	"walk":      {4, 2},
	"call":      {4, 2},
	"close":     {3, 2, 4, 7},
	"free":      {2, 7, 4, 3},
	"play":      {2, 4},
	"work":      {4, 2},
	"overdress": {4},
	"stare":     {4, 2},
	"star":      {4, 3, 2},
	"easy":      {3, 7},
	"fast":      {2, 3, 7, 4},
	"hot":       {3},
	"fine":      {7, 3, 2, 4},
	"wise":      {2, 3},
	"wolf":      {2, 4},
	"jockey":    {4, 2},
	"class":     {2, 4},
	"classify":  {4},
	"potato":    {2},
	"woman":     {2},
}

var testSet = [][]string{
	{"walked", "walk"},
	{"called", "call"},
	{"calling", "call"},
	{"closed", "close"},
	{"closing", "close"},
	{"freed", "free"},
	{"freeing", "free"},
	{"played", "play"},
	{"playing", "play"},
	{"working", "work"},
	{"overdressed", "overdress"},
	{"staring", "stare"},
	{"starring", "star"},
	{"stared", "stare"},
	{"starred", "star"},
	{"calls", "call"},
	{"easier", "easy"},
	{"faster", "fast"},
	{"easiest", "easy"},
	{"fastest", "fast"},
	{"hotter", "hot"},
	{"hottest", "hot"},
	{"finest", "fine"},
	{"finer", "fine"},
	{"classifies", "classify"},
	{"classes", "class"},
	{"potatoes", "potato"},
	{"wolves", "wolf"},
	{"women", "woman"},
	// {"wiseish", "wise"},
	// {"wolfish", "wolf"},
	// {"jokeyish", "jockey"},
}

func TestMorphRules(t *testing.T) {
	assert := assert.New(t)

	lmChecker := lm.NewLemmaIndex(testLemmaIdx)

	lzr := lm.NewLemmatizer(lmChecker,
		[]lm.LmResolver{lm.NewSuffixRuleResolver(MorphRules, lmChecker)})

	for _, v := range testSet {
		ll := lzr.Lemmatize(v[0])
		assert.NotEmpty(ll)
		assert.Equal(v[1], ll.Val)
	}
}
