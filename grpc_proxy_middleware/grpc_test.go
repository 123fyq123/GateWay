package grpc_proxy_middleware

import (
	"log"

	"google.golang.org/grpc"
)

// 流式RPC拦截器
func GrpcAuthStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	err := handler(srv, ss)
	if err != nil {
		log.Printf("RPC failed with error %v\n", err)
	}
	return err
}
