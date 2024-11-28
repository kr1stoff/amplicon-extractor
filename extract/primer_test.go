package extract

import (
	"fmt"
	"os"
	"testing"
)

func TestExpandDegenerateBases(t *testing.T) {
	var primer string
	// ? primer := "CARGACATNGTTYAGTGGATGAG"
	expandedPrimers := ExpandDegenerateBases(primer)

	for _, seq := range expandedPrimers {
		fmt.Fprintln(os.Stdout, seq)
	}
}
