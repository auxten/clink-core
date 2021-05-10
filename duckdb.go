// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package clinkcore implements a databse/sql driver for the clink-core database.
package clinkcore

/*
#cgo CFLAGS: -I/home/auxten/Codes/go/src/github.com/cwida/duckdb/src/include
#cgo LDFLAGS: -static /home/auxten/Codes/go/src/github.com/cwida/duckdb/build/release/src/libduckdb_static.a
#cgo LDFLAGS: /home/auxten/Codes/go/src/github.com/cwida/duckdb/build/release/third_party/utf8proc/libutf8proc.a
#cgo LDFLAGS: /home/auxten/Codes/go/src/github.com/cwida/duckdb/build/release/third_party/fmt/libfmt.a
#cgo LDFLAGS: /home/auxten/Codes/go/src/github.com/cwida/duckdb/build/release/third_party/re2/libduckdb_re2.a
#cgo LDFLAGS: /home/auxten/Codes/go/src/github.com/cwida/duckdb/build/release/third_party/miniz/libminiz.a
#cgo LDFLAGS: /home/auxten/Codes/go/src/github.com/cwida/duckdb/build/release/third_party/libpg_query/libpg_query.a
#cgo LDFLAGS: /home/auxten/Codes/go/src/github.com/cwida/duckdb/build/release/third_party/hyperloglog/libhyperloglog.a
#cgo LDFLAGS: -lm -lstdc++ -lgcc -ldl
#include <duckdb.h>
*/
import "C"

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"unsafe"
)

func init() {
	sql.Register("clink", impl{})
}

type impl struct{}

func (impl) Open(name string) (driver.Conn, error) {
	var db C.duckdb_database
	var con C.duckdb_connection

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	if err := C.duckdb_open(cname, &db); err == C.DuckDBError {
		return nil, errError
	}
	if err := C.duckdb_connect(db, &con); err == C.DuckDBError {
		return nil, errError
	}

	return &conn{db: &db, con: &con}, nil
}

var (
	errError = errors.New("could not open database")
)
