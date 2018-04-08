//
// service - define the routes and start the service
//
// @author darryl.west@ebay.com
// @created 2018-04-08 09:26:59
//

package edit

import (
	"fmt"
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

// Start start the admin listener and event loop
func (svc Service) Start() error {
	log.Info("start the hub service...")
	// cfg := svc.cfg

	// create and open the registry database

	// start the listener
	if err := svc.startServer(); err != nil {
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
