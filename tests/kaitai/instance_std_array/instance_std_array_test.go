// Autogenerated from KST: please remove this line if doing any edits by hand!

package instance_std_array

import (
	"os"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstanceStdArray(t *testing.T) {
	f, err := os.Open("../../../testdata/kaitai/instance_std_array.bin")
	if err != nil {
		t.Fatal(err)
	}

	var r InstanceStdArray
	err = r.Decode(f)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 16, r.Ofs())
	assert.EqualValues(t, 3, r.QtyEntries())
	assert.EqualValues(t, 4, r.EntrySize())
	assert.EqualValues(t, 3, len(r.Entries()))
	assert.EqualValues(t, []uint8{17, 17, 17, 17}, r.Entries()[0])
	assert.EqualValues(t, []uint8{34, 34, 34, 34}, r.Entries()[1])
	assert.EqualValues(t, []uint8{51, 51, 51, 51}, r.Entries()[2])
}
