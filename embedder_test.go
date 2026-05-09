package embedcontracts_test

import (
	"context"
	"testing"

	embedcontracts "github.com/hollis-labs/go-embed-contracts"
)

type fake struct{ vec []float32 }

func (f *fake) Embed(_ context.Context, _ string, _ string) (*embedcontracts.EmbeddingResult, error) {
	return &embedcontracts.EmbeddingResult{Embedding: f.vec, TokenCount: 1}, nil
}

func (f *fake) EmbedBatch(_ context.Context, texts []string, _ string) ([]embedcontracts.EmbeddingResult, error) {
	out := make([]embedcontracts.EmbeddingResult, len(texts))
	for i := range texts {
		out[i] = embedcontracts.EmbeddingResult{Embedding: f.vec, TokenCount: 1}
	}
	return out, nil
}

func (f *fake) EmbeddingDimensions(_ string) int { return len(f.vec) }

func TestFakeSatisfiesEmbedder(t *testing.T) {
	var _ embedcontracts.Embedder = (*fake)(nil)
}
