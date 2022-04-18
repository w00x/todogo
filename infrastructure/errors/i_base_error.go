package errors

type IBaseError interface {
	Error() map[string]interface{}
	HttpStatusCode() int
}