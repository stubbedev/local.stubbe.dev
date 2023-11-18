package server

import (
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.tmpl")
	r.Static("/static", "./static/")

	tmpl_routes := getStaticRoutes("templates")
	for _, p := range tmpl_routes {
		r.GET(path.Clean(p), s.frontendHandler)
		r.POST(path.Clean(p), s.backendHandler)
	}

	return r
}

func (s *Server) frontendHandler(c *gin.Context) {
	tv := getTemplateVars(c.Request.URL.Path)
	pp := strings.TrimPrefix(strings.TrimSuffix(c.Request.URL.Path, "/")+"/index.tmpl", "/")
	c.HTML(http.StatusOK, pp, tv)
}

func (s *Server) backendHandler(c *gin.Context) {
	status, data := processRequest(s, c)
	c.JSON(status, data)
}

// func (s *Server) helloWorldHandler(c *gin.Context) {
// 	resp := make(map[string]string)
// 	resp["message"] = "Hello World"
//
// 	c.JSON(http.StatusOK, resp)
// }
