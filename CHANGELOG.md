# Changelog

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
