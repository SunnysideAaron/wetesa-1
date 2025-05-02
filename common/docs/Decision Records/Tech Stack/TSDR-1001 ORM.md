# TSDR-1001 ORM

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context

I hate ORMs. Avoid if possible. We did think about them.

## Decision

sqlc?

Squirrel or sqrl. for more complex queries?

jet? 
  - https://www.reddit.com/r/golang/comments/1j6fzsz/we_made_writing_typesafe_sql_queries_in_go_even/
  - https://github.com/go-jet/jet

Comparing database/sql, GORM, sqlx, and sqlc
https://blog.jetbrains.com/go/2023/04/27/comparing-db-packages/

### sqlc

[sqlc](https://sqlc.dev/)

Plugins:
- [sqlc-gen-go](https://github.com/sqlc-dev/sqlc-gen-go)
  - code gen plugin template
- [sqlc-gen-go-server](https://github.com/walterwanderley/sqlc-gen-go-server)  
- [sqlc-http](https://github.com/walterwanderley/sqlc-http)
  - generates server from sqlc output.
- [Using plugins](https://docs.sqlc.dev/en/stable/guides/plugins.html)

Notes:

- [Reddit: How popular is sqlc in production go projects?](https://www.reddit.com/r/golang/comments/1imhs5l/how_popular_is_sqlc_in_production_go_projects/)
  - good stuff about pagination and dynamic queries concerns
  - example code
- [4 Tips for Working with Sqlc in Go](https://haykot.dev/blog/4-tips-for-working-with-sqlc-in-go/)
- [Reddit: sqlc dynamic queries with pagination and filters](https://www.reddit.com/r/golang/comments/1aiooft/sqlc_dynamic_queries_with_pagination_and_filters/)
  - pagination and dynamic queries
- [Reddit: Best practices with sqlc and dynamic filters](https://www.reddit.com/r/golang/comments/183292y/best_practices_with_sqlc_and_dynamic_filters/?share_id=fTJsRkD8PPGBA2jm4VXhD&utm_medium=android_app&utm_name=androidcss&utm_source=share&utm_term=1)
- [Is there a way for sqlc to generate code that can use pgxpool](https://stackoverflow.com/questions/76848733/is-there-a-way-for-sqlc-to-generate-code-that-can-use-pgxpool)

Intro Guides:
- [Creating an API using Go and sqlc](https://dev.to/eminetto/creating-an-api-using-go-and-sqlc-364o)
- [Bringing it Together: Go, SQL, Code Gen](https://brojonat.com/posts/go-postgres-sqlc-atlas/)
  - shows migrations using atlas
- [Introducing sqlc](https://conroy.org/introducing-sqlc) (2019-12)

Misc Links:
- [sqlc-pgx-monitoring ](https://github.com/amirsalarsafaei/sqlc-pgx-monitoring)


## Why / Notes

- Sounds amazing.

## Consequences

## Other Options

Possibilities:
- [Bun](https://bun.uptrace.dev/)
- GORM
- [ent.](https://entgo.io/)
- [Awesome Go: ORMs](https://github.com/avelino/awesome-go?tab=readme-ov-file#orm)
