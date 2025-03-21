# ADR-100
# TSDR-100

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context

We want code generation for (should include basic validation):
- API End points
- API to DB crud
- DB Schema
- OpenAPI 3 (Swagger 1 and 2 are previous versions of OpenAPI)
- Web Client use of API
- Mobile Client use of API

## Decision

short list of options:
- SQLc? 
- [oto](https://github.com/pacedotdev/oto/tree/main/otohttp) by Mat Ryer **liked this** can generate anything I want from templates.
- goa (retry?)

## Why / Notes



## Consequences



## Other Possible Options

- [awesome-go](https://github.com/avelino/awesome-go?tab=readme-ov-file#go-generate-tools)
- [reddit confused_by_the_openapi_options](https://www.reddit.com/r/golang/comments/1gmhz08/confused_by_the_openapi_options_for_go/)
- [Laurence de Jong](https://ldej.nl/post/generating-go-from-openapi-3/)
- [Goa](https://github.com/goadesign/goa) - Goa provides a holistic approach for developing remote APIs and microservices in Go.
  - may give another chance.
- Gleece (OpenApi generation. code generation?)
  - https://zuplo.com/blog/2025/03/04/holistic-view-of-apis
- [reddit](https://www.reddit.com/r/golang/comments/1avsog1/go_openapi_codegen/)
- Code → Spec
  - Astra by
  - (Gin-only, Echo & Fiber WIP as of 2024-02-22)
  - Huma by
  - Fuego (built with go 1.22 uses generics and standard http started end of 2023)
  - Tonic
  - libopenapi
  - kin-openapi (seems semi-abandoned as of 2024-02-22)
  - go restful
    - generates open api spec from code.
- Spec → SDK
  - https://packagemain.tech/p/practical-openapi-in-golang **SHORT LIST TO TRY**
  - oapi-codegen
    doesn't do OpenAPI 3?
  - ogen **TRIED**
  - openapi-generator
  - swagger-codegen
  - microsoft/kiota
- DSL → Spec + Code
  - goa  **TRIED**
- Unknown if Spec or SDK first
  - [swaggest - rest](https://github.com/swaggest/rest)
  - [swaggest - openapi-go](https://github.com/swaggest/openapi-go)

## Not an Option
- swaggo/swag (OAS3 Beta)
  - **TRIED** supported frameworks:
  - gin, echo, buffalo, net/http (standard library), gorilla/mux, go-chi/chi,
    flamingo, fiber, atreugo, hertz
- - [speakeasy](https://www.speakeasy.com/docs/languages/golang/oss-comparison-go) trying to sell speakeasy




