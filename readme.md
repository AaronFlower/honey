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

我们用正则来匹配路由，在路由时还需要将路由参数给一一解析出来，然后根据 HTTP 请求的方法，通过反映到调用对应 Controller 的方法。所以我们可以把 Route 定义如下结构：

```go
// Route 基本的路由结构体，包含一个路由 path 和对应的 controller.
type Route struct {
	path string
	controller reflect.Type
}
```

然后有一个统一的路由管理器，来完成应用程序的路由注册，与路由分发。

```go
// RouteBus 统一的路由管理器。
type RouteBus struct {
    routes []route
    App *Application
}
// 路由注册
func (bus *RouteBus)Add(path string, controller ControllerInterface) {
	// TODO 路由注册
}
// 路由分发
func (bus *RouteBus)ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // TODO 路由参数解析和分发
}
```

### `controller.go`

The controller.go file defines the controller interface.

传统的 MVC Controller 都基于 `xxAction` 即 `Action`后缀的形式来进行映射。我们在这里直接根据 REST 风格的来设计我们的 Controller.

我们 Controller 包含请求的上下文 `Context` , 对应的渲染模版等。而 Controller 对应的接口设计支持基本的 REST 方法和常用的方法即可。

```go
type Controller struct {
    Ct *Context
    Tpl *template.Template
}

type ControllerInterface interface {
	Init(ct *Context, cn string)    //初始化上下文和子类名称
	Prepare()                       //开始执行之前的一些处理
	Get()                           //method=GET的处理
	Post()                          //method=POST的处理
	Delete()                        //method=DELETE的处理
	Put()                           //method=PUT的处理
	Head()                          //method=HEAD的处理
	Patch()                         //method=PATCH的处理
	Options()                       //method=OPTIONS的处理
	Finish()                        //执行完成之后的处理		
	Render() error                  //执行完method对应的方法之后渲染页面
}
```

在新建自己的 Controller 时根据需要实现对应的方法就可以啦。

### `context.go`

The Context includes http request info.

### `config.go`

The config.go file defines the Config struct.



