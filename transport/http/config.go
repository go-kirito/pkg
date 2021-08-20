/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/8/20 10:44
 */
package http

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
	if err := zconfig.UnmarshalKey("server.http", &serverConfig); err != nil {
		zlog.Error("读取http的配置文件信息失败")
		return nil
	}

	opts := make([]ServerOption, 0)

	opts = append(opts, Address(serverConfig.Address))

	timeout := time.Duration(serverConfig.Timeout) * time.Second
	opts = append(opts, Timeout(timeout))

	return opts
}
