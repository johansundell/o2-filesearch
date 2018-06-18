package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/mattn/go-zglob"
)

func init() {
	routes = append(routes, Route{"searchHandler", "GET", "/file/{id}", searchHandler})
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pattern := vars["id"]

	files, err := zglob.Glob(settings.Dir + "**/*" + pattern + "*.*")
	if err != nil {
		logger.Error(err)
	}
	//logger.Info(files)
	for _, pa := range files {
		_, f := filepath.Split(pa)
		fmt.Fprintf(w, "<a target=\"_blank\" href=\"file://"+pa+"\">"+f+"</a></br>")
		// fmt.Fprintf(w, "<a target=\"_blank\" href=\"file://"+pa[2:]+"\">"+f+"</a></br>")

	}
}
