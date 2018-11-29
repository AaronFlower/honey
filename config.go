package honey

import "time"

// Config configs the app settings.
type Config struct {
	HTTPAddr     string
	HTTPPort     int
	TemplatePath string
	RecoverPanic bool
	RunMode      int8 // 0 = prod, 1 = dev
	UseFcgi      bool
	ReadTimeout  time.Duration // maximum duration before timing out read of the request.
	WriteTimeout time.Duration // maximum duration before timing out read of the request.
}

const (
	// RunModeProd means production enviroment.
	RunModeProd int8 = 0
	// RunModeDev means development enviroment.
	RunModeDev int8 = 1
)
