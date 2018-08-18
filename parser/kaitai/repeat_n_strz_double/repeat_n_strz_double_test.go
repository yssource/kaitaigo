// Autogenerated from KST: please remove this line if doing any edits by hand!

package repeat_n_strz_double

import (
	"os"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatNStrzDouble(t *testing.T) {
	f, err := os.Open("../../../testdata/kaitai/repeat_n_strz.bin")
	if err != nil {
		t.Fatal(err)
	}

	var r RepeatNStrzDouble
	err = r.Decode(f)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 2, r.Qty())
	assert.EqualValues(t, "foo", r.Lines1()[0])
	assert.EqualValues(t, "bar", r.Lines2()[0])
}
