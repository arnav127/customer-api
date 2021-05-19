package main

import (
	"context"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"gitlab.com/arnavdixit/customer-api/proto"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwmux, err := newGateway(ctx)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	log.Println("grpc-gateway listen on localhost:8888")
	return http.ListenAndServe(":8888", mux)
}

func newGateway(ctx context.Context) (http.Handler, error) {

	opts := []grpc.DialOption{grpc.WithInsecure()}

	gwmux := runtime.NewServeMux()
	if err := proto.RegisterCustomerServiceHandlerFromEndpoint(ctx, gwmux, ":8080", opts); err != nil {
		return nil, err
	}

	return gwmux, nil
}

func main() {
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
