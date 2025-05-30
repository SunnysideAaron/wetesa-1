# ADR-1016 App Struct vs Parameters

## Status

Accepted

## Context

It is a commonly spotted pattern is to have a global app struct that contains all the
"global" variables and then pass that around. or make the functions methods on that struct.
This may include:
- context
- logger
- database
- config
- template cache

## Decision

This project wont use a global app struct. Passing each "global" variable as a
parameter to the functions that need them.

## Why / Notes

Passing one struct does simplify function calls. However, it also tightly couples
all the code together. Most functions don't need all the "globals". Most will just
use a subset. A global struct also means that the whole thing has to be created
in order to test any of the code. By passing parameters only what is needed is
passed.

Go specifically discourages the use of passing context in a struct. As it hides
the duration of the context. Other reasons?

TODO better write up.

- "Lets Go!" by Alex Edwards
- [Organizing Database Access](https://www.alexedwards.net/blog/organising-database-access)

## Consequences


