# kaitai.go

kaitai.go is a compiler and runtime to create Go parsers from [Kaitai Struct](http://kaitai.io/) files.

## Installation
```sh
go get gitlab.com/cugu/kaitai.go
```

## Usage

### 1. Create ksy file

**my_format.ksy**
```yaml
meta:
  id: my_format
seq:
  - id: data_size
    type: u1
  - id: data
    size: data_size 
```

### 2. Create go parser

```sh
kaitai.go test.ksy # creates my_format.ksy.go
```

### 3. Use parser

**main.go**
```go
package main

import (
	"bytes"
	"log"
)

func main() {
    f := bytes.NewReader([]byte("\x05Hello world!"))
    var r MyFormat
    err := r.Decode(f) // Decode takes any io.ReadSeeker
    if err != nil {
        log.Fatal(err)
    }
    log.Print(string(r.Data())) // Prints "Hello"
}
```
## kaitai.go features

### Supported kaitai features:

 - Type specification
    - meta
        - endianess*
    - doc
    - seq
    - instances
    - enums
 - Attribute specification
    - id
    - doc
    - contents
    - repeat, repeat-expr, repeat-until
    - if
    - size, size-eos
    - process
    - terminator
    - consume
    - include
    - pad
    - eos-error
 - Primitive data types
 - Processing specification
    - xor
    - rol
    - ror
    - zlib
 - Instance specification
    - pos
    - value

_*partially_

### Additional features:

#### whence

Can be used togher with `pos` the define the reference point of the position. Valid values are `seek_set`, `seek_end` and `seek_cur` (default).

### Limitations

 - No _io (Most uses can be replaced with [whence](#whence))
 - Accessing nested types with `::` is not allowed
 - No fancy enums
 - No nested endianess
 - No encoding
 - No comparison of string, []byte or custom types
 - No min or max functions
 - fix type inference
 - -2 % 8 = -2
 - xor, ror, rol and zlib only work on bytes
 - float + int fails