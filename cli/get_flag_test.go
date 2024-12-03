package cli

import (
	"fmt"
	"os"
	"testing"
)

func TestGetFlag(t *testing.T) {
	args := GetFlag()
	fmt.Fprintln(os.Stdout, args)
}
