package context

import "github.com/gin-gonic/gin"

type GinContextAdapter struct {
	Ctx *gin.Context
}

func (g *GinContextAdapter) ShouldBindUri(obj interface{}) error {
	return g.Ctx.ShouldBindUri(obj)
}

func (g *GinContextAdapter) ShouldBindJSON(obj interface{}) error {
	return g.Ctx.ShouldBindJSON(obj)
}

func (g *GinContextAdapter) JSON(code int, obj interface{}) {
	g.Ctx.JSON(code, obj)
}

func (g *GinContextAdapter) Param(key string) string {
	return g.Ctx.Param(key)
}
