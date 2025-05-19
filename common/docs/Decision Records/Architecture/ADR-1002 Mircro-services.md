# ADR-1002 Micro-services

## Status

Accepted

## Context

At very large scales micro-services are a standard way to split concerns and more importantly handle the work load.

## Decision

We will **not** implement micro-services until we are forced to in order to meet scale or break up team responsibilities. If the only reason is team responsibilities then consider the creation of external packages.

When we implement micro-services we will first find a set of guidelines to follow. (perhaps a book) we will research the cons of those guidelines as well.

we don't need event driven architecture until we do micro-services.

## Why / Notes

Micro-services add complexity and are very costly in developer time and support. They are easy to implement poorly.

Micro-services are an organizational solution not a technical solution.

[So you want to break down monolith?](https://www.architecture-weekly.com/p/so-you-want-to-break-down-monolith)

## Consequences

Incurs technical debt when we have to refactor. Our application might struggle to handle scale for a bit until we implement.

## Other Options

Possibilities:

- [go-zero](https://github.com/zeromicro/go-zero)
  - chinese.

Not an option:

