package framework

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"net/textproto"
)

type MyContext struct {
	rw     http.ResponseWriter
	r      *http.Request
	params map[string]string
}

func NewMyContext(rw http.ResponseWriter, r *http.Request) *MyContext {
	return &MyContext{
		rw:     rw,
		r:      r,
		params: map[string]string{},
	}
}

func (ctx *MyContext) BindJson(data any) error {
	byteData, err := io.ReadAll(ctx.r.Body)
	if err != nil {
		return err
	}

	ctx.r.Body = io.NopCloser(bytes.NewBuffer(byteData))

	return json.Unmarshal(byteData, data)
}

func (ctx *MyContext) Json(data any) {
	responseData, err := json.Marshal(data)
	if err != nil {
		ctx.rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx.rw.Header().Set("Content-Type", "application/json")
	ctx.rw.WriteHeader(http.StatusOK)

	ctx.rw.Write(responseData)
}

func (ctx *MyContext) WriteString(data string) {
	ctx.rw.WriteHeader(http.StatusOK)
	fmt.Fprint(ctx.rw, data)
}

func (ctx *MyContext) WriteHeader(httpCode int) {
	ctx.rw.WriteHeader(httpCode)
}

func (ctx *MyContext) QueryAll(query string) map[string][]string {
	return ctx.r.URL.Query()
}

func (ctx *MyContext) QueryKey(key string, defaultValue string) string {
	values := ctx.r.URL.Query()
	if target, ok := values[key]; ok {
		if len(target) == 0 {
			return defaultValue
		}
		return target[len(target)-1]
	}
	return defaultValue
}

func (ctx *MyContext) SetParams(dicts map[string]string) {
	ctx.params = dicts
}

func (ctx *MyContext) GetParam(key string, defaultValue string) string {
	params := ctx.params

	if v, ok := params[key]; ok {
		return v
	}

	return defaultValue
}

func (ctx *MyContext) FormKey(key string, defaultValue string) string {
	if ctx.r.Form == nil {
		ctx.r.ParseMultipartForm(32 << 20)
	}
	if vs := ctx.r.Form[key]; len(vs) > 0 {
		return vs[0]
	}
	return defaultValue
}

type FormFileInfo struct {
	Data     []byte
	Filename string
	Header   textproto.MIMEHeader
	Size     int64
}

func (ctx *MyContext) FormFile(key string) (*FormFileInfo, error) {
	file, fileHeader, err := ctx.r.FormFile(key)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return &FormFileInfo{
		Data:     data,
		Filename: fileHeader.Filename,
		Header:   fileHeader.Header,
		Size:     fileHeader.Size,
	}, nil
}

func (ctx *MyContext) RenderHtml(filepath string, data any) error {
	t, err := template.ParseFiles(filepath)
	if err != nil {
		return err
	}
	return t.Execute(ctx.rw, data)
}
