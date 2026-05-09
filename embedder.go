package embedcontracts

import "context"

// EmbeddingResult holds the output of an embedding request.
type EmbeddingResult struct {
	// Embedding is the vector representation of the input text.
	Embedding []float32
	// TokenCount is the number of tokens consumed (for billing/rate tracking).
	TokenCount int
}

// Embedder is implemented by providers that produce text embeddings.
//
// Implementers commit to:
//   - Single-input Embed: one round-trip per call, returning a single vector.
//   - EmbedBatch: a single API call when the underlying provider supports
//     native batching; otherwise an internal loop over Embed (callers should
//     not assume the batch path is faster than single by default).
//   - EmbeddingDimensions: synchronous lookup with no I/O. Returns 0 when the
//     model is unknown to the implementer.
//
// Consumers type-assert when needed; not every Provider in the portfolio
// supports embedding.
type Embedder interface {
	Embed(ctx context.Context, text string, model string) (*EmbeddingResult, error)
	EmbedBatch(ctx context.Context, texts []string, model string) ([]EmbeddingResult, error)
	EmbeddingDimensions(model string) int
}
