# ADR-0000 TL;DR of Decisions

- All Decisions from wetesa-0
  - These are copied from wetesa-0 so AI can find them. Don't edit them here.
    Create a new decision record.
- ADR-1001 Cryptography
  - Will not roll our own. **PENDING** decide which packages to use.
- ADR-1002 Microservices
  - We will **not** implement micro-services until we are forced to in order to
    meet scale
- ADR-1003
  - Dates will always be presented in YYYY-MMM-DD format. Leading 0s on DD.  YEAR-MM-DD is allowable for data entry
- ADR-1004 Authentication
  - **PENDING**
- ADR-1005 Authorization
  - **PENDING**
- ADR-1006 Sessions
  - **PENDING**
- ADR-1007 Linters 2
  - **PENDING**
- ADR-1008 CSS Coding Standards
  - **PENDING**
- ADR-1009 Security
  - **PENDING**
- ADR-1010 Audit Tables
  - **PENDING**
- ADR-1011 Validation
  - **PENDING**
- ADR-1012 SQL Coding Standards
  - **PENDING**
- ADR-1013 Common container
  - common docker container
- ADR-1014 Pagination
  - **PENDING**
- ADR-1015 Middleware 2
  - **PENDING**
- ADR-1016 App Struct vs Parameters
  - Pass parameters not a struct.



## TODO

https://www.moesif.com/blog/api-analytics/api-strategy/API-Logs/
Performance Logs
Security Logs


https://www.reddit.com/r/devops/comments/1jkgdiq/how_much_do_you_spend_on_cicd/


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

https://pb33f.io/wiretap/

use long form flags in make files. self documenting.

https://www.reddit.com/r/webdev/comments/1k2uuoh/i_hate_timezones/

https://www.reddit.com/r/golang/comments/1k3sb3j/json_schema_to_go_struct_or_alternatives/


https://threedots.tech/post/common-anti-patterns-in-go-web-applications/