package Jgrpc_auth_interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Validator 实现 PGV Validator 接口
type Validator interface {
	ValidateAll() error
}

// ValidationUnaryInterceptor PGV 验证拦截器
func ValidationUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	if r, ok := req.(Validator); ok {
		if err = r.ValidateAll(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	return handler(ctx, req)

}

// AuthenticationInterceptor 权限验证拦截器
func AuthenticationInterceptor(ctx context.Context) (context.Context, error) {
	// 该验证使用 AuthFuncOverride 重新实现
	return ctx, nil
}
