# EmbeddedStructFieldCheck

[![Go](https://github.com/manuelarte/embeddedstructfieldcheck/actions/workflows/go.yml/badge.svg)](https://github.com/manuelarte/embeddedstructfieldcheck/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/manuelarte/embeddedstructfieldcheck)](https://goreportcard.com/report/github.com/manuelarte/embeddedstructfieldcheck)
![version](https://img.shields.io/github/v/release/manuelarte/embeddedstructfieldcheck)

Linter that checks that embedded types should be at the top of the field list of a struct.
And there must be an empty line separating embedded fields from regular fields.

<table>
<thead><tr><th>❌ Bad</th><th>✅ Good</th></tr></thead>
<tbody>
<tr><td>

```go
type Client struct {
  version int
  http.Client
}
```

</td><td>

```go
type Client struct {
  http.Client

  version int
}
```

</td></tr>

</tbody>
</table>

## ⬇️  Getting Started

### As a golangci-lint linter

Enable the linter in your golangci-lint configuration file, e.g:

```yaml
linters:
  enable:
    - embeddedstructfieldcheck 
    ...

  settings:
    embeddedstructfieldcheck:
      # Checks that sync.Mutex is not used as embedded field.
      # Default: false
      forbid-mutex: true
```

### Standalone application

Install EmbeddedStructFieldCheck by running:

```bash
go install github.com/manuelarte/embeddedstructfieldcheck@latest
```

And then use it as:

```bash
embeddedstructfieldcheck [-forbid-mutex=true|false] [--fix]
```

- `forbid-mutex`: `true|false` (default `false`) Checks that `sync.Mutex` is not used as an embedded field.
- `fix`: Fix the case when there is no space between the embedded fields and the regular fields.

## Resources

- [Embedding in structs](https://github.com/uber-go/guide/blob/master/style.md#embedding-in-structs)
