package execute

import ()

// Service
type Service interface {
	// Name
	Name() string
	// Init
	Init() error
	// Start
	Start() error
	// Stop
	Stop() error
}
