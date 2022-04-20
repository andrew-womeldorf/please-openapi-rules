package validate

import (
	_ "embed"
	"fmt"
	"net/http"
	"strings"
)

//go:embed petstore.yaml
var openapiSpec string

type SpecApiController struct{}

func NewSpecApiController() Router {
	return &SpecApiController{}
}

func (c *SpecApiController) Routes() Routes {
	return Routes{
		{
			"GetSpec",
			strings.ToUpper("Get"),
			"/openapi.yaml",
			c.GetSpec,
		},
	}
}

func (c *SpecApiController) GetSpec(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, openapiSpec)
}
