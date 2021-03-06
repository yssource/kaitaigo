// Autogenerated from KST: please remove this line if doing any edits by hand!

package position_abs

import (
	"os"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositionAbs(t *testing.T) {
	f, err := os.Open("../../../testdata/kaitai/position_abs.bin")
	if err != nil {
		t.Fatal(err)
	}

	var r PositionAbs
	err = r.Decode(f)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 32, r.IndexOffset())
	assert.EqualValues(t, "foo", r.Index().Entry())
}
