package honey

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"time"
)

const (
	// Version tells honey's version
	Version string = "1.0.0"
)

var (
	// MyApp is the application instance
	MyApp *App
	// AppPath is the application path.
	AppPath string
)

// App defines web application.
type App struct {
	Handlers *ControllerRegister
	config   *Config
	// StaticDirs stores all static files directories.
	StaticDirs map[string]string
}

func init() {
	MyApp = newApp(nil)
	AppPath, _ = os.Getwd()
}

func newApp(config *Config) *App {
	cr := &ControllerRegister{}
	app := &App{
		Handlers:   cr,
		config:     config,
		StaticDirs: make(map[string]string),
	}
	cr.App = app
	return app
}

// Run starts the server
func (app *App) Run() {
	if app.config.HTTPAddr == "" {
		app.config.HTTPAddr = "127.0.0.1"
	}

	addr := fmt.Sprintf("%s:%d", app.config.HTTPAddr, app.config.HTTPPort)

	var err error
	for {
		if app.config.UseFcgi {
			fmt.Println("The Fcgi")
			listener, err := net.Listen("tcp", addr)
			if err != nil {
				log.Print("Listen:", err)
			}
			err = fcgi.Serve(listener, app.Handlers)
		} else {
			err = httpListenAndServe(addr, app.Handlers, app.config)
			if err != nil {
				os.Exit(1)
			}
		}

		if err != nil {
			log.Print("ListenAndServe:", err)
		}
		time.Sleep(time.Second * 2)
	}
}

// SetStaticPath serves static files.
func (app *App) SetStaticPath(url string, path string) *App {
	app.StaticDirs[url] = path
	return app
}

func httpListenAndServe(addr string, handler http.Handler, config *Config) error {
	readTimeout := 5 * time.Second
	if config.ReadTimeout != 0 {
		readTimeout = config.ReadTimeout
	}

	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  readTimeout,
		WriteTimeout: config.WriteTimeout,
	}
	return server.ListenAndServe()
}

// Run starts the server
func Run(config *Config) {
	MyApp.config = config
	MyApp.Run()
}
