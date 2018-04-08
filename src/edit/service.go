//
// service - define the routes and starts the web service
//
// @author darryl.west@ebay.com
// @created 2018-04-08 09:26:59
//

package edit

import (
	"fmt"
    "net/http"
)

// Service - the service struct
type Service struct {
	cfg *Config
}

// NewService create a new service by passing in config
func NewService(cfg *Config) (*Service, error) {
	svc := Service{
		cfg: cfg,
	}

	return &svc, nil
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    log.Info("show the about page")

    // serve a template file...
    // w.Header().Set("Content-Type", "plain/html")
    fmt.Fprintf(w, "<!doctype html><html><pre>%s</pre><h5>Version %s</h5><html>", logo, Version())
}

// Start start the admin listener and event loop
func (svc Service) Start() error {
	log.Info("start the hub service...")
	cfg := svc.cfg

    fs := http.FileServer(http.Dir(cfg.StaticFolder))
    http.Handle("/", fs)

    http.HandleFunc("/about", aboutHandler)

	// start the listener
    host := fmt.Sprintf(":%d", cfg.Port)
    log.Info("start listening on port: %s", host)

    err := http.ListenAndServe(host, nil)
    if err != nil {
        log.Error("http error: %s", err)
        return err
    }

	return nil
}

func (svc Service) startServer() error {
	cfg := svc.cfg
	port := cfg.Port

	host := fmt.Sprintf(":%d", port)
	log.Info("start listening on port %s", host)

	return nil
}
