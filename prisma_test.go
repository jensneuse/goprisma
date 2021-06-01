package goprisma

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	queryString = `postgresql://admin:admin@localhost:54321/example?schema=public&connection_limit=20&pool_timeout=5`
)

func TestIntrospect(t *testing.T) {
	schema := fmt.Sprintf(`datasource db {
		provider = "%s"
		url      = "%s"
	}`, "postgresql", queryString)

	schema, sdl, err := Introspect(schema)
	assert.NoError(t, err)
	assert.NotEqual(t, "", schema)
	assert.NotEqual(t, "", sdl)
}

func TestNewEngine(t *testing.T) {
	schema := fmt.Sprintf(`datasource db {
		provider = "%s"
		url      = "%s"
	}`, "postgresql", queryString)

	query := `{
            "query": "query Messages {findManymessages(take: 20 orderBy: [{id: desc}]){id message users {id name}}}",
            "variables": {}
        }`

	engine, err := NewEngine(schema)
	assert.NoError(t, err)
	assert.NotNil(t, engine)

	defer engine.Close()

	response,err := engine.Execute(query)
	assert.NoError(t, err)
	if strings.Contains(response, "errors") {
		t.Fatal(response)
	}
}

func BenchmarkEngine_Execute(b *testing.B) {
	schema := fmt.Sprintf(`datasource db {
		provider = "%s"
		url      = "%s"
	}`, "postgresql", queryString)

	query := `{
            "query": "query Messages {findManymessages(take: 20 orderBy: [{id: desc}]){id message users {id name}}}",
            "variables": {}
        }`

	engine, err := NewEngine(schema)
	if err != nil {
		b.Fatal(err)
	}

	defer engine.Close()

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			response,err := engine.Execute(query)
			if err != nil {
				b.Fatal(err)
			}
			if strings.Contains(response, "errors") {
				b.Fatal(response)
			}
		}
	})
}

func PrintMemUsage() {
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v KiB", bToMb(m.Alloc)/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
