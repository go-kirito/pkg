/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/8/24 18:09
 */
package errors

import (
	"log"
	"testing"
)

var (
	ErrUserNotExists = NotFound("100001", "用户不存在")
)

func TestGenErrCodeDoc(t *testing.T) {
	got, err := GenErrCodeDoc("demo")
	if err != nil {
		t.Error(err)
	}

	log.Print("got:", string(got))
}
