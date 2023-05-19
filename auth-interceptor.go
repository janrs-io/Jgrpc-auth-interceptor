package Jgrpc_auth_interceptor

import (
	"context"
)

// AuthenticationInterceptor 权限验证拦截器
func AuthenticationInterceptor(ctx context.Context) (context.Context, error) {
	// 该验证使用 AuthFuncOverride 重新实现
	return ctx, nil
}
