/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/8/20 17:59
 */
package response

import (
	"net/http"

	"github.com/go-kirito/pkg/encoding"
	"github.com/go-kirito/pkg/errors"
	"github.com/go-kirito/pkg/internal/httputil"
)

type response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// DefaultResponseEncoder encodes the object to the HTTP response.
func Encoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	codec, _ := CodecForRequest(r, "Accept")

	resp := response{
		Code:    "0",
		Message: "",
		Data:    v,
	}
	data, err := codec.Marshal(resp)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", httputil.ContentType(codec.Name()))
	w.Write(data)
	return nil
}

// CodecForRequest get encoding.Codec via http.Request
func CodecForRequest(r *http.Request, name string) (encoding.Codec, bool) {
	for _, accept := range r.Header[name] {
		codec := encoding.GetCodec(httputil.ContentSubtype(accept))
		if codec != nil {
			return codec, true
		}
	}
	return encoding.GetCodec("json"), false
}

func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := errors.FromError(err)
	codec, _ := CodecForRequest(r, "Accept")

	resp := response{
		Code:    se.Reason,
		Message: se.Message,
		Data:    nil,
	}

	body, err := codec.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", httputil.ContentType(codec.Name()))
	w.WriteHeader(int(se.Code))
	w.Write(body)
}
