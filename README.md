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

Install it by running:

```bash
go install github.com/manuelarte/embeddedstructfieldcheck@latest
```

## Resources

- [Embedding in structs](https://github.com/uber-go/guide/blob/master/style.md#embedding-in-structs)
