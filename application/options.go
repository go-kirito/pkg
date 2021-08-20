package application

import (
	"context"
	"net/url"
	"os"

	"github.com/go-kirito/pkg/registry"
	"github.com/go-kirito/pkg/transport"
	"github.com/go-kirito/pkg/transport/grpc"
	"github.com/go-kirito/pkg/transport/http"
	"github.com/go-kirito/pkg/zlog"
)

// Option is an application option.
type Option func(o *options)

// options is an application options.
type options struct {
	id        string
	name      string
	version   string
	metadata  map[string]string
	endpoints []*url.URL

	ctx  context.Context
	sigs []os.Signal

	logger    zlog.Logger
	registrar registry.Registrar

	servers    []transport.Server
	httpServer *http.Server
	grpcServer *grpc.Server
}

// ID with service id.
func ID(id string) Option {
	return func(o *options) { o.id = id }
}

// Name with service name.
func Name(name string) Option {
	return func(o *options) { o.name = name }
}

// Version with service version.
func Version(version string) Option {
	return func(o *options) { o.version = version }
}

// Metadata with service metadata.
func Metadata(md map[string]string) Option {
	return func(o *options) { o.metadata = md }
}

// Endpoint with service endpoint.
func Endpoint(endpoints ...*url.URL) Option {
	return func(o *options) { o.endpoints = endpoints }
}

// Context with service context.
func Context(ctx context.Context) Option {
	return func(o *options) { o.ctx = ctx }
}

// Logger with service logger.
func Logger(logger zlog.Logger) Option {
	return func(o *options) { o.logger = logger }
}

// Server with transport servers.
func Server(srv ...transport.Server) Option {
	return func(o *options) { o.servers = srv }
}

// Signal with exit signals.
func Signal(sigs ...os.Signal) Option {
	return func(o *options) { o.sigs = sigs }
}

// Registrar with service registry.
func Registrar(r registry.Registrar) Option {
	return func(o *options) { o.registrar = r }
}

func HttpServer(srv *http.Server) Option {
	return func(o *options) { o.httpServer = srv }
}

func GrpcServer(srv *grpc.Server) Option {
	return func(o *options) { o.grpcServer = srv }
}
