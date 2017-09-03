package handler

import (
	"net/http"

	"golang.org/x/net/context"
)

type ContextGenerator interface {
	Generate(rqs *http.Request) context.Context
}

type ContextGeneratorFunc func(rqs *http.Request) context.Context

func (ctxGen ContextGeneratorFunc) Generate(rqs *http.Request) context.Context {
	return ctxGen(rqs)
}
