# ADR-1018 API Guidelines

## 1. Introduction

This ADR is a bit different because there are so many minor and interdependent decisions to be made.
We are starting with [Zalando RESTful API and Event Guidelines](https://opensource.zalando.com/restful-api-guidelines/)
and then adjusting to fit our own decisions.

## 2. Principles

### Our API is NOT a product

The APIs primary user is our web and mobile clients. The project that will be
based on WETESA-1 is not a public API. Not for the wider world. While
it may someday be used by third parties. They are not the primary concern.

### Solution first. Not API first

Most of the documentation on APIs say "API First". We are not building a public
API for consumption by the world. We are building a business (crud)solution for users.
To be honest my brain isn't wired to think in resources. I think in user problems,
user input, reports, and database structures. It may be shooting myself in the foot. but for
now I'm going to tailer the API to the screens the UI needs. Hopefully this will
cut down on over and under fetching, prevent the API from just mirroring the database
model, and prematurely optimizing sorting, filtering, and embedding.

## 3. General Guidelines

## 4. REST Basics - Meta information

## 5. REST Basics - Security

### MUST secure endpoints

TODO How

### MUST define and assign permissions (scopes)

TODO How

### PENDING MUST follow the naming convention for permissions (scopes)

https://opensource.zalando.com/restful-api-guidelines/#225

## 6. REST Basics - Data formats

### PENDING MUST use standard data formats

TODO go through and pick formats.
https://opensource.zalando.com/restful-api-guidelines/#238
https://opensource.zalando.com/restful-api-guidelines/#publications-specifications-and-standards

### PENDING MUST define a format for number and integer types

### PENDING MUST use standard formats for date and time properties

### PENDING SHOULD select appropriate one of date or date-time format

### PENDING SHOULD use standard formats for time duration and interval properties

### PENDING MUST use standard formats for country, language and currency properties

### PENDING SHOULD use content negotiation, if clients may choose from different resource representations

### SHOULD use UUIDs

See: ADR-002 Database Columns

## 7. REST Basics - URLs

### PENDING SHOULD not use /api as base path

### MUST pluralize resource names

### MUST use URL-friendly resource identifiers

### MUST use kebab-case for path segments

### MUST use normalized paths without empty path segments and trailing slashes

### MUST keep URLs verb-free

### MUST avoid actions — think about resources

### SHOULD define useful resources

### MUST use domain-specific resource names

### SHOULD model complete business processes

### MUST identify resources and sub-resources via path segments

### MAY expose compound keys as resource identifiers

### MAY consider using (non-) nested URLs

### SHOULD limit number of resource types

Here Zalando says 4 to 8 resource types. That only applies because they are
building using microservices. This will be a monolith. We will have many more
resource types but let's see if we can limit them and not be too crazy or just mirror db.

### SHOULD limit number of sub-resource levels

Avoid requiring resource URIs more complex than collection/item/collection.

2 levels OK. 3 Max

"There are main resources (with root url paths) and sub-resources (or nested resources with non-root urls paths). Use sub-resources if their life cycle is (loosely) coupled to the main resource, i.e. the main resource works as collection resource of the subresource entities. You should use ⇐ 3 sub-resource (nesting) levels — more levels increase API complexity and url path length. (Remember, some popular web browsers do not support URLs of more than 2000 characters.)"

### PENDING MUST use snake_case (never camelCase) for query parameters

### PENDING MUST stick to conventional query parameters

If you provide query support for searching, sorting, filtering, and paginating, you must stick to the following naming conventions:

- **q**: default query parameter, e.g. used by browser tab completion; should have an entity specific alias, e.g. sku.
- **sort**: comma-separated list of fields (as defined by MUST define collection format of header and query parameters) to define the sort order. To indicate sorting direction, fields may be prefixed with + (ascending) or - (descending), e.g. /sales-orders?sort=+id.
- **filter**: comma-separated list of fields (as defined by MUST define collection format of header and query parameters) to define the filter criteria. To indicate filtering direction, fields may be prefixed with + (include) or - (exclude), e.g. /sales-orders?filter=-cancelled.
- **fields**: field name expression to retrieve only a subset of fields of a resource. See SHOULD support partial responses via filtering below.
- **embed**: field name expression to expand or embedded sub-entities, e.g. inside of an article entity, expand silhouette code into the silhouette object. Implementing embed correctly is difficult, so do it with care. See SHOULD allow optional embedding of sub-resources below.
- **offset**: numeric offset of the first element provided on a page representing a collection request. See REST Design - Pagination section below.
- **cursor**: an opaque pointer to a page, never to be inspected or constructed by clients. It usually (encrypted) encodes the page position, i.e. the identifier of the first or last page element, the pagination direction, and the applied query filters to recreate the collection. See Cursor-based pagination in RESTful APIs or REST Design - Pagination section below.
- **limit**: client suggested limit to restrict the number of entries on a page. See REST Design - Pagination section below.

## 8. REST Basics - JSON payload

### MUST use JSON as payload data interchange format

### PENDING SHOULD design single resource schema for reading and writing

### PENDING SHOULD be aware of services not fully supporting JSON/unicode

### MAY pass non-JSON media types using data specific standard formats

### PENDING SHOULD use standard media types

### PENDING SHOULD pluralize array names

### PENDING MUST property names must be snake_case (and never camelCase)

### PENDING SHOULD declare enum values using UPPER_SNAKE_CASE string

### PENDING SHOULD use naming convention for date/time properties

### PENDING SHOULD define maps using additionalProperties

### PENDING MUST use same semantics for null and absent properties

### PENDING MUST not use null for boolean properties

### PENDING SHOULD not use null for empty arrays

### PENDING MUST use common field names and semantics

### PENDING MUST use the common address fields

### PENDING MUST use the common money object

## 9. REST Basics - HTTP requests

### PENDING MUST use HTTP methods correctly

### PENDING MUST fulfill common method properties

### PENDING SHOULD consider to design POST and PATCH idempotent

### PENDING SHOULD use secondary key for idempotent POST design

### MAY support asynchronous request processing

### PENDING MUST define collection format of header and query parameters

### PENDING SHOULD design simple query languages using query parameters

### PENDING SHOULD design complex query languages using JSON

### PENDING MUST document implicit response filtering

## 10. REST Basics - HTTP status codes

### MUST use official HTTP status codes

### PENDING MUST specify success and error responses

### PENDING SHOULD only use most common HTTP status codes

### PENDING MUST use most specific HTTP status codes

### PENDING MUST use code 207 for batch or bulk requests

### PENDING MUST use code 429 with headers for rate limits

### PENDING MUST support problem JSON

### MUST not expose stack traces

### PENDING SHOULD not use redirection codes

## 11. REST Basics - HTTP headers

## 12. REST Design - Hypermedia

### MAY use href links on sub items.

Use these if the UI will use them.

```json
{
  "id": "446f9876-e89b-12d3-a456-426655440000",
  "name": "Peter Mustermann",
  "spouse": {
    "href": "https://...",
    "since": "1996-12-19",
    "id": "123e4567-e89b-12d3-a456-426655440000",
    "name": "Linda Mustermann"
  }
}
```

### MUST use full, absolute URI for resource identification

Links to other resource must always use full, absolute URI.

Motivation: Exposing any form of relative URI (no matter if the relative URI uses an absolute or relative path) introduces avoidable client side complexity. It also requires clarity on the base URI, which might not be given when using features like embedding subresources. The primary advantage of non-absolute URI is reduction of the payload size, which is better achievable by following the recommendation to use gzip compression

### MUST NOT use link headers with JSON entities

For flexibility and precision, we prefer links to be directly embedded in the JSON payload instead of being attached using the uncommon link header syntax. As a result, the use of the Link Header defined by RFC 8288 in conjunction with JSON media types is forbidden.

### SHOULD NOT turn into HATEOAS

HATEOAS = links in response

- Don't do it. The premise is that clients can discover and build a ui / app based
on the links provided but there are no clients that do this. FUll HATEOAS is just overhead.
Does not mean we can't include links in the response if our UI will use them.
- [Why HATEOAS is useless and what that means for REST](https://www.reddit.com/r/programming/comments/80bul4/why_hateoas_is_useless_and_what_that_means_for/)

## 13. REST Design - Performance

### PENDING SHOULD reduce bandwidth needs and improve responsiveness

Common techniques include:
- compression of request and response bodies (see SHOULD use gzip compression)
- querying field filters to retrieve a subset of resource attributes (see SHOULD support partial responses via filtering below)
- ETag and If-Match/If-None-Match headers to avoid re-fetching of unchanged resources (see MAY consider to support ETag together with If-Match/If-None-Match header)
- Prefer header with return=minimal or respond-async to anticipate reduced processing requirements of clients (see MAY consider to support Prefer header to handle processing preferences)
- REST Design - Pagination for incremental access of larger collections of data items
- caching of master data items, i.e. resources that change rarely or not at all after creation (see MUST document cacheable GET, HEAD, and POST endpoints).

### PENDING SHOULD use gzip compression

### PENDING SHOULD support partial responses via filtering

### PENDING SHOULD allow optional embedding of sub-resources

### PENDING MUST document cacheable GET, HEAD, and POST endpoints

## Notes

- [OpenAPI specification](https://github.com/OAI/OpenAPI-Specification/)
  - [OpenAPI Map](https://openapi-map.apihandyman.io/)
- [Zalando: References](https://opensource.zalando.com/restful-api-guidelines/#dissertations)

## TODO

- [11. REST Basics - HTTP headers](https://opensource.zalando.com/restful-api-guidelines/#headers)
  - some stuff about language and location I might want to come back to in here.
- [MUST provide API identifiers [215]](https://opensource.zalando.com/restful-api-guidelines/#215)
- [optimistic-locking](https://opensource.zalando.com/restful-api-guidelines/#optimistic-locking)

