/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/12/2 17:12
 */
package response

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/go-kirito/pkg/encoding"
	"github.com/go-kirito/pkg/internal/httputil"
)

func Decoder(ctx context.Context, res *http.Response, out interface{}) error {
	codec := CodecForResponse(res)

	var resp response

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	err = codec.Unmarshal(body, &resp)

	if err != nil {
		return err
	}

	data, err := codec.Marshal(resp.Data)

	if err != nil {
		return err
	}

	err = codec.Unmarshal(data, &out)

	if err != nil {
		return err
	}

	return nil
}

func CodecForResponse(r *http.Response) encoding.Codec {
	codec := encoding.GetCodec(httputil.ContentSubtype(r.Header.Get("Content-Type")))
	if codec != nil {
		return codec
	}
	return encoding.GetCodec("json")
}
