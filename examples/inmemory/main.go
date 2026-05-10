// Package main demonstrates implementing the embedcontracts.Embedder interface
// with a tiny in-memory stub.
//
// The example shows:
//   - implementing Embed, EmbedBatch, and EmbeddingDimensions on a custom type;
//   - returning EmbeddingResult values with TokenCount populated;
//   - exercising the contract via the standard interface.
//
// Real implementations call out to a provider API (OpenAI, Voyage, a local
// model server, etc.); this stub just hashes the input into a fixed-width
// vector so the example is self-contained and runs offline.
//
// Run with:
//
//	go run ./examples/inmemory
package main

import (
	"context"
	"fmt"
	"hash/fnv"

	embedcontracts "github.com/hollis-labs/go-embed-contracts"
)

// dims is the vector width this stub embedder produces.
const dims = 8

// hashEmbedder is a deterministic offline Embedder for demonstration.
// It hashes each input string into a fixed-width float32 vector. It is not
// useful for similarity search — it exists only to show the interface shape.
type hashEmbedder struct{}

func (hashEmbedder) Embed(_ context.Context, text, _ string) (*embedcontracts.EmbeddingResult, error) {
	return &embedcontracts.EmbeddingResult{
		Embedding:  vectorize(text),
		TokenCount: len(text),
	}, nil
}

func (h hashEmbedder) EmbedBatch(ctx context.Context, texts []string, model string) ([]embedcontracts.EmbeddingResult, error) {
	out := make([]embedcontracts.EmbeddingResult, len(texts))
	for i, t := range texts {
		r, err := h.Embed(ctx, t, model)
		if err != nil {
			return nil, fmt.Errorf("embed %d: %w", i, err)
		}
		out[i] = *r
	}
	return out, nil
}

func (hashEmbedder) EmbeddingDimensions(_ string) int { return dims }

func vectorize(text string) []float32 {
	v := make([]float32, dims)
	h := fnv.New64a()
	_, _ = h.Write([]byte(text))
	sum := h.Sum64()
	for i := 0; i < dims; i++ {
		v[i] = float32((sum>>(i*8))&0xff) / 255.0
	}
	return v
}

func main() {
	// Compile-time check: confirm we satisfy the contract.
	var emb embedcontracts.Embedder = hashEmbedder{}

	ctx := context.Background()
	const model = "demo-model"

	// Single-input path.
	single, err := emb.Embed(ctx, "hello, embeddings", model)
	if err != nil {
		panic(err)
	}
	fmt.Printf("single: dims=%d tokens=%d vec[0..2]=%v\n",
		len(single.Embedding), single.TokenCount, single.Embedding[:3])

	// Batch path.
	batch, err := emb.EmbedBatch(ctx, []string{"alpha", "beta", "gamma"}, model)
	if err != nil {
		panic(err)
	}
	for i, r := range batch {
		fmt.Printf("batch[%d]: dims=%d tokens=%d\n", i, len(r.Embedding), r.TokenCount)
	}

	// Synchronous dimension lookup.
	fmt.Printf("declared dims for %q: %d\n", model, emb.EmbeddingDimensions(model))
}
