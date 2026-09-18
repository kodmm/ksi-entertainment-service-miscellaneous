// Package command はCQRSの書き込み側（コマンド）ハンドラーを置く。
package command

import (
	"context"
	"errors"
	"time"

	"github.com/kodmm/ksi-entertainment-service-miscellaneous/internal/domain"
)

// CreateLiveRecordInput は LiveRecord 作成コマンドの入力。
type CreateLiveRecordInput struct {
	ID        string
	UserID    string
	EventName string
	Venue     string
	LiveDate  time.Time
}

// Store はコマンド側（書き込み）が必要とする永続化のポート。
// 実装は internal/infrastructure に置く。
type Store interface {
	Save(ctx context.Context, record domain.LiveRecord) error
}

// CreateLiveRecordHandler は LiveRecord 作成コマンドを処理する。
type CreateLiveRecordHandler struct {
	store Store
}

// NewCreateLiveRecordHandler は store に依存する CreateLiveRecordHandler を組み立てる。
func NewCreateLiveRecordHandler(store Store) *CreateLiveRecordHandler {
	return &CreateLiveRecordHandler{store: store}
}

// Handle は LiveRecord を作成して保存する。
func (h *CreateLiveRecordHandler) Handle(ctx context.Context, input CreateLiveRecordInput) (domain.LiveRecord, error) {
	if input.ID == "" {
		return domain.LiveRecord{}, errors.New("command: id is required")
	}

	record := domain.LiveRecord{
		ID:        input.ID,
		UserID:    input.UserID,
		EventName: input.EventName,
		Venue:     input.Venue,
		LiveDate:  input.LiveDate,
		CreatedAt: time.Now(),
	}

	if err := h.store.Save(ctx, record); err != nil {
		return domain.LiveRecord{}, err
	}
	return record, nil
}
