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

	zlog.Infof("[config] Load Config File:%s\n", v.ConfigFileUsed())

	includes := v.GetStringSlice("includes")
	for _, i := range includes {

		fd, err := ioutil.ReadFile(i)
		if err != nil {
			zlog.Fatal("load config err:", err.Error())
			return err
		}

		v.MergeConfig(bytes.NewReader(fd))

		zlog.Infof("[config] Load Config File:%s\n", i)

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

func GetStringMap(key string) map[string]interface{} {
	return conf.GetStringMap(key)
}
