package nlpgo

type POS string

// A compact encoding of POS value
type POSId uint8

// The POSId values are specified so that to be compatible with the ones in my
// other projects
const (
	PosIdNoun POSId = 2
	PosIdAdj  POSId = 3
	PosIdVerb POSId = 4
	PosIdPron POSId = 5
	PosIdNum  POSId = 6
	PosIdAdv  POSId = 7
)

// https://universaldependencies.org/u/pos/
const (
	PosNoun POS = "NOUN"
	PosAdj  POS = "ADJ"
	PosVerb POS = "VERB"
	PosPron POS = "PRON"
	PosNum  POS = "NUM"
	PosAdv  POS = "ADV"
)
