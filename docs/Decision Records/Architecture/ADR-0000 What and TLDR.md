# ADR-0000 What and TLDR

## What we are building.

See "../Tech Stack/TSDR-0000 What and TLDR.md"

## TL;DR of decisions

- All Decisions from wetesa-0
  - These will be pulled into example 1 later. Once that example is "finished".
- ADR-1001 Cryptography
  - Will not roll our own. **PENDING** decide which packages to use.
- ADR-1002 Microservices
  - We will **not** implement micro-services until we are forced to in order to meet scale
- ADR-1003
  - Dates will always be presented in YYYY-MMM-DD format. Leading 0s on DD.  YEAR-MM-DD is allowable for data entry

TODO

https://www.moesif.com/blog/api-analytics/api-strategy/API-Logs/
Performance Logs
Security Logs


https://www.reddit.com/r/devops/comments/1jkgdiq/how_much_do_you_spend_on_cicd/



**PENDING**

DB default encodings
utf8 (other encodings?)

UTC for server and db

ACID Compliant

DB knows user logged in. Not just a general web log in.

how to handle api breaking changes.



- Alex Edwards’ book “Let’s Go Further”
  - 6 sql migrations
  - 8 Advanced crud
    - partial updates
    - optimistic concurancy (two updates to same records same time)
    - sql query timeouts
  - 11 rate limiting
  - 13, 15-18 authentication and authorization
  - 14 sending emails
  - 19 metrics
  - 20 QA, versioning
  - 22 Appendices