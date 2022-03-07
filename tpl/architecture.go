package tpl

func ArchitectureTemplate() []byte {
	return []byte(`
# Project Layout

This is a basic layout for our service. It uses a [standard-like](https://github.com/golang-standards/project-layout) layout with some modification.

The changes are made to adapt the following technique/principles:

- Clean Architecture/Hexagonal Architecture/Ports and Adapters Architecture.
- Domain-Driven Design.
- Command and Query Responsibility Segregation.

## Goals

- The business rules are testable by design and independent of frameworks, databases, and other external applications/interfaces.
- Ease to understand, use and work with.
- Maintainable when the project grow big.
- Tests are easy to write.

## Directories

` + "```" + `
├── api
├── cmd
│   └── root.go
├── docs
│   └── ARCHITECTURE.md
├── go.mod
├── go.sum
├── internal
│   ├── adapters
│   ├── app
│   ├── domain
│   └── ports
└── main.go
` + "```" + `

### ` + "`" + "/api" + "`" + `

OpenAPI/Swagger specs, JSON schema files, protocol definition files.

### ` + "`" + "/cmd" + "`" + `

Contain running code for programs of this service.

### ` + "`" + "/docs" + "`" + `

Design and user documents (include generated docs).

### ` + "`" + "/internal" + "`" + `

Our core application and library code that run the business. And inside this folder, we divide it into sub-folders which are:

- ` + "`" + "adapters" + "`" + `: infrastructure code that talk to the external world (DB queries, HTTP/gRPC clients, file reader/writer, PubSub producers,..)
- ` + "`" + "ports" + "`" + `: code that expose our application to the external world (HTTP/gRPC servers, PubSub subscribers,...)
- ` + "`" + "domain" + "`" + `: contains our domain entities and holds just business logic of our service.
- ` + "`" + "app" + "`" + `: a thin layer that glues together other layers.
`)
}
