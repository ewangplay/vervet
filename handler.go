package vervet

import (
	"net/http"
)

type ProcessFunc func(
	version int,
	resource string,
	method string,
	params map[string]string,
	body []byte,
	result map[string]interface{}) error

type Handler interface {
	Process(
		r *http.Request,
		version int,
		resource string,
		f ProcessFunc) (string, error)

	ProcessFunc(
		version int,
		resource string,
		method string,
		params map[string]string,
		body []byte,
		result map[string]interface{}) error
}
