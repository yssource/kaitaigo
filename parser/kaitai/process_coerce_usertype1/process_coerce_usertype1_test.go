// Autogenerated from KST: please remove this line if doing any edits by hand!

package spec

import (
	"os"
	"testing"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	. "test_formats"
	"github.com/stretchr/testify/assert"
)

func TestProcessCoerceUsertype1(t *testing.T) {
	f, err := os.Open("../../src/process_coerce_bytes.bin")
	if err != nil {
		t.Fatal(err)
	}
	s := kaitai.NewStream(f)
	var r ProcessCoerceUsertype1
	err = r.Read(s, &r, &r)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 0, r.Records[0].Flag)
	tmp1, err := r.Records[0].Buf()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, 1094795585, tmp1.Value)
	assert.EqualValues(t, 1, r.Records[1].Flag)
	tmp2, err := r.Records[1].Buf()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, 1111638594, tmp2.Value)
}