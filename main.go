// ksi-entertainment-service-miscellaneous のエントリーポイント。
//
// ここはコンポジションルート（依存関係を組み立てるだけの場所）。
// 実際のロジックは internal/ 配下の各レイヤーに置く。層構成は README.md を参照。
package main

import (
	"log"
	"net"
	"os"

	"github.com/kodmm/ksi-entertainment-service-miscellaneous/internal/application/command"
	"github.com/kodmm/ksi-entertainment-service-miscellaneous/internal/application/query"
	"github.com/kodmm/ksi-entertainment-service-miscellaneous/internal/infrastructure"
	grpcserver "github.com/kodmm/ksi-entertainment-service-miscellaneous/internal/interfaces/grpc"
)

const defaultPort = "50051"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	store := infrastructure.NewInMemoryStore()
	createLiveRecord := command.NewCreateLiveRecordHandler(store)
	listLiveRecords := query.NewListLiveRecordsHandler(store)

	server := grpcserver.New(createLiveRecord, listLiveRecords)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	log.Printf("gRPC server listening on :%s", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
