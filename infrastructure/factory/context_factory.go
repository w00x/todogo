package factory

import (
	"github.com/gin-gonic/gin"
	context2 "todogo/infrastructure/api/context"
	"todogo/infrastructure/errors"
)

func ContextFactory(adapter string, ctx interface{}) (context2.IContextAdapter, errors.IBaseError) {
	if adapter == "gin" {
		return &context2.GinContextAdapter{Ctx: ctx.(*gin.Context)}, nil
	}

	return nil, errors.NewNotFoundError("Todo Factory not found")
}