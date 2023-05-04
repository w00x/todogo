package infrastructure

import (
	"github.com/gin-gonic/gin"
	"todogo/infrastructure/factory"
)

func GinRoutes(factoryAdapter string) *gin.Engine {
	route := gin.Default()
	contextAdapter := "gin"

	version1 := route.Group("/v1")
	{
		todo := version1.Group("todo")
		{
			todoController := InitializeTodoController(factoryAdapter)
			todo.GET("/", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				todoController.Index(ctx)
			})
			todo.GET("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				todoController.Get(ctx)
			})
			todo.POST("", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				todoController.Create(ctx)
			})
			todo.PATCH("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				todoController.Update(ctx)
			})
			todo.DELETE("/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				todoController.Delete(ctx)
			})

			todo.GET("/toggle/:id", func(context *gin.Context) {
				ctx, err := factory.ContextFactory(contextAdapter, context)
				if err != nil {
					panic(err)
				}
				todoController.Toggle(ctx)
			})
		}
	}

	return route
}
