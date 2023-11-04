package framework

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func TimeOutMiddleWare(ctx *MyContext) {
	successCh := make(chan struct{})
	panicCh := make(chan struct{})
	durationContext, cancel := context.WithTimeout(ctx.Request().Context(), time.Second*5)

	defer cancel()

	go func() {

		defer func() {
			if err := recover(); err != nil {
				panicCh <- struct{}{}
			}
		}()

		ctx.Next()
		// time.Sleep(time.Second * 6)
		successCh <- struct{}{}
	}()

	select {
	case <-durationContext.Done():
		ctx.WriteString("timeout")
		ctx.SetHasTimeout(true)
	case <-panicCh:
		ctx.WriteString("panic")
	case <-successCh:
		fmt.Println("success")
	}
}

func AuthUserMiddleware(ctx *MyContext) {
	ctx.Set("AuthUser", "test")
}

func TimeCostMiddleware(ctx *MyContext) {
	now := time.Now()
	ctx.Next()
	fmt.Println("time cost:", time.Since(now).Milliseconds())
}

func StaticFileMiddleware(ctx *MyContext) {
	fileServer := http.FileServer(http.Dir("./static"))

	pathname := ctx.Request().URL.Path
	pathname = strings.TrimSuffix(pathname, "/")

	fPath := path.Join("./static", pathname)
	fInfo, err := os.Stat(fPath)

	fExist := err == nil && !fInfo.IsDir()
	if fExist {
		fileServer.ServeHTTP(ctx.ResponseWritor(), ctx.Request())
		ctx.Abort()
		return
	}

}
