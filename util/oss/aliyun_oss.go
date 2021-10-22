/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/10/19 19:15
 */
package util

import (
	"errors"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/go-kirito/pkg/zconfig"
)

type OssConfig struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
	RegionId        string
	Bucket          string
}

func GetOssConfig() (*OssConfig, error) {

	config := zconfig.GetStringMap("oss")
	if config == nil {
		return nil, errors.New("oss config not found")
	}

	var regionId string
	var accessKeyId string
	var accessKeySecret string
	var arn string
	var bucket string
	var ok bool

	if regionId, ok = config["regionid"].(string); !ok {
		return nil, errors.New("oss config regionId not found")
	}

	if accessKeyId, ok = config["accesskeyid"].(string); !ok {
		return nil, errors.New("oss config accessKeyId not found")
	}

	if accessKeySecret, ok = config["accesskeysecret"].(string); !ok {
		return nil, errors.New("oss config accessKeySecret not found")
	}

	if arn, ok = config["arn"].(string); !ok {
		return nil, errors.New("oss config arn not found")
	}

	if bucket, ok = config["bucket"].(string); !ok {
		return nil, errors.New("oss config bucket not found")
	}
	//构建一个阿里云客户端, 用于发起请求。
	//构建阿里云客户端时，需要设置AccessKey ID和AccessKey Secret。
	client, err := sts.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)

	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	//设置参数。关于参数含义和设置方法，请参见API参考。
	request.RoleArn = arn
	request.RoleSessionName = "stsSession"

	//发起请求，并得到响应。
	response, err := client.AssumeRole(request)
	if err != nil {
		return nil, err
	}

	ossConfig := &OssConfig{
		AccessKeyId:     response.Credentials.AccessKeyId,
		AccessKeySecret: response.Credentials.AccessKeySecret,
		SecurityToken:   response.Credentials.SecurityToken,
		RegionId:        regionId,
		Bucket:          bucket,
	}

	return ossConfig, nil
}
