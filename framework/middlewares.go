package framework

import (
	"context"
	"fmt"
	"time"
)

func TimeOutMiddleWare(ctx *MyContext, next func(ctx *MyContext)) {

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

		time.Sleep(time.Second * 6)
		next(ctx)
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
