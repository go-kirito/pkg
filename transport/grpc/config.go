/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/8/20 11:19
 */
package grpc

import (
	"time"

	"github.com/go-kirito/pkg/zconfig"
	"github.com/go-kirito/pkg/zlog"
)

type ServerConfig struct {
	Address string `json:"address"`
	Timeout int    `json:"timeout"`
}

func OptionsWithConfig() []ServerOption {
	var serverConfig ServerConfig
	if err := zconfig.UnmarshalKey("server.grpc", &serverConfig); err != nil {
		zlog.Error("读取http的配置文件信息失败")
		return nil
	}

	opts := make([]ServerOption, 0)

	if serverConfig.Address == "" {
		serverConfig.Address = ":9000"
	}

	opts = append(opts, Address(serverConfig.Address))

	if serverConfig.Timeout > 0 {
		timeout := time.Duration(serverConfig.Timeout) * time.Second
		opts = append(opts, Timeout(timeout))
	}

	return opts
}
