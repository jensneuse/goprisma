// package prisma exposes a CGO wrapper around the prisma introspection- and query engine to make it available to go
package goprisma

/*
#include <stdlib.h>
#include "lib/query_engine_c_api.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"

	// do not remove these side effect imports
	// they are required so that go mod properly imports the cgo files
	_ "github.com/jensneuse/goprisma/lib"
	_ "github.com/jensneuse/goprisma/lib/darwin"
	_ "github.com/jensneuse/goprisma/lib/darwin-aarch64"
	_ "github.com/jensneuse/goprisma/lib/linux"
	_ "github.com/jensneuse/goprisma/lib/windows"
)

func Introspect(schema string) (prismaSchema, prismaSDL string, err error) {
	inputSchemaCstring := C.CString(schema)
	defer C.free(unsafe.Pointer(inputSchemaCstring))
	introspectionResult := C.prisma_introspect(inputSchemaCstring)
	if introspectionResult == nil {
		return "", "", errors.New("introspection failed")
	}
	defer C.free_introspection_result(introspectionResult)
	if introspectionResult.error != nil {
		introspectionError := C.GoString(introspectionResult.error)
		return "", "", errors.New(introspectionError)
	}
	prismaSchema = C.GoString(introspectionResult.schema)
	prismaSDL = C.GoString(introspectionResult.sdl)
	return
}

// Engine is holding a C pointer of the prisma engine created in Rust
type Engine struct {
	ptr C.PrismaPtr
}

// NewEngine requires a valid prisma schema with one db defined
// please read the prisma docs for further information
//
// The returned Engine object contains a pointer to the prisma Rust engine
// the engine is already connected
// If you want to disconnect it, call Close()
//
// You're able to configure the number of database connections in the connection pool, timeouts etc. by using standard prisma configuration
// More info here: https://www.prisma.io/docs/concepts/components/prisma-client/working-with-prismaclient/connection-pool
//
// Do not forget to call Close() before the Engine struct goes out of scope,
// otherwise you have a memory leak!
func NewEngine(schema string) (*Engine, error) {
	inputSchemaCstring := C.CString(schema)
	defer C.free(unsafe.Pointer(inputSchemaCstring))
	ptr := C.prisma_new(inputSchemaCstring)
	if ptr == nil {
		return nil,fmt.Errorf("unable to create engine")
	}
	return &Engine{
		ptr: ptr,
	}, nil
}

// Execute takes a JSON encoded GraphQL query and executes it using the Rust prisma engine
// make sure to add an empty variables object to the query, otherwise the request will fail
// that said, variables are not supported by prisma, all values must be inlined into the query
func (e *Engine) Execute(query string) (string,error) {
	if e == nil || e.ptr == nil {
		return "",fmt.Errorf("engine uninitialized")
	}
	queryCString := C.CString(query)
	defer C.free(unsafe.Pointer(queryCString))
	responseCString := C.prisma_execute(e.ptr, queryCString)
	defer C.free(unsafe.Pointer(responseCString))
	return C.GoString(responseCString),nil
}

// Close disconnects the Rust prisma engine from the database and cleans up all structs and pointers on both sides of the bridge.
// If you forget to call Close before the Engine struct goes out of scope, you have a memory leak!
func (e *Engine) Close() {
	if e == nil || e.ptr == nil {
		return
	}
	C.free_prisma(e.ptr)
	e.ptr = nil
}
