# TSDR-1002 DB migrations  

## Status

Pending (waiting to determine if sqlc will be part of this.)

## Context

We will have to push database changes to production. How?

## Decision



## Why / Notes



## Consequences



## Other Options

Possibilities:
- [Tern](https://github.com/jackc/tern)
  - From the person behind pgx
- Goose
- https://github.com/golang-migrate/migrate

Not an option:
- [Bun](https://bun.uptrace.dev/guide/)
  - We are already using pgx. Bun comes with a lot more stuff.
