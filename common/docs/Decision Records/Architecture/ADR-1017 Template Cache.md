# ADR-1017 Template Cache

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context

How exactly we want to cache templates.

## Decision



## Why / Notes

It seems like most of the examples are incomplete and don't do exactly what we want.

- See [Template Caching Example (Google AI)](https://github.com/SunnysideAaron/wetesa-0/blob/main/docs/Decision%20Records/Info%20Source%20Backups/Template%20Caching%20Example%20(Google%20AI).md)
- "Let's Go!" by Alex Edwards Chapter 5.3
  - loads all sub templates on every main template. Even one's that aren't called by the main template.
  - no sub directories?
- [How To Create A Template Cache For Your Golang Web Application](https://andrew-mccall.com/blog/2022/06/create-a-template-cache-for-a-go-application/)
  - very similar to Let's Go! by Alex Edwards
  - says it's based on Trevor Sawler’s udemy course "Building Modern Web Applications With Go"
- [template-parse-recursive](https://github.com/karelbilek/template-parse-recursive)
- [stencil](https://gitlab.com/kylehqcom/stencil)
  - [Go Templates — a final outcry](https://kylehqcom.medium.com/go-templates-a-final-outcry-1b8c9f7d046d)
- [Golang Templates — What I missed!](https://kylehqcom.medium.com/golang-templates-what-i-missed-abd1add92791)
  - https://gitlab.com/-/snippets/1662623

  
## Consequences



## Other Possible Options

## Not an Option

