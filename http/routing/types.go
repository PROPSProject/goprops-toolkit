package routing

import (
	"bytes"
	"fmt"
	"net/http"
)

//Route route struct
type Route struct {
	Name         string
	Method       string
	ResourcePath string
	Version      string
	NameSpace    string
	HandlerFunc  http.HandlerFunc
}

func (r *Route) String() string {
	return fmt.Sprintf("Name: %v, Method: %v, Version: %v, URI: %v",
		r.Name, r.Method, r.Version, r.GetURI(),
	)
}

// GetURI ...
func (r *Route) GetURI() string {
	return fmt.Sprintf("%v/%v/%v", r.NameSpace, r.Version, r.ResourcePath)
}

// Routes ...
type Routes []Route

func (r *Routes) String() string {
	var buffer bytes.Buffer
	for _, route := range *r {
		buffer.WriteString(fmt.Sprintf("\t%v\n", route))
	}

	return buffer.String()
}
