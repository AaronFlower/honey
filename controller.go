package honey

import (
	"html/template"
	"net/http"
)

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
	Finish()
}

// Init context and child name.
func (c *Controller) Init(ct *Context, cn string) {
	c.Data = make(map[interface{}]interface{})
	c.Layout = make([]string, 0)
	c.TplNames = ""
	c.ChildName = cn
	c.Ct = ct
	c.TplExt = "tpl"
}

// Prepare handles some common operations.
func (c *Controller) Prepare() {
}

// Get Handler
func (c *Controller) Get() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

// Post Handler
func (c *Controller) Post() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

// Head Handler
func (c *Controller) Head() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

// Delete Handler
func (c *Controller) Delete() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

// Put Handler
func (c *Controller) Put() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

// Patch Handler
func (c *Controller) Patch() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

// Options Handler
func (c *Controller) Options() {
	http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

// Render Handler
func (c *Controller) Render() error {
	return nil
}

// Finish handles some common operations.
func (c *Controller) Finish() {
}
