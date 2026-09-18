// ksi-entertainment-service-miscellaneous のエントリーポイント。
//
// 現時点では gRPC Health Checking Protocol のみを登録した最小構成。
// session_service.proto から生成される SessionService や、将来の
// LiveRecordService の実装はフォローアップで結線する。
package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

const defaultPort = "50051"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	server := grpc.NewServer()

	healthServer := health.NewServer()
	// 個別サービス名を指定しないと全体としての SERVING しか返せないため、
	// ここでは overall status のみ SERVING にしている。個別サービスの
	// 登録は実際の RPC を実装するタイミングで追加する。
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(server, healthServer)

	log.Printf("gRPC server listening on :%s", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
