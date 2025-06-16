# ADR-1018 API Guidelines

## 1. Introduction

See ADR-016 API Standard.md for initial API design / standards.

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

### It is OK to be different.

In "Design and Build Great Web APIs" by Mike Amundsen it is stated that givin the
same spec 5 api designers would create 5 different APIs. I'm getting the impression
that APIs are like UI and Resumes. Who ever you talk to will be able to point out
how you did it wrong. Here is my permission to flip them the bird and just do what
makes sense for this problem.

### API is not the database model.

Do not mirror DB tables. Because this example is so small this might be a bit hard to see in this
example.

- [Data model vs. API: what’s the difference?](https://tyk.io/blog/your-data-model-is-not-an-api/)
- "data model relates to how information is stored and retrieved. An API, on the other hand, relates to the way in which consumers experience your app."
- [Design question on structs and responses for a REST API](https://www.reddit.com/r/golang/comments/pf4vjv/design_question_on_structs_and_responses_for_a/)
- [Common Anti-Patterns in Go Web Applications](https://threedots.tech/post/common-anti-patterns-in-go-web-applications/#a-single-model-couples-your-application)
- [Backend API design principles: Don’t mirror your data](https://ravendb.net/articles/backend-api-design-principles-dont-mirror-your-data)

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

### PENDING MUST define collection format of header and query parameters

Pick one of the following. I prefer the first. But do we have to escape commas?

```
?param=value1,value2
```	

```
?param=value1&param=value2
```

### PENDING MUST use URL parameters for result set limiting

- **fields**: field name expression to retrieve only a subset of fields of a resource. See SHOULD support partial responses via filtering below.
  - We don't need to allow for all fields to be filterable. Only the ones the UI needs.

- **embed**: field name expression to expand or embedded sub-entities, e.g. inside of an article entity, expand silhouette code into the silhouette object. Implementing embed correctly is difficult, so do it with care. See SHOULD allow optional embedding of sub-resources below.

### PENDING MUST use URL parameters for sorting

- **sort**: comma-separated list of fields (as defined by MUST define collection format of header and query parameters) to define the sort order. To indicate sorting direction, fields may be prefixed with + (ascending) or - (descending), e.g. /sales-orders?sort=+id.
  - We don't need to allow for all fields to be sortable. Only the ones the UI needs.



### PENDING MUST use URL parameters for queries

- **q**: default query parameter, e.g. used by browser tab completion; should have an entity specific alias, e.g. sku.

- **filter**: comma-separated list of fields (as defined by MUST define collection format of header and query parameters) to define the filter criteria. To indicate filtering direction, fields may be prefixed with + (include) or - (exclude), e.g. /sales-orders?filter=-cancelled.

- [SHOULD design simple query languages using query parameters [236]](https://opensource.zalando.com/restful-api-guidelines/#236)


- URL Formatting

  - [FIQL](https://datatracker.ietf.org/doc/html/draft-nottingham-atompub-fiql-00)
    - since 2007
  - [RSQL]()
    - slightly more modern version of FIQL
    - [Here: RSQL](https://www.here.com/docs/bundle/data-client-library-developer-guide-java-scala/page/client/rsql.html)  
  - [OData](https://www.odata.org/)
    - came from microsoft
    - [OData adoption rate?](https://www.reddit.com/r/dotnet/comments/11eoa6d/odata_adoption_rate/)
    - maybe too flexible?
  - [speakeasy: Filtering Collections](https://www.speakeasy.com/api-design/filtering-responses)
  - [Correct way to pass multiple values for same parameter name in GET request](https://stackoverflow.com/questions/24059773/correct-way-to-pass-multiple-values-for-same-parameter-name-in-get-request)
- [json-server](https://github.com/typicode/json-server)
  - example


### MUST use URL query parameters for pagination

SEE ADR-1019 API Pagination.md

### MUST NOT design complex query languages using JSON

For our example API we wont need complex queries. It's enough to know we can.




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

### MUST use success and messages in response

- [Why do so many standards for JSON API response formats contain a "success" property in the response body instead of just using HTTP status codes?](https://softwareengineering.stackexchange.com/questions/437529/why-do-so-many-standards-for-json-api-response-formats-contain-a-success-prope)




## 9. REST Basics - HTTP requests

### MUST use POST for create

Do not use PUT. Clients will not create IDs

### PENDING update methods

- PENDING PUT vs PATCH use one throughout unless have specific use case.
- TODO: https://opensource.zalando.com/restful-api-guidelines/#http-requests

### TODO

- idempotent
  - same request results in same response / action. ie delete deletes once but subsequent calls 
  - still say deleted. For how long? before not not found?
  - creates once. don't create multiple of same record?
  - update once. don't keep overwriting?
  - ### PENDING SHOULD use secondary key for idempotent POST design
  - ### PENDING SHOULD consider to design POST and PATCH idempotent
  - Idempotency: PUT requests are generally idempotent, meaning that multiple identical PUT requests should have the same effect as a single request. POST requests are not idempotent, as each request can result in the creation of a new resource.
- caching
- support asynchronous request processing
  - MAY support asynchronous request processing
- implicit response filtering
  - [MUST document implicit response filtering [226]](https://opensource.zalando.com/restful-api-guidelines/#226)

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

- https://www.vinaysahni.com/best-practices-for-a-pragmatic-restful-api
- https://github.com/interagent/http-api-design
- [Ask HN: What are good reads for designing APIs?](https://news.ycombinator.com/item?id=12262586)
- https://stackoverflow.com/questions/978061/http-get-with-request-body


- [Microsoft REST API Guidelines](https://news.ycombinator.com/item?id=12122828)
- [OpenAPI specification](https://github.com/OAI/OpenAPI-Specification/)
  - [OpenAPI Map](https://openapi-map.apihandyman.io/)
- [Zalando: References](https://opensource.zalando.com/restful-api-guidelines/#dissertations)
- "Design and Build Great Web APIs" by Mike Amundsen
  - Aaron notes: seemed a bit too basic? or maybe a bit to much on process of documenting?
    Something seemed off. Read several chapters.
- https://amberonrails.com/building-stripes-api


## TODO

- [11. REST Basics - HTTP headers](https://opensource.zalando.com/restful-api-guidelines/#headers)
  - some stuff about language and location I might want to come back to in here.
- [MUST provide API identifiers [215]](https://opensource.zalando.com/restful-api-guidelines/#215)
- [optimistic-locking](https://opensource.zalando.com/restful-api-guidelines/#optimistic-locking)

## Example APIs
- [Best Practices for Structuring JSON API Responses](https://sahinur.medium.com/best-practices-for-structuring-json-api-responses-24881e7add2f)
- [json:api](https://jsonapi.org/)
  - link heavy. HATEOAS?
- [PokeAPI](https://pokeapi.co/docs/v2#info)
- [Github](https://docs.github.com/en/rest?apiVersion=2022-11-28)
- [json-server](https://github.com/typicode/json-server)
- [Todd Motto: public-apis](https://github.com/toddmotto/public-apis)
- https://docs.stripe.com/api stripe api is often recommended as a good api