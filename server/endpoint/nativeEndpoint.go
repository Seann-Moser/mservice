package endpoint

import "net/http"

var _ Endpoint = &NativeEndpoint{}

type NativeEndpoint struct {
	id          string
	description string
	service     string

	path string
	vars map[string]Value

	roles map[string]struct{}

	access  map[string]string // method : role
	methods []string

	response map[string][]interface{}
	request  map[string]interface{}
}

func (n *NativeEndpoint) GetRoles() []string {
	//TODO implement me
	panic("implement me")
}

func (n *NativeEndpoint) GetPathValues() map[string]string {
	//TODO implement me
	panic("implement me")
}

func (n *NativeEndpoint) GetHeaders() map[string]string {
	//TODO implement me
	panic("implement me")
}

func (n *NativeEndpoint) GetResponseTypes() map[int][]interface{} {
	//TODO implement me
	panic("implement me")
}

func (n *NativeEndpoint) GetRequestTypes() map[string]interface{} {
	//TODO implement me
	panic("implement me")
}

func (n *NativeEndpoint) GetRequest(r *http.Request) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NativeEndpoint) IsValidRequest(r *http.Request) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NativeEndpoint) GetVars(r *http.Request) map[string]string {
	//TODO implement me
	panic("implement me")
}

func (n *NativeEndpoint) AddToRoute() {
	//TODO implement me
	panic("implement me")
}

func (n *NativeEndpoint) HasAccess(role string, method string) bool {
	//TODO implement me
	panic("implement me")
}

func (n *NativeEndpoint) AddResponse(i interface{}, responseCode int) Endpoint {
	//TODO implement me
	panic("implement me")
}

func (n *NativeEndpoint) AddRequest(i interface{}, method string) Endpoint {
	//TODO implement me
	panic("implement me")
}

func (n *NativeEndpoint) SetPath(path ...string) Endpoint {
	//TODO implement me
	panic("implement me")
}

func (n *NativeEndpoint) AddVar(varName string, parser Parser, access HasAccess) Endpoint {
	//TODO implement me
	panic("implement me")
}

func (n *NativeEndpoint) AddRole(role string, method string) Endpoint {
	//TODO implement me
	panic("implement me")
}

func (n *NativeEndpoint) SetDescription(description string) Endpoint {
	//TODO implement me
	panic("implement me")
}

func (n *NativeEndpoint) GetService() string {
	return n.service
}

func (n *NativeEndpoint) SetService(service string) Endpoint {
	n.service = service
	return n
}

func (n *NativeEndpoint) GetEndpoint() string {
	return n.path
}

func (n *NativeEndpoint) GetPath() string {
	return n.path
}

func (n *NativeEndpoint) GetID() string {
	if n.id == "" {
		n.id = "" // methods, path, roles
	}
	return n.id
}

func (n *NativeEndpoint) GetMethods() []string {
	return n.methods
}

func (n *NativeEndpoint) GetDescription() string {
	return n.description
}
