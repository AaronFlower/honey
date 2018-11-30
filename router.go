package honey

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

// controllerInfo stores a controller informationt.
type controllerInfo struct {
	regex          *regexp.Regexp
	params         map[int]string
	controllerType reflect.Type
}

// ControllerRegister contains all registered router rules.
type ControllerRegister struct {
	routes []*controllerInfo
	App    *App
}

// Add controller pattern rules and handler to ControllerRegister
// usage:
// Add("/", &MainController{})
// Add("/user/:param", &UserController{})
// @TODO complete register with methods.
func (cr *ControllerRegister) Add(pattern string, c ControllerInterface) {
	parts := strings.Split(pattern, "/")

	// parsing the url to get the parameter and set regexp group capture placeholders.
	j := 0
	params := make(map[int]string)
	for i, part := range parts {
		// check if is a parameter.
		if strings.HasPrefix(part, ":") {
			expr := "([^/]+)"
			params[j] = part
			parts[i] = expr // replace with group capture placehodler.
			j++
		}
	}

	// recreate the url pattern, with parameters replaced by regular expressions.
	// then compile the regexp.
	pattern = strings.Join(parts, "/")
	regex, err := regexp.Compile(pattern)
	if err != nil {
		// @TODO add error handling here to avoid panic.
		panic(err)
	}

	// create the Route
	t := reflect.Indirect(reflect.ValueOf(c)).Type()
	route := &controllerInfo{}
	route.regex, route.params, route.controllerType = regex, params, t

	cr.routes = append(cr.routes, route)

}

// Use the routes of ControllerRegister to auto route.
func (cr *ControllerRegister) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var started bool

	fmt.Printf("cr = %+v\n", cr)
	fmt.Printf("cr.App = %+v\n", cr.App)
	// first check static path
	for prefix, staticDir := range cr.App.StaticDirs {
		if strings.HasPrefix(r.URL.Path, prefix) {
			file := staticDir + r.URL.Path[len(prefix):]
			http.ServeFile(w, r, file)
			started = true
			return
		}
	}
	fmt.Println("Server HTTP 2")

	// find a matching Route
	path := r.URL.Path
	for _, route := range cr.routes {
		// check if route pattern matches url
		if !route.regex.MatchString(path) {
			continue
		}

		// capture group parameters
		// FindStringSubMatch returns an array, the first element is all matched string.
		// 2nd is the 1th group, 3rd is the 2nd group...
		matches := route.regex.FindStringSubmatch(path)

		// double check that the route matches the URL pattern.
		if len(matches[0]) != len(path) {
			continue
		}

		params := make(map[string]string)
		if len(route.params) > 0 {
			// add url parameters to the query param map
			values := r.URL.Query()
			for i, match := range matches[1:] {
				values.Add(route.params[i], match)
				params[route.params[i]] = match
			}

			// reassemble query params and add to RawQuery
			r.URL.RawQuery = url.Values(values).Encode() + "&" + r.URL.RawQuery
		}

		// Invoke the request handler
		// our controller should implements Init, Prepare, Get, Post, Head, Delete, Put, Patch,
		// Options, Finish methods.
		vc := reflect.New(route.controllerType)
		init := vc.MethodByName("Init")
		in := make([]reflect.Value, 2)
		ct := &Context{ResponseWriter: w, Request: r}
		in[0] = reflect.ValueOf(ct)
		in[1] = reflect.ValueOf(route.controllerType.Name())
		init.Call(in)

		in = make([]reflect.Value, 0)
		method := vc.MethodByName("Prepare")
		method.Call(in)

		methodsMap := map[string]string{
			"GET":     "Get",
			"POST":    "Post",
			"HEAD":    "Head",
			"DELETE":  "Delete",
			"PUT":     "Put",
			"PATCH":   "Patch",
			"OPTIONS": "Options",
		}

		method = vc.MethodByName(methodsMap[r.Method])
		method.Call(in)

		method = vc.MethodByName("Finish")
		method.Call(in)
		started = true
		break
	}

	if started == false {
		http.NotFound(w, r)
	}

}
