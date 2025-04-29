# ADR-100
# TSDR-100

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context

We want code generation for (one will be source, others derived from):
- DB Model / Chart / Diagram
- DB Schema
- DB Migrations
- API Model / DB Communication (should include basic validation)
- API Routing
- API Validation
- API controllers / handlers
- OpenAPI 3 (Swagger 1 and 2 are previous versions of OpenAPI) (API Documentation)
- Web Client use of API
- Web Client Validation
- Mobile Client use of API
- Mobile Client Validation
- SDKs for other languages for our clients

Code generation should allow for:
- Pagination
- Filters
- Overrides when we want our own code.

## Decision



## Why / Notes



## Consequences



## Short List of Options

- SQLc
- SQLc custom plugin
- open api codegen -> fallthrough podcast
- [oto](https://github.com/pacedotdev/oto/tree/main/otohttp) by Mat Ryer **liked this** can generate anything I want from templates.
- scrutinizer?
- goa (retry?)

## Other Possible Options

- [awesome-go](https://github.com/avelino/awesome-go?tab=readme-ov-file#go-generate-tools)
- [reddit confused_by_the_openapi_options](https://www.reddit.com/r/golang/comments/1gmhz08/confused_by_the_openapi_options_for_go/)
- [reddit](https://www.reddit.com/r/golang/comments/1avsog1/go_openapi_codegen/)

- Code → Spec
  - Astra by
  - (Gin-only, Echo & Fiber WIP as of 2024-02-22)
  - Fuego (built with go 1.22 uses generics and standard http started end of 2023)
  - [Gleece](https://github.com/gopher-fleece/gleece)
    - Code -> Code + Spec
    - Generates routers for Gin, Echo v4, Gorilla Mux, Chi v5, and Fiber v2 routers. Probably I could add standard library router.
    - Generates HTTP routing generation, authorization enforcement, payload validation, error handling, and OpenAPI v3 specification generation
    - https://zuplo.com/blog/2025/03/04/holistic-view-of-apis
  - Huma by
  - Tonic
  - libopenapi
  - kin-openapi (seems semi-abandoned as of 2024-02-22)
  - go restful
    - generates open api spec from code.
- Spec → SDK
  - https://packagemain.tech/p/practical-openapi-in-golang **SHORT LIST TO TRY**
  - [Laurence de Jong](https://ldej.nl/post/generating-go-from-openapi-3/)
  - oapi-codegen
    doesn't do OpenAPI 3?
  - ogen **TRIED**
  - openapi-generator
  - swagger-codegen
  - microsoft/kiota
- DSL → Spec + Code
  - [Goa](https://github.com/goadesign/goa)
    - Goa provides a holistic approach for developing remote APIs and microservices in Go.
    - **TRIED** may give another chance.
- Unknown if Spec or SDK first
  - [Prism](https://stoplight.io/open-source/prism)
    - use for testing?
  - [swaggest - rest](https://github.com/swaggest/rest)
  - [swaggest - openapi-go](https://github.com/swaggest/openapi-go)

## Not an Option

- swaggo/swag (OAS3 Beta)
  - **TRIED** supported frameworks:
  - gin, echo, buffalo, net/http (standard library), gorilla/mux, go-chi/chi,
    flamingo, fiber, atreugo, hertz
- [speakeasy](https://www.speakeasy.com/docs/languages/golang/oss-comparison-go)
  - trying to sell speakeasy
  - generates sdks
