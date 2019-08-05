package middle

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go/mocktracer"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"os"
)

var logsk = io.MultiWriter(os.Stdout)
var zapLogger *zap.Logger
var customFunc grpc_zap.CodeToLevel

// grpc流访问拦截器
func GetStreamServerOption() grpc.ServerOption {
	mockTracer := mocktracer.New()
	opts := []grpc_opentracing.Option{
		grpc_opentracing.WithTracer(mockTracer),
	}
	return grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		// 为上下文增加Tag map对象
		grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		// 支持opentracing/zipkin
		grpc_opentracing.StreamServerInterceptor(opts...),
		grpc_auth.StreamServerInterceptor(myAuthFunction),
	))
}

// 非grpc流访问拦截器
func GetUnaryServerOption() grpc.ServerOption {
	mockTracer := mocktracer.New()
	opts := []grpc_opentracing.Option{
		grpc_opentracing.WithTracer(mockTracer),
	}
	grpc_middleware.ChainUnaryClient()
	return grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		// 为上下文增加Tag map对象
		grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		// 支持opentracing/zipkin
		grpc_opentracing.UnaryServerInterceptor(opts...),
		grpc_auth.UnaryServerInterceptor(myAuthFunction),
	))
}

func myAuthFunction(ctx context.Context) (context.Context, error) {
	var md metadata.MD
	var user string
	var password string
	var ok bool
	if md, ok = metadata.FromIncomingContext(ctx); !ok {
		return nil, status.Errorf(codes.PermissionDenied, "沒有访问权限")
	}
	if val, ok := md["user"]; ok {
		user = val[0]
	}
	if val, ok := md["password"]; ok {
		password = val[0]
	}
	if user != "chc" || password != "123456" {
		return nil, status.Errorf(codes.PermissionDenied, "用户名或密码错误")
	}
	return ctx, nil
}
