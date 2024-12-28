package endpoint

import (
	"net/http"
)

type Endpoint interface {
	GetEndpoint() string
	GetPath() string
	GetID() string
	GetMethods() []string
	GetDescription() string
	GetService() string
	GetRoles() []string
	GetPathValues() map[string]string
	GetHeaders() map[string]string
	GetResponseTypes() map[int][]interface{} // response_code: possible response
	GetRequestTypes() map[string]interface{} //method:body
	GetRequest(r *http.Request) (interface{}, error)
	IsValidRequest(r *http.Request) (bool, error)

	GetVars(r *http.Request) map[string]string

	AddToRoute()
	HasAccess(role string, method string) bool

	SetService(service string) Endpoint
	AddResponse(i interface{}, responseCode int) Endpoint
	AddRequest(i interface{}, method string) Endpoint
	SetPath(path ...string) Endpoint
	AddVar(varName string, parser Parser, access HasAccess) Endpoint
	AddRole(role string, method string) Endpoint
	SetDescription(description string) Endpoint
}
