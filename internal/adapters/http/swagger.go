package http

import (
	"embed"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed swagger/openapi.yaml
var swaggerYAML []byte

const swaggerUIPage = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Mafia API Docs</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.onload = () => {
      window.ui = SwaggerUIBundle({
        url: '__SPEC_URL__',
        dom_id: '#swagger-ui',
      });
    };
  </script>
</body>
</html>`

// registerSwaggerRoutes mounts both the static OpenAPI document and a dynamic Swagger UI.
func registerSwaggerRoutes(r *gin.Engine) {
	r.GET("/swagger/openapi.yaml", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/x-yaml", swaggerYAML)
	})

	r.GET("/swagger/openapi-dynamic.yaml", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/x-yaml", renderDynamicSpec(c.Request))
	})

	r.GET("/swagger", swaggerUIHandler)
	r.GET("/swagger/", swaggerUIHandler)
}

func swaggerUIHandler(c *gin.Context) {
	html := strings.ReplaceAll(swaggerUIPage, "__SPEC_URL__", "/swagger/openapi-dynamic.yaml")
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}

func renderDynamicSpec(req *http.Request) []byte {
	scheme := "http"
	if req.TLS != nil {
		scheme = "https"
	}
	serverURL := fmt.Sprintf("%s://%s", scheme, req.Host)
	return []byte(strings.ReplaceAll(string(swaggerYAML), "http://localhost:8080", serverURL))
}
