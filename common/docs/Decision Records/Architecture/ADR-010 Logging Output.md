# ADR-006 Logging Output

## Status

Accepted

## Context

Where to send logs

## Decision

log to STDOUT

## Why

We will log to STDOUT not to a file. A different tool will send the logs to where
ever they need to go. Docker has means to see the last logs if needed in case of
needing to debug server crashes.

Is slog asynchronous? Not a problem since we are logging to STDOUT. If we were
logging to a file we would need to find out and code a solution if not.

## Notes

- TODO https://www.reddit.com/r/golang/comments/1kcmdy7/whats_your_logging_strategy_for_go_backend/
- https://www.reddit.com/r/webdev/comments/1kdikta/is_it_good_practice_to_log_every_single_api/
  - places to send logs for storage

## Consequences
