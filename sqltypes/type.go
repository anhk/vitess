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

package sqltypes

type Type int

const (
	Null       Type = 0
	VarBinary  Type = 10262
	Bit        Type = 2073
	Expression Type = 31
)

type Flag int32

const (
	flagIsQuoted Flag = 2048
)

// IsQuoted returns true if querypb.Type is a quoted text or binary.
// If you have a Value object, use its member function.
func IsQuoted(t Type) bool {
	return (int32(t)&int32(flagIsQuoted)) == int32(flagIsQuoted) && t != Bit
}
