package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/kardianos/service"
)

const appVersionStr = "v1.0"
const filenameSettings = "settings.json"

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		defaultHandler,
	},
}

var router *mux.Router
var settings appSettings

type appSettings struct {
	Debug bool   `json:"debug"`
	Port  string `json:"port"`
	Dir   string `json:"dir"`
}

func main() {
	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()

	svcConfig := &service.Config{
		Name:        "o2-filesearch",
		DisplayName: "o2-filesearch",
		Description: "o2-filesearch is a simple filesearcher",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	dir, _ := filepath.Split(ex)
	dat, err := ioutil.ReadFile(dir + filenameSettings)
	if err != nil {
		data, _ := json.Marshal(settings)
		ioutil.WriteFile(dir+filenameSettings, data, 0664)
		log.Fatal("settings.json missing, " + err.Error())
	}

	if err := json.Unmarshal(dat, &settings); err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body>We are up and running ;)</body></html>")
}
