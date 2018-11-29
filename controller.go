package honey

import "html/template"

// Controller defines some basic http request handler operations, such as
// http context, tempate and view, session ans xsrf
type Controller struct {
	Ct        *Context
	Tpl       *template.Template
	Data      map[interface{}]interface{}
	ChildName string
	TplNames  string
	Layout    []string
	TplExt    string
}

// ControllerInterface is an interface to uniform all contorller handler.
type ControllerInterface interface {
	Init(ct *Context, cn string) // Init context and child name.
	Prepare()                    // Prepare handles some common operations.
	Get()                        // method = GET Handler
	Post()                       // method = Post Handler
	Head()                       // method = Head Handler
	Delete()                     // method = Delete Handler
	Put()                        // method = Put Handler
	Patch()                      // method = Patch Handler
	Options()                    // method = Options Handler
	Render() error               // Render renders the page after method called.
}
