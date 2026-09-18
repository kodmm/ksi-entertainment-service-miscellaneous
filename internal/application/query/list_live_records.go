// Package query はCQRSの読み取り側（クエリ）ハンドラーを置く。
package query

import (
	"context"

	"github.com/kodmm/ksi-entertainment-service-miscellaneous/internal/domain"
)

// Store はクエリ側（読み取り）が必要とする永続化のポート。
// 実装は internal/infrastructure に置く。
type Store interface {
	List(ctx context.Context) ([]domain.LiveRecord, error)
}

// ListLiveRecordsHandler は LiveRecord 一覧を取得するクエリを処理する。
type ListLiveRecordsHandler struct {
	store Store
}

// NewListLiveRecordsHandler は store に依存する ListLiveRecordsHandler を組み立てる。
func NewListLiveRecordsHandler(store Store) *ListLiveRecordsHandler {
	return &ListLiveRecordsHandler{store: store}
}

// Handle は LiveRecord の一覧を返す。
func (h *ListLiveRecordsHandler) Handle(ctx context.Context) ([]domain.LiveRecord, error) {
	return h.store.List(ctx)
}
