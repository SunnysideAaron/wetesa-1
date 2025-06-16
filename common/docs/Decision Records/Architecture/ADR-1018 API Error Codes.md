# ADR-020 API Error Codes

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context



## Decision

API will return the following status codes:

### 200 OK

Standard response for successful HTTP requests. The actual response will depend on the request method used. In a GET request, the response will contain an entity corresponding to the requested resource. In a POST request, the response will contain an entity describing or containing the result of the action.

All good! Just be careful to only provide this status code if everything actually is OK. Facebook’s Graph API has been known to provide a 200 code whenever it successfully returns some output, even if that output contains an error code.

### 201 Created

The request has been fulfilled, resulting in the creation of a new resource.

### 202 Accepted

The request has been accepted for processing, but the processing has not been completed. The request might or might not be eventually acted upon, and may be disallowed when processing occurs.
    
### 3XX Redirection

**TODO** I'll have to look up which codes make sense at the time. Probably 303 and 307 but not 100% certain today.

### 400 Bad Request

This is almost always due to a typo in your user’s input. But that doesn’t mean you’re off the hook! Make sure your error message provides some specifics about the faulty input so the user can quickly fix it.

### 401 Unauthorized

This status means the input is fine but the users’ request is missing an authorization code. Not to be confused with…

### 403 Forbidden

This means the authorization code is recognized as valid but the user doesn’t have permission. For example, a user could be trying to access something only available to the admins, an increasing security concern with remote staff.

### 404 Not Found

The user’s request is valid but the endpoint or resource they’re asking for doesn’t exist. This might be because the file has since been deleted, but make sure this isn’t caused by an HTTP/HTTPS error.

### 429 Too Many Requests

rate limiting

### 500 Internal Server Error

## Why / Notes

Seems like it might be a good idea to limit the number of error codes we use.

- [List of HTTP status codes](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes)
- [6 Frequently Occurring API Errors And How to Prevent Them From Happening](https://www.astera.com/type/blog/api-errors/)
- [RESTful API Best Practices and Common Pitfalls](https://medium.com/@schneidenbach/restful-api-best-practices-and-common-pitfalls-7a83ba3763b5)

## Consequences



## Other Possible Options



## Not an Option

Don't use 422 code for validation. Simpler. Doesn't matter unless we have consumers that can tell the difference.
  - [400 vs 422 for validation](https://www.reddit.com/r/rest/comments/iv45gj/400_vs_422_for_validation/)