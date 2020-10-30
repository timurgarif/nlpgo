// Package nlpgo provides basic NLP defs used in sub-packages.
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

	PosIdNns POSId = 30
	PosIdJjr POSId = 40
	PosIdJjs POSId = 41
	PosIdRbr POSId = 42
	PosIdRbs POSId = 43
	PosIdVbd POSId = 44
	PosIdVbn POSId = 45
	PosIdVbg POSId = 46
	PosIdVbp POSId = 47
	PosIdVbz POSId = 48
)

const (
	PosNoun POS = "NOUN"
	PosAdj  POS = "ADJ"
	PosVerb POS = "VERB"
	PosPron POS = "PRON"
	PosNum  POS = "NUM"
	PosAdv  POS = "ADV"

	PosNns POS = "NNS"
	PosJjr POS = "JJR"
	PosJjs POS = "JJS"
	PosRbr POS = "RBR"
	PosRbs POS = "RBS"
	PosVbd POS = "VBD"
	PosVbn POS = "VBN"
	PosVbg POS = "VBG"
	PosVbp POS = "VBP"
	PosVbz POS = "VBZ"
)

var posForms = map[POSId]POSId{
	PosIdNns: PosIdNoun,
	PosIdJjr: PosIdAdj,
	PosIdJjs: PosIdAdj,
	PosIdRbr: PosIdAdv,
	PosIdRbs: PosIdAdv,
	PosIdVbd: PosIdVerb,
	PosIdVbn: PosIdVerb,
	PosIdVbg: PosIdVerb,
	PosIdVbp: PosIdVerb,
	PosIdVbz: PosIdVerb,
}

func (p POSId) HasForm(f POSId) bool {
	if pos, ok := posForms[f]; ok {
		return pos == p
	}
	return false
}
