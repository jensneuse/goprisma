# GoPrisma - a Go wrapper for the Prisma Engines

## What's this?

Introspect a database and use it as a GraphQL API using Go.

Supported Databases:
- [x] SQLite
- [x] PostgreSQL
- [x] MySQL
- [ ] MongoDB
- [ ] MS SQL Server

### go get

```shell
go get github.com/jensneuse/goprisma
```

### Introspect a database

```go
prismaSchema := "datasource db {provider = \"postgresql\" url = \"postgresql://admin:admin@localhost:54321/example?schema=public&connection_limit=20&pool_timeout=5\"}"
schema, sdl, err := Introspect(prismaSchema)
```

### Make a Query

```go
query := "{\"query\": \"query Messages {findManymessages(take: 20 orderBy: [{id: desc}]){id message users {id name}}}","variables\": {}}"
engine, err := NewEngine(schema)
if err != nil {
	return
}
defer engine.Close()
response := engine.Execute(query)
```

## Why?

GraphQL is a nice abstraction layer on top of other APIs.
Making databases available can be a productivity gainer when no direct database access is required and GraphQL clients already exist in the codebase.

## How does it work

Prisma is a really powerful ORM.
If you look closely, you'll realize that internally,
Prisma 2 is powered by a GraphQL Engine to abstract away the database layer.

The GraphQL Engine is written in Rust.

This library is a CGO wrapper to make the Prisma Rust GraphQL Engine available to write Go programs on top of it.

## Is it fast?

I've measured ~0.3ms latency (Go -> Rust -> Database -> Go) on my laptop using a database in docker.

```
goos: darwin
goarch: amd64
pkg: github.com/jensneuse/goprisma/pkg/prisma
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkEngine_Execute
BenchmarkEngine_Execute-16    	   43815	    304507 ns/op	     144 B/op	       3 allocs/op
PASS
```

You can find the benchmark in the prisma package tests and run it yourself.

## What architectures are supported?

- [x] darwin x86_64
- [x] darwin aarch64
- [x] linux x86_64
- [ ] windows x86_64 (needs verification)

Other targets can be added.

## Where is the Rust wrapper?

https://github.com/jensneuse/prisma-engines/tree/goprisma

## Example / How to run the tests?

First, clone the repo.

```shell
cd docker-example
docker-compose up
```

Now you're able to run the tests.

## How can you help?

### Testing

You can help by running the tests and architectures that need verification.
Please report any issues and tell me if it works.

### Automate building the C-Wrapper

You might want to increase the level of trust by adding a github action to automate building the Rust C-Wrapper.
Currently, I'm building it on my MacBook.

Ideally, this github action would fetch the latest stable of prisma-engines, compile the C-Wrapper and updates the lib folder via PR.
A blueprint how to do this can be found here: https://github.com/rogchap/v8go/blob/master/.github/workflows/v8build.yml

If you're interested in adding this, please open a PR.

## Considerations

It might be obvious, but I just want to make it clear that you should NOT publicly expose your database with this library.
This is only intended for use cases where the GraphQL API is shielded by some other component.
