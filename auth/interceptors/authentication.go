package interceptors

import (
	"context"
	"strings"

	"github.com/ruancaetano/grpc-graphql-store/auth/utils"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryAuthServerInterceptor(mappedRoutes map[string][]string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if mappedRoutes[info.FullMethod] == nil {
			return nil, status.Errorf(codes.Unimplemented, "this route authotization config was not mapped at server auth interceptor")
		}

		if slices.Contains(mappedRoutes[info.FullMethod], "public") {
			return handler(ctx, req)
		}

		// Get the metadata from the incoming context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "couldn't parse incoming context metadata")
		}

		authenticationHeader := md.Get("authorization")

		if len(authenticationHeader) == 0 && mappedRoutes[info.FullMethod][0] != "public" {
			return nil, status.Errorf(codes.Unauthenticated, "Authorization token is required")
		}

		token := strings.TrimPrefix(authenticationHeader[0], "Bearer ")

		decodedToken, err := utils.ValidateJwt(token, mappedRoutes[info.FullMethod])
		if err != nil {
			return nil, err
		}

		ctx = metadata.AppendToOutgoingContext(ctx, "userId", decodedToken.Id)

		return handler(ctx, req)
	}
}

func UnaryAuthClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// Get the metadata from the incoming context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return status.Errorf(codes.Unauthenticated, "couldn't parse incoming context metadata sad")
		}

		authenticationHeader := md.Get("authorization")

		if len(authenticationHeader) == 0 {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", authenticationHeader[0])

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
