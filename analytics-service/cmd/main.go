package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"analytics-service/internal/grpcserver"
	"analytics-service/internal/handler"
	"analytics-service/internal/storage"
	"analytics-service/proto"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "5432"
	}
	store, err := storage.NewStorage("postgres://user:password@" + dbHost + ":" + dbPort + "/dbname?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// gRPC
	grpcSrv := grpc.NewServer()
	proto.RegisterAnalyticsServer(grpcSrv, &grpcserver.Server{Store: store})
	go func() {
		lis, _ := net.Listen("tcp", ":50051")
		log.Println("gRPC on :50051")
		grpcSrv.Serve(lis)
	}()

	// HTTP
	h := &handler.Handler{Store: store}
	r := mux.NewRouter()
	r.HandleFunc("/stats/{code}", h.GetStats).Methods("GET")
	log.Println("HTTP on :8081")
	http.ListenAndServe(":8081", r)
}
