/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/12/2 17:12
 */
package response

import (
	"context"
	"io"
	"net/http"

	"github.com/go-kirito/pkg/encoding"
	"github.com/go-kirito/pkg/errors"
	"github.com/go-kirito/pkg/internal/httputil"
)

func Decoder(ctx context.Context, res *http.Response, out interface{}) error {
	codec := CodecForResponse(res)

	var resp response

	body, err := io.ReadAll(res.Body)
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

func DecodeError(ctx context.Context, res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err == nil {
		e := new(errors.Error)
		var resp response
		if err = CodecForResponse(res).Unmarshal(data, &resp); err == nil {
			e.Code = int32(res.StatusCode)
			e.Reason = resp.Code
			e.Message = resp.Message
			return e
		}
	}
	return errors.Errorf(res.StatusCode, errors.UnknownReason, err.Error())
}
