# ADR-100
# TSDR-100

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context

format the slog output for human readability.

## Decision



## Why / Notes



## Consequences




## Links
- See ADR-011 Logging Package.
- [Awesome-Go: Logging](https://github.com/avelino/awesome-go?tab=readme-ov-file#logging)
- [Go Wiki: Resources for slog](https://go.dev/wiki/Resources-for-slog)
- **[slog-multi](https://github.com/samber/slog-multi)**
  - saving in case we want to log to multiple places down the road.
- [slogd](https://github.com/kaihendry/slogd)
  - slog with duration
  - with video https://youtu.be/IsPa11N5pzI

## Possible Options (short list)

- [slog-formatter](https://github.com/samber/slog-formatter)
  - Common formatters for slog + helpers for building your own
  - seems like a lot there. but not widely used?
- [tinted](https://pkg.go.dev/github.com/lmittmann/tint)
  - colorized output

## Other Possible Options

- [ConsoleHandler](https://gist.github.com/wijayaerick/de3de10c47a79d5310968ba5ff101a19)
  - similar to Zap’s ConsoleEncoder
  - very small code
- [humane](https://github.com/telemachus/humane)
  - a human-friendly (but still largely structured) slog Handler
- [klog](https://github.com/kubernetes/klog)
  - the text format used by Kubernetes. Provides klog output routing when using the main package’s logger and a simpler logger that just writes to stderr. Both slog/logr and go-logr/logr APIs are supported.
- [slogjson](https://github.com/veqryn/slog-json)
  - Format using the upcoming **JSON v2** library, with optional single-line pretty-printing: https://github.com/veqryn/slog-json
- [slogor](https://gitlab.com/greyxor/slogor)
  - A colorful slog handler

## Not an Option

- [logf](https://pkg.go.dev/github.com/AndrewHarrisSPU/logf)
  - (attr {key} interpolation, rich tty output):  (uses lazy Handler stores: https://go.dev/play/p/psdD7KDF5fp )
  not maintained?
- [standard library slog: Writing a handler](https://pkg.go.dev/log/slog@go1.24.1#hdr-Writing_a_handler)
  - we don't need to write our own.
  - [slgotest](https://pkg.go.dev/testing/slogtest@go1.24.1)
- [slug](https://github.com/dotse/slug)
  - a handler that prints colourful logs for humans
  - not maintained