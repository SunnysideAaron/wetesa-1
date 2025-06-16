# ADR-020 API Error Codes

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context



## Decisions

- MUST use code 207 for batch or bulk requests
- MUST use code 429 with headers for rate limits
- MUST NOT use 422 code for validation.
- MUST not expose stack traces

### PENDING MUST support multiple errors in JSON response.

API will use the following status codes:

### 200 OK

Standard response for successful HTTP requests. The actual response will depend on the request method used. In a GET request, the response will contain an entity corresponding to the requested resource. In a POST request, the response will contain an entity describing or containing the result of the action.

All good! Just be careful to only provide this status code if everything actually is OK. Facebook’s Graph API has been known to provide a 200 code whenever it successfully returns some output, even if that output contains an error code.

### 201 Created

The request has been fulfilled, resulting in the creation of a new resource.

### 202 Accepted

The request has been accepted for processing, but the processing has not been completed. The request might or might not be eventually acted upon, and may be disallowed when processing occurs.

### 204 No content PENDING

Returned instead of 200, if no response payload is returned. Normally only applicable to methods which change something.

### 207 Multi-Status POST (DELETE)

The response body contains status information for multiple different parts of a batch/bulk request (see MUST use code 207 for batch or bulk requests for details). Normally used for POST, in some cases also for DELETE

### 3XX Redirection

**TODO** I'll have to look up which codes make sense at the time. Probably 303 and 307 but not 100% certain today.
- [SHOULD not use redirection codes [251]](https://opensource.zalando.com/restful-api-guidelines/#251)

Zalando recommends not using redirect codes as there are better alternatives.
- https://opensource.zalando.com/restful-api-guidelines/#251

Zalando uses a 304 instead of 200 if the resouce wasn't changed because someone else changed it first.
- https://opensource.zalando.com/restful-api-guidelines/#status-code-304

### 400 Bad Request

This is almost always due to a typo in your user’s input. But that doesn’t mean you’re off the hook! Make sure your error message provides some specifics about the faulty input so the user can quickly fix it.

### 401 Unauthorized

This status means the input is fine but the users’ request is missing an authorization code. Not to be confused with…

### 403 Forbidden

This means the authorization code is recognized as valid but the user doesn’t have permission. For example, a user could be trying to access something only available to the admins, an increasing security concern with remote staff.

### 404 Not Found

The user’s request is valid but the endpoint or resource they’re asking for doesn’t exist. This might be because the file has since been deleted, but make sure this isn’t caused by an HTTP/HTTPS error.

### 405 Method Not Allowed PENDING

The request method is not supported for this resource. In theory, this can be returned for all resources for all the methods except the ones documented. Using this response code for an existing endpoint (usually with path parameters) only makes sense if it depends on some internal resource state whether a specific method is allowed, e.g. an order can only be canceled via DELETE until the shipment leaves the warehouse. Do not use it unless you have such a special use case, but then make sure to document it, making it clear why a resource might not support a method.

### 409 Conflict PENDING

The request cannot be completed due to conflict with the current state of the target resource.
 For example, you may get a 409 response when updating a resource that is older than the existing one on the server, 
 resulting in a version control conflict. 
 
 **PENDING** If this is used, it MUST be documented. For successful robust creation of resources (PUT or POST) you should always return 200 or 204 and not 409, even if the resource exists already. If any If-* headers cause a conflict, you should use 412 and not 409. Only applicable to methods which change something.

### 410 Gone PENDING

The resource does not exist any longer. It did exist in the past, and will most likely not exist in the future. This can be used, e.g. when accessing a resource that has intentionally been deleted. This normally does not need to be documented, unless there is a specific need to distinguish this case from the normal 404.

### 422 Unprocessable Content DO NOT USE

Use 400 instead. Simpler. Doesn't matter unless we have consumers that can tell the difference.
  - [400 vs 422 for validation](https://www.reddit.com/r/rest/comments/iv45gj/400_vs_422_for_validation/)

### 429 Too Many Requests

rate limiting

### 500 Internal Server Error

specific errors should be in the JSON response body.

## Why / Notes

Seems like it might be a good idea to limit the number of error codes we use.

- [List of HTTP status codes](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes)
- [Zalando: HTTP status codes](https://opensource.zalando.com/restful-api-guidelines/#150)
- [6 Frequently Occurring API Errors And How to Prevent Them From Happening](https://www.astera.com/type/blog/api-errors/)
- [RESTful API Best Practices and Common Pitfalls](https://medium.com/@schneidenbach/restful-api-best-practices-and-common-pitfalls-7a83ba3763b5)

## Consequences


## Other Possible Options



## Not an Option

