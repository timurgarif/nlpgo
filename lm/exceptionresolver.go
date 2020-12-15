package lm

type exceptResolver struct {
	excidx map[string][]Lemma
}

func NewExceptionResolver(excidx map[string][]Lemma) LmResolver {
	return exceptResolver{excidx: excidx}
}

func (r exceptResolver) Resolve(word string, acc LemmaAccumulator, max int) {
	if lemmata, ok := r.excidx[word]; ok {
		for _, lm := range lemmata {
			acc.Set(lm.Val, lm.Pos)
			if len(acc) >= max {
				return
			}
		}
	}
}
