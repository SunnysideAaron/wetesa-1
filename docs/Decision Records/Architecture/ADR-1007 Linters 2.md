# ADR-100
# TSDR-100

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context



## Decision

Will use the following linters:
- https://github.com/go-simpler/sloglint

A linter like sloglint can help you enforce various rules for Slog based on your preferred code style. Here's an example configuration when used through golangci-lint:


## Why / Notes

 - [golangci-lint](https://golangci-lint.run/)
 - [revive](https://revive.run/)
  - TODO Additional rules

  aaron does not like single line if else statements
they easily hide scope and lessen readability.
have to reedit later when maintaining anyway.
can these be linted out?
    https://www.calhoun.io/one-liner-if-statements-with-errors/

    markdown linter

- https://github.com/avelino/awesome-go?tab=readme-ov-file#code-analysis


https://www.reddit.com/r/golang/comments/1in0tiw/simple_strategy_to_understand_error_handling_in_go/
errorlint in golangci-lint will throw a lint warning if you use == or != to compare an error instead of errors.Is. I find it very useful, and it is even in the class of linter that has caught an actual bug in my code and not just given style suggestions. I recommend it.

Note that this is not errcheck, which is on by default. errorlint must be enabled explicitly.

API Toolkit
	project for verifing api contract.

Look up United Kingdoms GDS departments website access guidlines
	?must work without javascript?

TODO
Lint SQL
  - [Postreges Wiki: Don't Do This](https://wiki.postgresql.org/wiki/Don't_Do_This)
  - https://github.com/kristiandupont/schemalint

I liked the uber or google style guide that said things with no consensus. good to mark that out.


100 mistakes book to style guide / linter

## Consequences



## Other Possible Options

## Not an Option

