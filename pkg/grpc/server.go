package grpc

import (
	"log/slog"
	"net"

	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type Server interface {
	grpc.ServiceRegistrar
	Listen(addr string) error
	Stop()
}

type ServerOpts struct {
	Logger  logging.Logger
	Slogger *slog.Logger
}

type grpcServer struct {
	*grpc.Server
	// Deprecated: use slogger
	logger  logging.Logger
	slogger *slog.Logger
}

func (g *grpcServer) Listen(addr string) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.NewEf(err, "could not listen to net/tcp server")
	}
	g.slogger.Info("grpc server listening", "at", addr)
	return g.Serve(listen)
}

func (g *grpcServer) Stop() {
	g.Server.GracefulStop()
}

func NewGrpcServer(opts ServerOpts) (Server, error) {
	if opts.Slogger == nil {
		opts.Slogger = slog.Default()
	}

	server := grpc.NewServer(
		grpc.StreamInterceptor(func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			p, ok := peer.FromContext(stream.Context())
			if ok {
				_ = p.Addr.String()
				// if opts.Slogger != nil {
				// 	opts.Slogger.Debug("new grpc connection", "from", p.Addr.String())
				// } else {
				// 	opts.Logger.Debugf("[Stream] New connection from %s", p.Addr.String())
				// }
			}
			return handler(srv, stream)
		}),
	)

	return &grpcServer{Server: server, logger: opts.Logger, slogger: opts.Slogger}, nil
}

// Type guard to ensure grpcServer implements Server interface, at compile time
var _ Server = &grpcServer{}
