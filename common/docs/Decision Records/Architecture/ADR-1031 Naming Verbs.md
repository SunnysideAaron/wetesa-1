# ADR-1031 Naming Verbs

## Status

Accepted

## Decision

In code procedure names use Get [sigular], Get [plural], Post, Put, Patch, Delete for Verbs.

API will be verb noun. Web will be noun verb.

## Why / Notes

Don't use Read, List, Create, Update, Delete or other verbs unless we are doing something unique.

We don't have any control over HTTP method names. Might as well use them and keep it consistent. No need to translate.
End users do not have to see HTTP method names. The UI can and should use whatever makes sense.

Swapping verb and noun order between web and api handlers is an attempt to make sure
they don't have the same name and hopefully avoid confusion.