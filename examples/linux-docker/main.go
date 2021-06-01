package main

import (
	"fmt"

	"github.com/jensneuse/goprisma"
)

const (
	queryString = `postgresql://admin:admin@host.docker.internal:54321/example?schema=public&connection_limit=20&pool_timeout=5`
)

func main() {

	schema := fmt.Sprintf(`datasource db {
		provider = "%s"
		url      = "%s"
	}`, "postgresql", queryString)

	schema, sdl, err := goprisma.Introspect(schema)
	if err != nil {
		panic(err)
	}
	fmt.Println(schema)
	fmt.Println(sdl)
}
