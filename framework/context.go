package framework

import (
	"encoding/json"
	"fmt"
	"net/http"
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
