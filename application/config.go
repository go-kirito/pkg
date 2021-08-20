/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/8/20 10:22
 */
package application

import (
	"github.com/go-kirito/pkg/transport"
	"github.com/go-kirito/pkg/transport/grpc"
	"github.com/go-kirito/pkg/transport/http"
	"github.com/go-kirito/pkg/zconfig"
)

type ServerConfig struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func loadServerConfig() *ServerConfig {
	serverConfig := new(ServerConfig)
	if err := zconfig.UnmarshalKey("application", serverConfig); err != nil {
		return nil
	}
	return serverConfig
}

// New create on application from config
func NewWithConfig() *App {

	serverConfig := loadServerConfig()

	opts := make([]Option, 0)

	if serverConfig.Name != "" {
		opts = append(opts, Name(serverConfig.Name))
	}

	if serverConfig.Version != "" {
		opts = append(opts, Version(serverConfig.Version))
	}

	tr := make([]transport.Server, 0)

	grpcOptions := grpc.OptionsWithConfig()

	if grpcOptions != nil {
		grpcServer := grpc.NewServer(grpcOptions...)
		opts = append(opts, GrpcServer(grpcServer))
		tr = append(tr, grpcServer)
	}

	httpOptions := http.OptionsWithConfig()

	if httpOptions != nil {
		//初始化http
		httpServer := http.NewServer(httpOptions...)
		opts = append(opts, HttpServer(httpServer))
		tr = append(tr, httpServer)
	}

	if len(tr) > 0 {
		opts = append(opts, Server(tr...))
	}

	return New(opts...)
}
