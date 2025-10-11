package notebook

import (
	"net/http"
)

type Handler interface {
	Get() http.HandlerFunc
	Post() http.HandlerFunc
	Put() http.HandlerFunc
}
