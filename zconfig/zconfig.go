/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/3/18 14:37
 */
package zconfig

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-kirito/pkg/util/crypt"
	"io/ioutil"
	"log"
	"os"
	"strings"

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

	log.Printf("[config] Load Config File:%s\n", v.ConfigFileUsed())

	includes := v.GetStringSlice("includes")
	for _, i := range includes {

		fd, err := ioutil.ReadFile(i)
		if err != nil {
			log.Fatal("[config] Load Config err:", err.Error())
			return err
		}

		v.MergeConfig(bytes.NewReader(fd))

		log.Printf("[config] Load Config File:%s\n", i)

	}

	conf = v

	return err
}

func LoadEncrypt(path string, secretKey string, iv string) error {
	keyLength := len(secretKey)
	if keyLength != 16 && keyLength != 24 && keyLength != 32 {
		return errors.New("key length must be 16, 24, 32")
	}
	var (
		err error
		v   = viper.New()
	)

	v.AddConfigPath(".")

	v.SetConfigFile(path)

	fd, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	plainText, err := crypt.AesDecrypt(string(fd), secretKey, iv)

	if err != nil {
		return err
	}

	if err := v.ReadConfig(bytes.NewReader(plainText)); err != nil {
		return err
	}

	log.Printf("[config] Load Config File:%s\n", v.ConfigFileUsed())

	includes := v.GetStringSlice("includes")
	for _, i := range includes {

		fd, err := os.ReadFile(i)
		if err != nil {
			log.Fatal("[config] Load Config err:", err.Error())
			return err
		}

		plainText, err := crypt.AesDecrypt(string(fd), secretKey, iv)

		if err != nil {
			return err
		}

		v.MergeConfig(bytes.NewReader(plainText))

		log.Printf("[config] Load Config File:%s\n", i)

	}

	conf = v

	return err
}

func WriteEncryptConfig(path string, secretKey string, iv string) error {
	var (
		err error
		v   = viper.New()
	)

	v.AddConfigPath(".")

	v.SetConfigFile(path)

	keyLength := len(secretKey)
	if keyLength != 16 && keyLength != 24 && keyLength != 32 {
		return errors.New("key length must be 16, 24, 32")
	}

	fd, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	encrypt, err := crypt.AseEncrypt(fd, []byte(secretKey), []byte(iv))

	if err != nil {
		return err
	}

	arr := strings.Split(path, ".")
	filename := arr[:len(arr)-1]
	ext := arr[len(arr)-1]

	newFileName := fmt.Sprintf("%s_encrypt.%s", filename, ext)

	err = os.WriteFile(newFileName, []byte(encrypt), 0644)

	if err != nil {
		return err
	}

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	log.Printf("[config] Load Config File:%s\n", v.ConfigFileUsed())

	includes := v.GetStringSlice("includes")
	for _, path := range includes {

		fd, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		encrypt, err := crypt.AseEncrypt(fd, []byte(secretKey), []byte(iv))

		if err != nil {
			return err
		}

		arr := strings.Split(path, ".")
		filename := arr[:len(arr)-1]
		ext := arr[len(arr)-1]

		newFileName := fmt.Sprintf("%s_encrypt.%s", filename, ext)

		err = os.WriteFile(newFileName, []byte(encrypt), 0644)

		if err != nil {
			return err
		}
	}

	return nil
}

func UnmarshalKey(key string, val interface{}) error {
	return conf.UnmarshalKey(key, val)
}

func GetString(key string) string {
	return conf.GetString(key)
}

func GetInt64(key string) int64 {
	return conf.GetInt64(key)
}

func GetInt(key string) int {
	return conf.GetInt(key)
}

func GetInt32(key string) int32 {
	return conf.GetInt32(key)
}

func GetStringMap(key string) map[string]interface{} {
	return conf.GetStringMap(key)
}
