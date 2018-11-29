package honey

const (
	// Version tells honey's version
	Version string = "1.0.0"
)

// App defines web application.
type App struct {
	Handlers *ControllerRegister
	config   *Config
	// StaticDirs stores all static files directories.
	StaticDirs map[string]string
}

// SetStaticPath serves static files.
func (app *App) SetStaticPath(url string, path string) *App {
	app.StaticDirs[url] = path
	return app
}
