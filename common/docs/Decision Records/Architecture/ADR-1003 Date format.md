# ADR-100
# TSDR-100

## Status

Accepted

## Context

There are a many different date formats around the world. 

## Decision

Dates will always be presented in YYYY-MMM-DD format. Leading 0s on DD.  YEAR-MM-DD is allowable for data entry.

## Why / Notes

MM-DD-YY, DD-MM-YY, and YY-MM-DD can all lead to confusion and mistakes when people come from different places.

Leading 0s on DD will ensure dates don't have a variable width in columns.

## Consequences



## Other Options

Possibilities:

Not an option:

