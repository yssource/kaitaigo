// Autogenerated from KST: please remove this line if doing any edits by hand!

package spec

import (
	"os"
	"testing"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	. "test_formats"
	"github.com/stretchr/testify/assert"
)

func TestExprIoPos(t *testing.T) {
	f, err := os.Open("../../src/expr_io_pos.bin")
	if err != nil {
		t.Fatal(err)
	}
	s := kaitai.NewStream(f)
	var r ExprIoPos
	err = r.Read(s, &r, &r)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, "CURIOSITY", r.Substream1.MyStr)
	assert.EqualValues(t, []uint8{17, 34, 51, 68}, r.Substream1.Body)
	assert.EqualValues(t, 66, r.Substream1.Number)
	assert.EqualValues(t, "KILLED", r.Substream2.MyStr)
	assert.EqualValues(t, []uint8{97, 32, 99, 97, 116}, r.Substream2.Body)
	assert.EqualValues(t, 103, r.Substream2.Number)
}
