package response

func GetOk(body string) *Response {
	return NewResponse("HTTP/1.1", "200", "Ok", body)
}

func GetNotFound(body string) *Response {
	return NewResponse("HTTP/1.1", "404", "Not Found", body)
}

func GetBadRequest(body string) *Response {
	return NewResponse("HTTP/1.1", "400", "Bad Request", body)
}

func GetInternalServerError(body string) *Response {
	return NewResponse("HTTP/1.1", "500", "Internal Server Error", body)
}

func GetCreated(body string) *Response {
	return NewResponse("HTTP/1.1", "201", "Created", body)
}

func GetNoContent() *Response {
	return NewResponse("HTTP/1.1", "204", "No Content", "")
}

func GetAccepted(body string) *Response {
	return NewResponse("HTTP/1.1", "202", "Accepted", body)
}

func GetUnauthorized(body string) *Response {
	return NewResponse("HTTP/1.1", "401", "Unauthorized", body)
}

func GetForbidden(body string) *Response {
	return NewResponse("HTTP/1.1", "403", "Forbidden", body)
}

func GetConflict(body string) *Response {
	return NewResponse("HTTP/1.1", "409", "Conflict", body)
}

func GetServiceUnavailable(body string) *Response {
	return NewResponse("HTTP/1.1", "503", "Service Unavailable", body)
}

func GetGatewayTimeout(body string) *Response {
	return NewResponse("HTTP/1.1", "504", "Gateway Timeout", body)
}
