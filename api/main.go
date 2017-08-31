package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"

	"feedback/api/routers/root"
)

// Version TODO
// Build TODO
var (
	Version string
	Build   string
	Hash    string
)

func initRouter(router *httprouter.Router) {
	root.NewRouter(router, Version, Build, Hash)
}

func main() {
	var (
		port = flag.String("port", os.Getenv("FEEDBACK_API_PORT"), "Port to listen to")
	)
	flag.Parse()
	log.Info("Initializing feedback api...")

	// Set the addr with the port
	var bindingPort int
	bindingPort, err := strconv.Atoi(*port)
	if err != nil {
		bindingPort = 8888
	}
	addr := fmt.Sprintf(":%d", bindingPort)

	log.Infof("Version: %s", Version)
	log.Infof("Git commit hash: %s", Build)

	// Setup the route from the api to the router
	router := httprouter.New()
	initRouter(router)

	// Start the api
	log.Infof("Feedback is running at %s.", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
