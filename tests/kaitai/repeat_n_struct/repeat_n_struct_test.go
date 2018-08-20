// Autogenerated from KST: please remove this line if doing any edits by hand!

package repeat_n_struct

import (
	"os"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatNStruct(t *testing.T) {
	f, err := os.Open("../../../testdata/kaitai/repeat_n_struct.bin")
	if err != nil {
		t.Fatal(err)
	}

	var r RepeatNStruct
	err = r.Decode(f)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 2, len(r.Chunks()))
	assert.EqualValues(t, 16, r.Chunks()[0].Offset())
	assert.EqualValues(t, 8312, r.Chunks()[0].Len())
	assert.EqualValues(t, 8328, r.Chunks()[1].Offset())
	assert.EqualValues(t, 15, r.Chunks()[1].Len())
}