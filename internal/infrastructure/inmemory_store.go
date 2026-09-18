// Package infrastructure は外部リソース（DB・キャッシュ等）との連携を置く。
// 今はDBの代わりにインメモリ実装のみを提供する。
package infrastructure

import (
	"context"
	"sync"

	"github.com/kodmm/ksi-entertainment-service-miscellaneous/internal/domain"
)

// InMemoryStore は DB 未結線の間だけ使う暫定の永続化実装。
// command.Store と query.Store の両方を構造的に満たす（CQRSだが実体は1つの
// インメモリスライスを共有する簡易実装）。
type InMemoryStore struct {
	mu      sync.RWMutex
	records []domain.LiveRecord
}

// NewInMemoryStore は空の InMemoryStore を作る。
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{}
}

// Save は LiveRecord を追加する（command.Store の実装）。
func (s *InMemoryStore) Save(_ context.Context, record domain.LiveRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.records = append(s.records, record)
	return nil
}

// List は保存済みの LiveRecord をすべて返す（query.Store の実装）。
func (s *InMemoryStore) List(_ context.Context) ([]domain.LiveRecord, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]domain.LiveRecord, len(s.records))
	copy(out, s.records)
	return out, nil
}
