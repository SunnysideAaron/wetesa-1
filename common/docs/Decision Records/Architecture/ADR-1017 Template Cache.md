# ADR-1017 Template Cache

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context

How exactly we want to cache templates.

## Decision

**For now** we will use the method from "Let's Go!" by Alex Edwards. With some minor
adjustments. We will adjust later if needed.

Don't use sub directories.

## Why / Notes

It seems like most of the examples are incomplete and don't do exactly what we want.
caching and parsing of main templates isn't hard. But referenced templates get's
wonky.

parsing of templates that reference other templates into a cache gets kinda wonky.
do to the way that `ParseFiles` works. Note there are two ParseFiles() one function one method.
https://go.dev/src/text/template/helper.go#L70

It appears that ParseFiles
1. creates a template with the base filename (without path).
2. creates templates named from {{define "name"}} code as well.
this can result in duplicates. but it makes sense since one template can have
multiple {{define}} blocks.

sub templates get added by file name and all of the defines from that file.

Go's template doesn't really like sub directories. It can be forced but in practice
makes things clunky.

**TODO** write up a demonstration apps that shows how convoluted wonky this all gets.
showing multiple ways of doing it.

## Sources

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

This seems to be something that many projects grapple with. Without a clear answer.
I see a few other options.

1. We could create a listing of each page and all of the templates it uses. Then passing that to ParseFiles.
  - big con would be maintaining that list.
  - a highbred would be to only have an override list. so when a specific page needs
    additional templates we can load just those.
2. We could try some kind of recursive reading of the templates directory to find each template
  - this gets weird because the {{template}} and {{define}} blocks don't have to have the file name.
  - so either we make them set to file names or we read each file and each define
  - to try and match them up. Even more complex if we have sub directories and allow
    defines to have the same name across files.
3. one thought would be to force the name of the sub templates to include the path and file name?

## Not an Option

