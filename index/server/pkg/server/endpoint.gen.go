//
// Copyright 2023 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4. **DO NOT EDIT**
package server // ServerInterface represents all server handlers.
import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

type ServerInterface interface {
	// Root endpoint of registry server.
	// (GET /)
	ServeRootEndpoint(c *gin.Context)
	// Get devfile by stack name.
	// (GET /devfiles/{stack})
	ServeDevfile(c *gin.Context, stack string)
	// Fetches starter project by stack and project name
	// (GET /devfiles/{stack}/starter-projects/{starterProject})
	ServeDevfileStarterProject(c *gin.Context, stack string, starterProject string)
	// Get devfile by stack name.
	// (GET /devfiles/{stack}/{version})
	ServeDevfileWithVersion(c *gin.Context, stack string, version string)
	// Fetches starter project by stack name, stack version, and project name
	// (GET /devfiles/{stack}/{version}/starter-projects/{starterProject})
	ServeDevfileStarterProjectWithVersion(c *gin.Context, stack string, version string, starterProject string)
	// Get health status.
	// (GET /health)
	ServeHealthCheck(c *gin.Context)
	// Gets index schemas of the stack devfiles.
	// (GET /index)
	ServeDevfileIndexV1(c *gin.Context, params ServeDevfileIndexV1Params)
	// Gets index schemas of the devfiles of specific type.
	// (GET /index/{indexType})
	ServeDevfileIndexV1WithType(c *gin.Context, indexType string, params ServeDevfileIndexV1WithTypeParams)
	// Gets V2 index schemas of the stack devfiles.
	// (GET /v2index)
	ServeDevfileIndexV2(c *gin.Context, params ServeDevfileIndexV2Params)
	// Gets V2 index schemas of the devfiles of specific type.
	// (GET /v2index/{indexType})
	ServeDevfileIndexV2WithType(c *gin.Context, indexType string, params ServeDevfileIndexV2WithTypeParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// ServeRootEndpoint operation middleware
func (siw *ServerInterfaceWrapper) ServeRootEndpoint(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ServeRootEndpoint(c)
}

// ServeDevfile operation middleware
func (siw *ServerInterfaceWrapper) ServeDevfile(c *gin.Context) {

	var err error

	// ------------- Path parameter "stack" -------------
	var stack string

	err = runtime.BindStyledParameter("simple", false, "stack", c.Param("stack"), &stack)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter stack: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ServeDevfile(c, stack)
}

// ServeDevfileStarterProject operation middleware
func (siw *ServerInterfaceWrapper) ServeDevfileStarterProject(c *gin.Context) {

	var err error

	// ------------- Path parameter "stack" -------------
	var stack string

	err = runtime.BindStyledParameter("simple", false, "stack", c.Param("stack"), &stack)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter stack: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "starterProject" -------------
	var starterProject string

	err = runtime.BindStyledParameter("simple", false, "starterProject", c.Param("starterProject"), &starterProject)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter starterProject: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ServeDevfileStarterProject(c, stack, starterProject)
}

// ServeDevfileWithVersion operation middleware
func (siw *ServerInterfaceWrapper) ServeDevfileWithVersion(c *gin.Context) {

	var err error

	// ------------- Path parameter "stack" -------------
	var stack string

	err = runtime.BindStyledParameter("simple", false, "stack", c.Param("stack"), &stack)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter stack: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "version" -------------
	var version string

	err = runtime.BindStyledParameter("simple", false, "version", c.Param("version"), &version)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter version: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ServeDevfileWithVersion(c, stack, version)
}

// ServeDevfileStarterProjectWithVersion operation middleware
func (siw *ServerInterfaceWrapper) ServeDevfileStarterProjectWithVersion(c *gin.Context) {

	var err error

	// ------------- Path parameter "stack" -------------
	var stack string

	err = runtime.BindStyledParameter("simple", false, "stack", c.Param("stack"), &stack)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter stack: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "version" -------------
	var version string

	err = runtime.BindStyledParameter("simple", false, "version", c.Param("version"), &version)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter version: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "starterProject" -------------
	var starterProject string

	err = runtime.BindStyledParameter("simple", false, "starterProject", c.Param("starterProject"), &starterProject)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter starterProject: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ServeDevfileStarterProjectWithVersion(c, stack, version, starterProject)
}

// ServeHealthCheck operation middleware
func (siw *ServerInterfaceWrapper) ServeHealthCheck(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ServeHealthCheck(c)
}

// ServeDevfileIndexV1 operation middleware
func (siw *ServerInterfaceWrapper) ServeDevfileIndexV1(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ServeDevfileIndexV1Params

	// ------------- Optional query parameter "arch" -------------

	err = runtime.BindQueryParameter("form", true, false, "arch", c.Request.URL.Query(), &params.Arch)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter arch: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "icon" -------------

	err = runtime.BindQueryParameter("form", true, false, "icon", c.Request.URL.Query(), &params.Icon)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter icon: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ServeDevfileIndexV1(c, params)
}

// ServeDevfileIndexV1WithType operation middleware
func (siw *ServerInterfaceWrapper) ServeDevfileIndexV1WithType(c *gin.Context) {

	var err error

	// ------------- Path parameter "indexType" -------------
	var indexType string

	err = runtime.BindStyledParameter("simple", false, "indexType", c.Param("indexType"), &indexType)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter indexType: %s", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params ServeDevfileIndexV1WithTypeParams

	// ------------- Optional query parameter "arch" -------------

	err = runtime.BindQueryParameter("form", true, false, "arch", c.Request.URL.Query(), &params.Arch)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter arch: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "icon" -------------

	err = runtime.BindQueryParameter("form", true, false, "icon", c.Request.URL.Query(), &params.Icon)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter icon: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ServeDevfileIndexV1WithType(c, indexType, params)
}

// ServeDevfileIndexV2 operation middleware
func (siw *ServerInterfaceWrapper) ServeDevfileIndexV2(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ServeDevfileIndexV2Params

	// ------------- Optional query parameter "arch" -------------

	err = runtime.BindQueryParameter("form", true, false, "arch", c.Request.URL.Query(), &params.Arch)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter arch: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "icon" -------------

	err = runtime.BindQueryParameter("form", true, false, "icon", c.Request.URL.Query(), &params.Icon)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter icon: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ServeDevfileIndexV2(c, params)
}

// ServeDevfileIndexV2WithType operation middleware
func (siw *ServerInterfaceWrapper) ServeDevfileIndexV2WithType(c *gin.Context) {

	var err error

	// ------------- Path parameter "indexType" -------------
	var indexType string

	err = runtime.BindStyledParameter("simple", false, "indexType", c.Param("indexType"), &indexType)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter indexType: %s", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params ServeDevfileIndexV2WithTypeParams

	// ------------- Optional query parameter "arch" -------------

	err = runtime.BindQueryParameter("form", true, false, "arch", c.Request.URL.Query(), &params.Arch)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter arch: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "icon" -------------

	err = runtime.BindQueryParameter("form", true, false, "icon", c.Request.URL.Query(), &params.Icon)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter icon: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ServeDevfileIndexV2WithType(c, indexType, params)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {

	errorHandler := options.ErrorHandler

	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/", wrapper.ServeRootEndpoint)

	router.GET(options.BaseURL+"/devfiles/:stack", wrapper.ServeDevfile)

	router.GET(options.BaseURL+"/devfiles/:stack/starter-projects/:starterProject", wrapper.ServeDevfileStarterProject)

	router.GET(options.BaseURL+"/devfiles/:stack/:version", wrapper.ServeDevfileWithVersion)

	router.GET(options.BaseURL+"/devfiles/:stack/:version/starter-projects/:starterProject", wrapper.ServeDevfileStarterProjectWithVersion)

	router.GET(options.BaseURL+"/health", wrapper.ServeHealthCheck)

	router.GET(options.BaseURL+"/index", wrapper.ServeDevfileIndexV1)

	router.GET(options.BaseURL+"/index/:indexType", wrapper.ServeDevfileIndexV1WithType)

	router.GET(options.BaseURL+"/v2index", wrapper.ServeDevfileIndexV2)

	router.GET(options.BaseURL+"/v2index/:indexType", wrapper.ServeDevfileIndexV2WithType)

	return router
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"//s8zlfcKAAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
