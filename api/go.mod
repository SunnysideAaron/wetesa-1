module api

go 1.24

toolchain go1.24.4

require (
	github.com/jackc/pgerrcode v0.0.0-20240316143900-6e2875d9b438
	github.com/jackc/pgx/v5 v5.7.5
	github.com/rbicker/go-rsql v0.0.0-00010101000000-000000000000
// shared-code v0.0.0
)

// replace shared-code => ./shared-code
// TODO get our fixes merged into their library.
replace github.com/rbicker/go-rsql => ./internal/shared-code/go-rsql

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/sync v0.13.0 // indirect
	golang.org/x/text v0.24.0 // indirect
)
