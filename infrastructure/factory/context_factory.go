package factory

import (
	"github.com/gin-gonic/gin"
	"todogohexa/infrastructure/context"
	"todogohexa/infrastructure/errors"
)

func ContextFactory(adapter string, ctx interface{}) (context.IContextAdapter, errors.IBaseError) {
	if adapter == "gin" {
		return &context.GinContextAdapter{Ctx: ctx.(*gin.Context)}, nil
	}

	return nil, errors.NewNotFoundError("Todo Factory not found")
}