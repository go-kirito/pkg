/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/12/13 16:22
 */
package writer

import (
	"io"
	"os"
	"time"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/protobuf/proto"
)

type AliyunWriter struct {
	opt      *AliyunOption
	instance *producer.Producer
}

type AliyunOption struct {
	AccessKey       string
	AccessKeySecret string
	EndPoint        string
	Project         string
	LogStore        string
	Topic           string
}

func NewAliyunWriter(opt *AliyunOption) io.Writer {
	a := &AliyunWriter{}
	a.opt = opt
	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = a.opt.EndPoint
	producerConfig.AccessKeyID = a.opt.AccessKey
	producerConfig.AccessKeySecret = a.opt.AccessKeySecret
	producerInstance := producer.InitProducer(producerConfig)
	producerInstance.Start()
	a.instance = producerInstance
	return a

}

func (a AliyunWriter) Write(p []byte) (n int, err error) {
	var data map[string]string
	err = jsoniter.Unmarshal(p, &data)
	if err != nil {
		return -1, err
	}

	var content []*sls.LogContent

	for k, v := range data {
		content = append(content, &sls.LogContent{
			Key:   proto.String(k),
			Value: proto.String(v),
		})
	}

	log := &sls.Log{
		Time:     proto.Uint32(uint32(time.Now().Unix())),
		Contents: content,
	}

	hostName, _ := os.Hostname()

	err = a.instance.SendLog(a.opt.Project, a.opt.LogStore, a.opt.Topic, hostName, log)
	if err != nil {
		return -1, err
	}

	return len(p), nil
}
