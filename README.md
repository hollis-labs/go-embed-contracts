# go-embed-contracts

Text-embedding interface for the hollis-labs portfolio. Contracts-only.

## Surface

- `Embedder` interface
- `EmbeddingResult` struct

## Consumers

- `vanta-conduit` — primary consumer (memory store, recall, mcpadapter)
- `stack-explorer` — code-embed indexing
- `nanite` — chat-side embedder selection (and a future SDK wrapper at
  `internal/llm/openai/embed.go` per SP-20260508-0001 CW-0012)

## Lineage

Originally in `github.com/hollis-labs/go-providers` (`provider.Embedder`).
v0.10.0 of go-providers narrowed scope to CLI/PTY/subprocess-only and
removed the HTTP-API surface; the embedding contract moved here.

## Companion module

Chat contracts and rate-budget primitives live in
`github.com/hollis-labs/go-llm-contracts`.
