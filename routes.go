package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func (s *Server) route() {
	s.router.GET("/", mainHandler())
	s.router.GET("/:name", nameHandler())
}

func mainHandler() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintln(ctx, string(ctx.UserAgent()))
	}
}

func nameHandler() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintln(ctx, ctx.UserValue("name"))
	}
}
