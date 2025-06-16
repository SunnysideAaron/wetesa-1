# ADR-1019 API Design

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context



## Decision



## Why / Notes

See ADR-016 API Standard.md for initial API design / standards.

- filtering and sorting. we don't need to allow for all fields to be sortable or filterable.
Only the ones the UI needs.



- API is not the database model. Do not mirror DB tables.
  - [Data model vs. API: what’s the difference?](https://tyk.io/blog/your-data-model-is-not-an-api/)
    - "data model relates to how information is stored and retrieved. An API, on the other hand, relates to the way in which consumers experience your app."
  - [Design question on structs and responses for a REST API](https://www.reddit.com/r/golang/comments/pf4vjv/design_question_on_structs_and_responses_for_a/)
  - [Common Anti-Patterns in Go Web Applications](https://threedots.tech/post/common-anti-patterns-in-go-web-applications/#a-single-model-couples-your-application)
  - [Backend API design principles: Don’t mirror your data](https://ravendb.net/articles/backend-api-design-principles-dont-mirror-your-data)

- success and messages in response.
  - [Why do so many standards for JSON API response formats contain a "success" property in the response body instead of just using HTTP status codes?](https://softwareengineering.stackexchange.com/questions/437529/why-do-so-many-standards-for-json-api-response-formats-contain-a-success-prope)

- Provide Next and Previous Links for pagination. because of how we are doing pagination.
  either the api server creates the links or the clients do. Might as well have the api server
  do it. Less duplicated code between web, mobile and 3rd parties.
  - [json:api](https://jsonapi.org/) example links

- HATEOAS (links in response)
  - Don't do it. The premise is that clients can discover and build a ui / app based
  on the links provided but there are no clients that do this. It's just overhead.
  - [Why HATEOAS is useless and what that means for REST](https://www.reddit.com/r/programming/comments/80bul4/why_hateoas_is_useless_and_what_that_means_for/)

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


per_page


General Examples
- [Best Practices for Structuring JSON API Responses](https://sahinur.medium.com/best-practices-for-structuring-json-api-responses-24881e7add2f)
- [json:api](https://jsonapi.org/)
  - link heavy. HATEOAS?
- [PokeAPI](https://pokeapi.co/docs/v2#info)
- [Github](https://docs.github.com/en/rest?apiVersion=2022-11-28)
- [json-server](https://github.com/typicode/json-server)


## Consequences


## Other Possible Options


## Not an Option

