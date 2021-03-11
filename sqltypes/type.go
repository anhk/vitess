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

import (
	querypb "github.com/anhk/vitess/vt/proto/query"
)

const (
	flagIsQuoted = int(querypb.Flag_ISQUOTED)
)

// IsQuoted returns true if querypb.Type is a quoted text or binary.
// If you have a Value object, use its member function.
func IsQuoted(t querypb.Type) bool {
	return (int(t)&flagIsQuoted == flagIsQuoted) && t != Bit
}

const (
	Null = querypb.Type_NULL_TYPE
	//Int8       = querypb.Type_INT8
	//Uint8      = querypb.Type_UINT8
	//Int16      = querypb.Type_INT16
	//Uint16     = querypb.Type_UINT16
	//Int24      = querypb.Type_INT24
	//Uint24     = querypb.Type_UINT24
	//Int32      = querypb.Type_INT32
	//Uint32     = querypb.Type_UINT32
	//Int64      = querypb.Type_INT64
	//Uint64     = querypb.Type_UINT64
	//Float32    = querypb.Type_FLOAT32
	//Float64    = querypb.Type_FLOAT64
	//Timestamp  = querypb.Type_TIMESTAMP
	//Date       = querypb.Type_DATE
	//Time       = querypb.Type_TIME
	//Datetime   = querypb.Type_DATETIME
	//Year       = querypb.Type_YEAR
	//Decimal    = querypb.Type_DECIMAL
	//Text       = querypb.Type_TEXT
	//Blob       = querypb.Type_BLOB
	//VarChar    = querypb.Type_VARCHAR
	VarBinary = querypb.Type_VARBINARY
	//Char       = querypb.Type_CHAR
	//Binary     = querypb.Type_BINARY
	Bit = querypb.Type_BIT
	//Enum       = querypb.Type_ENUM
	//Set        = querypb.Type_SET
	//Geometry   = querypb.Type_GEOMETRY
	//TypeJSON   = querypb.Type_JSON
	Expression = querypb.Type_EXPRESSION
)
