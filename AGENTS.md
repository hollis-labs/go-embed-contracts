# go-embed-contracts

The shared text-embedding interface: `Embedder` plus `EmbeddingResult`, and
nothing else. It ships no provider implementations and depends on nothing
beyond the standard library, so a library that produces embeddings and one that
consumes them can share an interface without inheriting each other's
dependency graphs.

## Start Here

- `embedder.go` is the entire surface, and its doc comment states the
  guarantees implementers commit to.
- `examples/inmemory/main.go` is a complete runnable implementation.
- `README.md` points at the companion `go-llm-contracts` for chat-model
  contracts.

## Commands

```bash
make vet
make test
make lint
```

## Boundaries

Contracts only. A provider implementation, an HTTP client or any third-party
dependency landing here would defeat the module's single purpose — consumers
import this precisely because it costs them nothing transitively.

The interface doc comment is the contract, and three parts of it are easy to
break by accident: `EmbedBatch` must issue one API call only where the provider
natively batches (callers must not assume it is faster), `EmbeddingDimensions`
must be synchronous with no I/O, and it returns 0 for an unknown model rather
than erroring.

Keep additions motivated by a real consumer. Not every provider in the
portfolio embeds, so consumers type-assert for this interface rather than
expecting it.
