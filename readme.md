# Honey
A simple web MVC framework.

The framework provides some essential functions.

1. Initialize configurations and start to listen port.
2. auto router
3. log
4. Database, Session support.

### `honey.go`

The entrance file to initialize the app and serve at the specified address.

### `router.go`

The router.go file supports route registration, routing and `ServeHTTP` calling.

### `controller.go`

The controller.go file defines the controller interface.

### `context.go`

The Context includes http request info.

### `config.go`

The config.go file defines the Config struct.



