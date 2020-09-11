package lm

import (
	"fmt"

	"github.com/timurgarif/nlpgo"
)

// POSIdNyble can be used to pack up to 4 PosId values (1 byte each) as a uint32
// (4 bytes) value.
// This allows to compact up to 4 POSId items (4 bytes) vs [4]POSId slice
// (24 + 4 bytes).
type POSIdNyble = uint32

// PackPosNyble is used to put slice of up to 4 POS elements to the PosIdNyble value.
// If len(ps) > 4 then only the first 4 elements will be packed.
func PackPosNyble(ps []nlpgo.POSId) POSIdNyble {
	var n POSIdNyble

	for i, v := range ps {
		if i > 3 {
			break
		}
		// shift v[i] by the 8 bit span
		n += POSIdNyble(uint32(v) << uint32(8*i))
		fmt.Println(n)
	}

	return n
}

func UnpackPosNyble(nyble POSIdNyble) []nlpgo.POSId {
	if nyble == 0 {
		return nil
	}

	ps := make([]nlpgo.POSId, 0, 4)

	for i := 0; i < 4; i++ {
		p := nlpgo.POSId(nyble >> uint32(8*i))
		if p == 0 {
			break
		}
		ps = append(ps, p)
	}

	return ps
}
