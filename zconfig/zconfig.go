/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/3/18 14:37
 */
package zconfig

import (
	"bytes"
	"io/ioutil"

	"github.com/go-kirito/pkg/zlog"

	"github.com/spf13/viper"
)

var conf *viper.Viper

func Load(path string) error {
	var (
		err error
		v   = viper.New()
	)

	v.AddConfigPath(".")
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	zlog.Infof("use config file -> %s\n", v.ConfigFileUsed())

	includes := v.GetStringSlice("includes")
	for _, i := range includes {

		fd, err := ioutil.ReadFile(i)
		if err != nil {
			zlog.Error("load config err:", err.Error())
			return err
		}

		v.MergeConfig(bytes.NewReader(fd))

		zlog.Infof("load config file -> %s\n", i)

	}

	conf = v

	return err
}

func UnmarshalKey(key string, val interface{}) error {
	return conf.UnmarshalKey(key, val)
}

func GetString(key string) string {
	return conf.GetString(key)
}