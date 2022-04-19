package context

type IContextAdapter interface {
	ShouldBindUri(obj interface{}) error
	ShouldBindJSON(obj interface{}) error
	JSON(code int, obj interface{})
	Param(key string) string
}
