# Changelog

## Repository retirement — 2026-10-09

- Deprecated this standalone repository in favor of `github.com/hollis-labs/substrate/llm-core@v0.1.0`
  ([migration guide](https://github.com/hollis-labs/substrate/blob/llm-core/v0.1.0/llm-core/embedcontracts/MIGRATION.md)).
- Preserved existing release tags and history. This documentation change does
  not create a new standalone release or migrate applications.

All notable changes to this project will be documented in this file. The
format is loosely based on [Keep a Changelog](https://keepachangelog.com/),
and this project adheres to [Semantic Versioning](https://semver.org/).

## v0.1.1 — 2026-05-10

### Added

- `examples/inmemory/` — runnable example showing how to satisfy the
  `Embedder` interface end-to-end (single, batch, dimension lookup).

### Changed

- README rewritten for public consumption: install snippet, quickstart code
  block, godoc badge, status banner, contributing pointer, license line.
- CHANGELOG reformatted with Keep-a-Changelog headings.

No public API changes in this release.

## v0.1.0 — 2026-05-09

### Added

- `Embedder` interface — `Embed`, `EmbedBatch`, `EmbeddingDimensions`.
- `EmbeddingResult` struct — embedding vector and token count.
