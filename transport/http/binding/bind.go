package binding

import (
	"net/http"
	"net/url"

	"github.com/go-kirito/pkg/encoding"
	"github.com/go-kirito/pkg/encoding/form"
)

// BindQuery bind vars parameters to target.
func BindQuery(vars url.Values, target interface{}) error {
	return encoding.GetCodec(form.Name).Unmarshal([]byte(vars.Encode()), target)
}

// BindForm bind form parameters to target.
func BindForm(req *http.Request, target interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	return encoding.GetCodec(form.Name).Unmarshal([]byte(req.Form.Encode()), target)
}
