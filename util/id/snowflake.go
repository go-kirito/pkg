/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/4/2 15:10
 */
package id

import (
	"os"

	"github.com/go-kirito/pkg/zlog"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	n, err := snowflake.NewNode(1)
	if err != nil {
		zlog.Panic("load snowflake err", err.Error())
		os.Exit(1)
	}
	node = n
}

func New() int64 {
	return node.Generate().Int64()
}

func NewString() string {
	return node.Generate().String()
}

func NewBase64() string {
	return node.Generate().Base64()
}
