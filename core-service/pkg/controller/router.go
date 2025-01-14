package controller

import (
	util "gym-core-service/pkg"
	http_error "gym-core-service/pkg/error/http_error"
	"log"
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

		log.Println(err)

		var httpError http_error.HTTPError
		if error, ok := err.(http_error.ToHttpError); ok {
			httpError = error.ToHTTPError()
		} else {
			httpError = http_error.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		util.WriteJSON(w, httpError.Code, httpError)
	}
}
