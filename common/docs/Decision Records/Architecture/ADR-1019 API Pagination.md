# ADR-100
# TSDR-100

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context



## Decision

- Use offset pagination.
- Do not return total records

## Why

Even if our data sets gets very large. For crud based apps
pages will usually be less than 10 pages of usable data for end users. Otherwise data
will be filtered, a report, or a data export. No user is going to click "Next page" 210 times.
Perhaps social media type data sets might need that. ie scrolling through a feed.

If we end up with a use case for cursor based pagination we can implement it then.

Total records is expensive to calculate. Have to run the query twice. Once for
the data and once for the count.

## Notes

offset vs cursor
offset 
  pro: is easier to implement
  pro: allows to jump to specific pages. 
  con: prone to drifting. ie if data is deleted or added while user goes between pages.
  con: the deeper into the pages the user goes the more data has to be read. then thrown away
       by the db.
  look into differed join if large data sets is an issue.
cursor 
  pro: is better for performance on large data sets
  pro: more resistant to drifting.
  con: can't jump to specific page. have to go through all previous pages to get to desired page.
  con: code to implement is more complex
  con: have to pass around cursors/tokens


## Consequences

Since we aren't returning total records the UI will not have specific pages for
the user to jump to. ie not page 1, 2, 3, ... 10, 11, 12 links. Users will have
to use filtering to quickly get to pages they are looking for.

## Other Possible Options
[Pagination in MySQL - offset vs. cursor](https://www.youtube.com/watch?v=zwDIN04lIpc)
  - [Pagination in MySQL](https://planetscale.com/blog/mysql-pagination)
  - https://planetscale.com/learn/courses/mysql-for-developers/examples/deferred-joins
  - [Efficient MySQL pagination using deferred joins](https://aaronfrancis.com/2022/efficient-mysql-pagination-using-deferred-joins-15d0de14)

https://www.depesz.com/2011/05/20/pagination-with-fixed-order/
https://dev.to/sadhakbj/implementing-cursor-pagination-in-golang-go-fiber-mysql-gorm-from-scratch-2p60

## Not an Option

