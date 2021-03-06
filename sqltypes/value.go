/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package sqltypes implements interfaces and types that represent SQL values.
package sqltypes

import (
	"fmt"
	"github.com/anhk/vitess/bytes2"
	"github.com/anhk/vitess/hack"
)

var (
	// NULL represents the NULL value.
	NULL = Value{}

	// DontEscape tells you if a character should not be escaped.
	DontEscape = byte(255)

	nullstr = []byte("null")
)

// BinWriter interface is used for encoding values.
// Types like bytes.Buffer conform to this interface.
// We expect the writer objects to be in-memory buffers.
// So, we don't expect the write operations to fail.
type BinWriter interface {
	Write([]byte) (int, error)
}

// Value can store any SQL value. If the value represents
// an integral type, the bytes are always stored as a canonical
// representation that matches how MySQL returns such values.
type Value struct {
	typ Type
	val []byte
}

// MakeTrusted makes a new Value based on the type.
// This function should only be used if you know the value
// and type conform to the rules. Every place this function is
// called, a comment is needed that explains why it's justified.
// Exceptions: The current package and mysql package do not need
// comments. Other packages can also use the function to create
// VarBinary or VarChar values.
func MakeTrusted(typ Type, val []byte) Value {

	if typ == Null {
		return NULL
	}

	return Value{typ: typ, val: val}
}

// Type returns the type of Value.
func (v Value) Type() Type {
	return v.typ
}

// Raw returns the internal representation of the value. For newer types,
// this may not match MySQL's representation.
func (v Value) Raw() []byte {
	return v.val
}

// ToBytes returns the value as MySQL would return it as []byte.
// In contrast, Raw returns the internal representation of the Value, which may not
// match MySQL's representation for newer types.
// If the value is not convertible like in the case of Expression, it returns nil.
func (v Value) ToBytes() []byte {
	if v.typ == Expression {
		return nil
	}
	return v.val
}

// Len returns the length.
func (v Value) Len() int {
	return len(v.val)
}

// ToString returns the value as MySQL would return it as string.
// If the value is not convertible like in the case of Expression, it returns nil.
func (v Value) ToString() string {
	if v.typ == Expression {
		return ""
	}
	return hack.String(v.val)
}

// String returns a printable version of the value.
func (v Value) String() string {
	if v.typ == Null {
		return "NULL"
	}
	if v.IsQuoted() || v.typ == Bit {
		return fmt.Sprintf("%v(%q)", v.typ, v.val)
	}
	return fmt.Sprintf("%v(%s)", v.typ, v.val)
}

// EncodeSQL encodes the value into an SQL statement. Can be binary.
func (v Value) EncodeSQL(b BinWriter) {
	switch {
	case v.typ == Null:
		b.Write(nullstr)
	case v.IsQuoted():
		encodeBytesSQL(v.val, b)
	case v.typ == Bit:
		encodeBytesSQLBits(v.val, b)
	default:
		b.Write(v.val)
	}
}

// IsQuoted returns true if Value must be SQL-quoted.
func (v Value) IsQuoted() bool {
	return IsQuoted(v.typ)
}

func encodeBytesSQL(val []byte, b BinWriter) {
	buf := &bytes2.Buffer{}
	buf.WriteByte('\'')
	for _, ch := range val {
		if encodedChar := SQLEncodeMap[ch]; encodedChar == DontEscape {
			buf.WriteByte(ch)
		} else {
			buf.WriteByte('\\')
			buf.WriteByte(encodedChar)
		}
	}
	buf.WriteByte('\'')
	b.Write(buf.Bytes())
}

func encodeBytesSQLBits(val []byte, b BinWriter) {
	fmt.Fprint(b, "b'")
	for _, ch := range val {
		fmt.Fprintf(b, "%08b", ch)
	}
	fmt.Fprint(b, "'")
}

// SQLEncodeMap specifies how to escape binary data with '\'.
// Complies to http://dev.mysql.com/doc/refman/5.1/en/string-syntax.html
var SQLEncodeMap [256]byte

// SQLDecodeMap is the reverse of SQLEncodeMap
var SQLDecodeMap [256]byte

var encodeRef = map[byte]byte{
	'\x00': '0',
	'\'':   '\'',
	'"':    '"',
	'\b':   'b',
	'\n':   'n',
	'\r':   'r',
	'\t':   't',
	26:     'Z', // ctl-Z
	'\\':   '\\',
}

func init() {
	for i := range SQLEncodeMap {
		SQLEncodeMap[i] = DontEscape
		SQLDecodeMap[i] = DontEscape
	}
	for i := range SQLEncodeMap {
		if to, ok := encodeRef[byte(i)]; ok {
			SQLEncodeMap[byte(i)] = to
			SQLDecodeMap[to] = byte(i)
		}
	}
}
