# ADR-100
# TSDR-100

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context

telemetry is different than monitoring

https://www.reddit.com/r/devops/comments/1jcym3x/k8s_monitoring_costs_is_exploding_at_my_startup/

https://www.reddit.com/user/honeycombio/comments/1j58dxy/why_do_observability_costs_keep_exploding_lets/?p=1&impressionid=4817690181194853274

Grafana monitoring

Grafana Cloud is the easiest way to get started with metrics, logs, traces, dashboards, and more. We have a generous forever-free tier and plans for every use case.

https://www.reddit.com/r/golang/comments/1iw07rm/what_is_your_logging_monitoring_observability/

## Decision

### How long to keep logs. EVALUATE
Podcast was recommending as little as 4 days (long weekend). 7 days, 14 days or 30 days. Don't need forever. Start small first. Only lengthen for business need.

Advice was to send all logs to standard out (standard error?) and then use a different tool to send standard out on to where ever to use logs. Research why that advice was given

Which tool to use monitoring logs?

https://go.dev/wiki/Resources-for-slog#log-sinks


## Why / Notes



## Consequences



## Other Options

Possibilities:

Not an option:

