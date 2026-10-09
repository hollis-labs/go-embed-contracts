# go-embed-contracts

## Moved to substrate

This standalone repository is deprecated. New development lives in the
[`github.com/hollis-labs/substrate/llm-core`](https://github.com/hollis-labs/substrate/tree/llm-core/v0.1.0/llm-core)
module, released as **`llm-core/v0.1.0`**.

```sh
go get github.com/hollis-labs/substrate/llm-core@v0.1.0
```

Follow the [package and API migration guide](https://github.com/hollis-labs/substrate/blob/llm-core/v0.1.0/llm-core/embedcontracts/MIGRATION.md) when updating imports;
the consolidation can include API changes. Existing standalone tags and history
are preserved. The documentation below describes the standalone releases and
is retained for historical reference. Applications migrate separately; this
redirect does not deploy or update any consumer.

[![Go Reference](https://pkg.go.dev/badge/github.com/hollis-labs/go-embed-contracts.svg)](https://pkg.go.dev/github.com/hollis-labs/go-embed-contracts)

Shared text-embedding interface for Go. Contracts only — no provider
implementations, no runtime dependencies beyond the Go standard library.

This module exists so that libraries which produce embeddings and libraries
which consume them can depend on the same interface without pulling in each
other's transitive dependencies.

> **Status:** pre-1.0. The API is small and stable in intent, but minor-version
> bumps may still adjust shapes until v1.0.0.

## Install

```
go get github.com/hollis-labs/go-embed-contracts
```

## Quickstart

Implement the `Embedder` interface on your provider type:

```go
package mine

import (
    "context"

    embedcontracts "github.com/hollis-labs/go-embed-contracts"
)

type MyEmbedder struct{ /* client, config, etc. */ }

func (e *MyEmbedder) Embed(ctx context.Context, text, model string) (*embedcontracts.EmbeddingResult, error) {
    // call your provider, return one vector
    return &embedcontracts.EmbeddingResult{Embedding: vec, TokenCount: tokens}, nil
}

func (e *MyEmbedder) EmbedBatch(ctx context.Context, texts []string, model string) ([]embedcontracts.EmbeddingResult, error) {
    // single API call when the provider supports native batching;
    // otherwise loop over Embed.
}

func (e *MyEmbedder) EmbeddingDimensions(model string) int {
    // synchronous, no I/O. Return 0 for unknown models.
    return 1536
}
```

A complete runnable example lives in [`examples/inmemory`](examples/inmemory/main.go).

## Surface

- `Embedder` interface — `Embed`, `EmbedBatch`, `EmbeddingDimensions`
- `EmbeddingResult` struct — vector + token count

See the godoc for the contractual guarantees implementers commit to.

## Companion module

Chat-model contracts and rate-budget primitives live in
[`github.com/hollis-labs/go-llm-contracts`](https://github.com/hollis-labs/go-llm-contracts).

## Contributing

Issues and pull requests welcome. The interface is intentionally minimal;
proposed additions should be motivated by a real consumer need and should not
introduce dependencies beyond the standard library.

## License

MIT — see [LICENSE](LICENSE).
