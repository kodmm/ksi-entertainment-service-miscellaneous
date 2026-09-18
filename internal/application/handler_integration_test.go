// application 層の command / query ハンドラーが、実際のInMemoryStoreを介して
// 正しく連携することを確認する統合テスト。
package application_test

import (
	"context"
	"testing"
	"time"

	"github.com/kodmm/ksi-entertainment-service-miscellaneous/internal/application/command"
	"github.com/kodmm/ksi-entertainment-service-miscellaneous/internal/application/query"
	"github.com/kodmm/ksi-entertainment-service-miscellaneous/internal/infrastructure"
)

func TestCreateThenListLiveRecord(t *testing.T) {
	store := infrastructure.NewInMemoryStore()
	createHandler := command.NewCreateLiveRecordHandler(store)
	listHandler := query.NewListLiveRecordsHandler(store)
	ctx := context.Background()

	created, err := createHandler.Handle(ctx, command.CreateLiveRecordInput{
		ID:        "test-1",
		UserID:    "user-1",
		EventName: "Test Live",
		Venue:     "Test Hall",
		LiveDate:  time.Now(),
	})
	if err != nil {
		t.Fatalf("createHandler.Handle() error = %v", err)
	}

	records, err := listHandler.Handle(ctx)
	if err != nil {
		t.Fatalf("listHandler.Handle() error = %v", err)
	}

	if len(records) != 1 {
		t.Fatalf("len(records) = %d, want 1", len(records))
	}
	if records[0].ID != created.ID {
		t.Fatalf("records[0].ID = %q, want %q", records[0].ID, created.ID)
	}
}

func TestCreateLiveRecordRequiresID(t *testing.T) {
	store := infrastructure.NewInMemoryStore()
	createHandler := command.NewCreateLiveRecordHandler(store)

	_, err := createHandler.Handle(context.Background(), command.CreateLiveRecordInput{})
	if err == nil {
		t.Fatal("createHandler.Handle() error = nil, want error for missing ID")
	}
}
