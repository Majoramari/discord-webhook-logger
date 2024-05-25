package server

import (
	"net/http"

	"github.com/majoramari/visitor-logger/utils"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/log-visitor", utils.LogVisitorInfo)
	return router
}
