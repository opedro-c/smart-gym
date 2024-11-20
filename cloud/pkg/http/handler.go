package http

import "net/http"

type NetHttpRouteHandler func(w http.ResponseWriter, r *http.Request)
type AppRouteHandler func(w http.ResponseWriter, r *http.Request) error

func MakeRouteHandler(handler AppRouteHandler) NetHttpRouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err == nil {
			return
		}

		if err, ok := err.(ToHttpError); ok {
			httpError := err.ToHTTPError()
			http.Error(w, httpError.Message, httpError.Code)
			return
		}

		// Default error return
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
