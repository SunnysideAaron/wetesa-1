package database

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"strconv"
	"strings"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rbicker/go-rsql"

	"api/internal/shared-code/model"
)

// InsertClient adds a new client record to the database.
// It returns an error if the insert operation fails.
func (pg *Postgres) InsertClient(ctx context.Context, c model.Client) error {
	query := `INSERT INTO client (name, address) VALUES (@name, @address)`
	args := pgx.NamedArgs{
		"name":    c.Name,
		"address": c.Address,
	}

	_, err := pg.pool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

// BulkInsertClients is slower than CopyInserts. Use bulk inserts if you need to know a particular insert failed.
func (pg *Postgres) BulkInsertClients(ctx context.Context, clients []model.Client) error {
	query := `INSERT INTO client (name, address) VALUES (@name, @address)`

	batch := &pgx.Batch{} //nolint:exhaustruct // works fine. we don't need to initialize. already handled.

	for _, client := range clients {
		args := pgx.NamedArgs{
			"name":    client.Name,
			"address": client.Address,
		}
		batch.Queue(query, args)
	}

	results := pg.pool.SendBatch(ctx, batch)
	// defer results.Close()
	defer func() {
		err := results.Close()
		if err != nil {
			slog.LogAttrs(
				ctx,
				slog.LevelError,
				"could not close results",
				slog.String("error", err.Error()),
			)
		}
	}()

	for _, client := range clients {
		_, err := results.Exec()
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
				slog.LogAttrs(
					ctx,
					slog.LevelInfo,
					"user already exists",
					slog.String("name", client.Name),
				)
				continue
			}

			return fmt.Errorf("unable to insert row: %w", err)
		}
	}

	err := results.Close()
	if err != nil {
		return fmt.Errorf("failed to close batch results: %w", err)
	}
	return nil
}

// CopyInsertClients if faster than BulkInsertClients
func (pg *Postgres) CopyInsertClients(ctx context.Context, clients []model.Client) error {
	entries := [][]any{}
	columns := []string{"name", "address"}
	tableName := "client"

	for _, client := range clients {
		entries = append(entries, []any{client.Name, client.Address})
	}

	_, err := pg.pool.CopyFrom(
		ctx,
		pgx.Identifier{tableName},
		columns,
		pgx.CopyFromRows(entries),
	)

	if err != nil {
		return fmt.Errorf("error copying into %s table: %w", tableName, err)
	}

	return nil
}

type QryStrings struct {
	Columns string
	Where   string
	OrderBy string
	Limit   string
	PerPage int
	Args    []any
}

// validateURLParamSort validates and processes sort parameters
// func validateURLParamFields(fields string, allowedColumns map[string]bool) (string, error) {
// 	fieldParts := strings.Split(fields, ":")
// 	columns := []string{}

// 	for _, field := range fieldParts {
// 		if !allowedColumns[field] {
// 			return "", fmt.Errorf("Invalid field: %s", field)
// 		}
// 		columns = append(columns, field)
// 	}

// 	return " " + strings.Join(columns, ", "), nil
// }

// validateUrlParamQuery validates and processes query parameters
func validateUrlParamQuery(
	qs *QryStrings,
	urlParams url.Values,
	allowedColumns map[string]bool,
) (err error) {
	queryStr := urlParams.Get("q")

	if len(queryStr) == 0 {
		return nil
	}

	qs.Where = ""

	parser, err := rsql.NewParser(rsql.Mongo())
	if err != nil {
		fmt.Errorf("error while creating parser: %s", err)
	}
	s := `status=="A",qty=lt=30`
	res, err := parser.Process(s)
	if err != nil {
		fmt.Errorf("error while parsing: %s", err)
	}
	println("result", res)

	println("queryStr", queryStr)

	return err
}

// validateUrlParamSort validates and processes sort parameters
func validateUrlParamSort(
	qs *QryStrings,
	urlParams url.Values,
	allowedColumns map[string]bool,
) (err error) {
	sortStr := urlParams.Get("sort")

	if len(sortStr) == 0 {
		return nil
	}

	orderByParts := []string{}

	values := strings.Split(sortStr, ",")

	column := ""
	order := "ASC"

	for _, value := range values {
		switch value[0] {
		case '-':
			column = value[1:]
			order = "DESC"
		// encountered + being url encoded as space. Don't use for now. Don't NEED to
		// so OK to just leave off for now. Leaving here as documentation.
		// case '+':
		// 	column = value[1:]
		// 	order = "ASC"
		default:
			column = value
			order = "ASC"
		}

		if !allowedColumns[column] {
			return fmt.Errorf("Invalid sort field: '%s' allowed fields are: %v", column, allowedColumns)
		}

		orderByParts = append(orderByParts, fmt.Sprintf("%s %s", column, order))
	}

	if len(orderByParts) > 0 {
		qs.OrderBy = " ORDER BY " + strings.Join(orderByParts, ", ")
	}

	return err
}

func validateUrlParamPage(qs *QryStrings, urlParams url.Values) (page int, err error) {
	err = nil
	page = 0
	pageStr := urlParams.Get("page")

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			return 0, errors.New("Invalid page parameter")
		}
		if page < 0 {
			page = 0
		}
	}

	perPage := 10 //TODO make this 25 when have more test data
	perPageStr := urlParams.Get("per_page")

	if perPageStr != "" {
		perPage, err = strconv.Atoi(perPageStr)
		if err != nil {
			return 0, errors.New("Invalid per_page parameter")
		}
		// Enforce reasonable size limits
		if perPage > 100 {
			perPage = 100
		} else if perPage < 1 {
			perPage = 1
		}
	}

	qs.PerPage = perPage

	offset := page * qs.PerPage
	argPosition := len(qs.Args) + 1

	qs.Limit = fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPosition, argPosition+1)
	// Request one extra record to determine if there are more results
	qs.Args = append(qs.Args, qs.PerPage+1, offset)

	return page, err
}

// GetClientsParseParams parses the parameters for the GetClients query.
// any errors from here is a http bad request
func ValidateGetClientsParams(urlParams url.Values) (qs QryStrings, page int, err error) {
	qs.Columns = " client_id, name, address" // default columns
	allowedColumns := map[string]bool{
		"client_id": true,
		"name":      true,
		"address":   true,
	}

	// TODO fields. I'm punting for now to prevent premature optimization.
	// should fields add to default list, subrtract from list, or replace list?
	// wait till I have a use case to implement it.
	// fields := urlParams.Get("fields")
	// if fields != "" {
	// 	qs.Columns, err = validateURLParamFields(fields, allowedColumns)
	// 	if err != nil {
	// 		return qs, page, err
	// 	}
	// }

	qs.Where = ""
	qs.Args = []any{}
	// argPosition := 1
	err = validateUrlParamQuery(&qs, urlParams, allowedColumns)
	if err != nil {
		return qs, 0, err
	}

	qs.OrderBy = "  ORDER BY name ASC" // default sort
	// based on this limited example I wouldn't normally allow sorting on client_id or address on list client
	// but I want to demonstrate how to do sorting
	err = validateUrlParamSort(&qs, urlParams, allowedColumns)
	if err != nil {
		return qs, 0, err
	}

	page, err = validateUrlParamPage(&qs, urlParams)

	return qs, page, err
}

// GetClients retrieves a list of clients from the database.
// any errors from here is an internal server error
func (pg *Postgres) GetClients(
	ctx context.Context,
	qs QryStrings,
) ([]model.Client, bool, error) {
	query := `SELECT` + qs.Columns
	query += ` FROM client
		 	   WHERE 1=1`

	query += qs.Where
	query += qs.OrderBy
	query += qs.Limit

	slog.LogAttrs(
		ctx,
		slog.LevelDebug,
		"query",
		slog.String("query", query),
		slog.Any("args", qs.Args),
	)

	rows, err := pg.pool.Query(ctx, query, qs.Args...)
	if err != nil {
		return nil, false, fmt.Errorf("unable to query clients: %w", err)
	}

	defer rows.Close()

	clients, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[model.Client])
	if err != nil {
		return nil, false, fmt.Errorf("failed to collect client rows: %w", err)
	}

	// Check if we got more results than requested
	hasNext := false
	if len(clients) > qs.PerPage {
		hasNext = true
		// Remove the extra record before returning
		clients = clients[:qs.PerPage]
	}

	return clients, hasNext, nil
}

// GetClient retrieves a single client by their ID.
// Returns the client data if found, or an error if the client doesn't exist
// or if the query fails.
func (pg *Postgres) GetClient(ctx context.Context, id string) (model.Client, error) {
	var client model.Client

	query := `SELECT client_id, name, address FROM client WHERE client_id = $1`

	row := pg.pool.QueryRow(ctx, query, id)

	err := row.Scan(&client.ClientID, &client.Name, &client.Address)
	if err != nil {
		return client, fmt.Errorf("failed to get client: %w", err)
	}
	return client, nil
}

// UpdateClient updates an existing client's information.
// Returns an error if the client doesn't exist or if the update operation fails.
func (pg *Postgres) UpdateClient(ctx context.Context, c model.Client) error {
	query := `UPDATE client 
			  SET name = @name, address = @address 
			  WHERE client_id = @client_id`

	args := pgx.NamedArgs{
		"client_id": c.ClientID,
		"name":      c.Name,
		"address":   c.Address,
	}

	result, err := pg.pool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update client: %w", err)
	}

	// Check if any row was actually updated
	if result.RowsAffected() == 0 {
		return fmt.Errorf("client with id %s not found", c.ClientID)
	}

	return nil
}

// DeleteClient removes a client from the database by their ID.
// Returns an error if the client doesn't exist or if the deletion fails.
func (pg *Postgres) DeleteClient(ctx context.Context, id string) error {
	query := `DELETE FROM client WHERE client_id = $1`

	result, err := pg.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("unable to delete client: %w", err)
	}

	// Check if any row was actually deleted
	if result.RowsAffected() == 0 {
		return fmt.Errorf("client with id %s not found", id)
	}

	return nil
}
