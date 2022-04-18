package errors

type BaseError struct {
	Message string
	Description string
	Code int
	httpStatusCode int
}

func (e *BaseError) Error() map[string]interface{} {
	err := make(map[string]interface{})
	err["message"] = e.Message
	err["description"] = e.Description
	err["code"] = e.Code
	return err
}

func (e *BaseError) HttpStatusCode() int {
	return e.httpStatusCode
}