package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func init() {
	routes = append(routes, Route{"searchHandler", "GET", "/file/{id}", searchHandler})
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pattern := vars["id"]
	files := make([]string, 0)
	_ = files
	logger.Info(settings.Dir)
	err := filepath.Walk(settings.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.Error(err)
		}
		if info.IsDir() {
			//files, err = filepath.Glob(path + "*(FM ID " + pattern + ")*.pdf")
			f, err := filepath.Glob(path + "*" + pattern + "*.*")

			if err != nil {
				logger.Error(err)
			}
			if len(f) > 0 {
				//logger.Info(f)
				files = append(files, f...)
			}

		}
		return nil
	})
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
