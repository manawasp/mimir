package feedback

import "github.com/julienschmidt/httprouter"

// NewRouter intialize a router with User ressources
func NewRouter(router *httprouter.Router) {
	router.GET("/api/feedback", scoresRetrieve)
}
