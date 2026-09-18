// Package domain は ksi-entertainment のコアとなるエンティティを定義する。
// 外部（gRPC・DB・他サービス）への依存を持たない。
package domain

import "time"

// LiveRecord はユーザーが参加した音楽ライブの参戦記録を表すエンティティ。
type LiveRecord struct {
	ID        string
	UserID    string
	EventName string
	Venue     string
	LiveDate  time.Time
	CreatedAt time.Time
}
