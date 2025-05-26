# TSDR-1002 DB migrations  

## Status

Pending (waiting to determine if sqlc will be part of this.)

## Context

We will have to push database changes to production. How?

## Decision



## Why / Notes

pick something compatible with sqlc https://docs.sqlc.dev/en/latest/howto/ddl.html

## Consequences



## Other Options

- [Awesome Go - Database](https://github.com/avelino/awesome-go?tab=readme-ov-file#database)
- [Tern](https://github.com/jackc/tern)
  - From the person behind pgx
- Goose
- https://github.com/golang-migrate/migrate
- https://www.reddit.com/r/golang/comments/1kghyca/how_to_manage_database_schema_in_golang/
- https://github.com/xataio/pgroll

Not an option:
- [Bun](https://bun.uptrace.dev/guide/)
  - We are already using pgx. Bun comes with a lot more stuff.
