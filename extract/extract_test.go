package extract

import (
	"fmt"
	"os"
	"testing"

	"github.com/agnivade/levenshtein"
)

func TestCalcHammingDistance(t *testing.T) {
	fmt.Fprintln(os.Stdout, levenshtein.ComputeDistance("AGCT", "AGTC"))
}
