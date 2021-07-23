package graphql

import "fmt"

type GraphQLError struct {
	err           error
	GraphqlErrors errors
}

type RequestError struct {
	err           error
	NetworkError *NetworkErrorInfo
}
type NetworkErrorInfo struct {
	StatusCode   int
	ErrorMessage string
}

func (e *GraphQLError) Error() string {
	var errMessage string

	for _, err := range e.GraphqlErrors {
		errMessage += err.Message + "\n"
	}

	return errMessage
}

func (e *RequestError) Error() string {
	var errMessage string
	if e.NetworkError != nil {
		errMessage += fmt.Sprintf("non-200 OK status code: %d", e.NetworkError.StatusCode)
	} else {
		return e.Error()
	}
	return errMessage
}
