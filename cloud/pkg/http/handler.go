package http

import (
	util "cloud-gym/pkg"
	"net/http"
)

type NetHttpRouteHandler func(w http.ResponseWriter, r *http.Request)
type AppRouteHandler func(w http.ResponseWriter, r *http.Request) error

func MakeRouteHandler(handler AppRouteHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err == nil {
			return
		}

		var httpError HTTPError
		if error, ok := err.(ToHttpError); ok {
			httpError = error.ToHTTPError()
		} else {
			httpError = NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		util.WriteJSON(w, httpError.Code, httpError)
	}
}
