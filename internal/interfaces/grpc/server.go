// Package grpcserver は gRPC サーバーの構築・サービス登録を行う interfaces 層。
package grpcserver

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/kodmm/ksi-entertainment-service-miscellaneous/internal/application/command"
	"github.com/kodmm/ksi-entertainment-service-miscellaneous/internal/application/query"
)

// Server はこのサービスが公開する gRPC サーバーをラップする。
type Server struct {
	grpcServer *grpc.Server

	// createLiveRecord / listLiveRecords は今のところ未使用（LiveRecordService の
	// RPC 登録は proto 生成コードの取り込み後に行う。フォローアップ: Issue #2）。
	// コンポジションルート（main.go）から結線だけ先に済ませておく。
	createLiveRecord *command.CreateLiveRecordHandler
	listLiveRecords  *query.ListLiveRecordsHandler
}

// New は gRPC サーバーを組み立て、Health Checking Protocol を登録する。
func New(createLiveRecord *command.CreateLiveRecordHandler, listLiveRecords *query.ListLiveRecordsHandler) *Server {
	grpcServer := grpc.NewServer()

	healthServer := health.NewServer()
	// 個別サービス名を指定しないと全体としての SERVING しか返せないため、
	// ここでは overall status のみ SERVING にしている。個別サービスの
	// 登録は実際の RPC を実装するタイミングで追加する。
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	return &Server{
		grpcServer:       grpcServer,
		createLiveRecord: createLiveRecord,
		listLiveRecords:  listLiveRecords,
	}
}

// Serve はリスナー上で gRPC サーバーを起動する。
func (s *Server) Serve(lis net.Listener) error {
	return s.grpcServer.Serve(lis)
}
