package execute

import (
	"fmt"
)

var srv Service

func StartService(srv Service) (err error) {
	// init
	err = srv.Init()
	if err != nil {
		return err
	}
	// start
	err = srv.Start()
	if err != nil {
		return err
	}

	return nil
}

func StopService() (err error) {
	if srv != nil {
		return srv.Stop()
	}
	return fmt.Errorf("Service is not started")
}

