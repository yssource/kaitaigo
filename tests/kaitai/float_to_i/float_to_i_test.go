// Autogenerated from KST: please remove this line if doing any edits by hand!

package float_to_i

import (
	"os"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloatToI(t *testing.T) {
	f, err := os.Open("../../../testdata/kaitai/floating_points.bin")
	if err != nil {
		t.Fatal(err)
	}

	var r FloatToI
	err = r.Decode(f)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 0.5, r.SingleValue())
	assert.EqualValues(t, 0.25, r.DoubleValue())
	assert.EqualValues(t, 0, r.SingleI())
	assert.EqualValues(t, 0, r.DoubleI())
	assert.EqualValues(t, 1, r.Float1I())
	assert.EqualValues(t, 1, r.Float2I())
	assert.EqualValues(t, 1, r.Float3I())
	assert.EqualValues(t, -2, r.Float4I())
}