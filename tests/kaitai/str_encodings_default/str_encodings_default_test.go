// Autogenerated from KST: please remove this line if doing any edits by hand!

package str_encodings_default

import (
	"os"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrEncodingsDefault(t *testing.T) {
	f, err := os.Open("../../../testdata/kaitai/str_encodings.bin")
	if err != nil {
		t.Fatal(err)
	}

	var r StrEncodingsDefault
	err = r.Decode(f)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, "Some ASCII", r.Str1())
	assert.EqualValues(t, "\u3053\u3093\u306b\u3061\u306f", r.Rest().Str2())
	assert.EqualValues(t, "\u3053\u3093\u306b\u3061\u306f", r.Rest().Str3())
	assert.EqualValues(t, "\u2591\u2592\u2593", r.Rest().Str4())
}
