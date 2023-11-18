package server

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func processRequest(s *Server, c *gin.Context) (int, map[string]interface{}) {
	p := strings.TrimSuffix(strings.TrimPrefix(c.Request.URL.Path, "/"), "/")
	a := action(c.Request.Header.Get("X-Action"))
	status, data := controller(a, p, c, s)
	return status, data
}

func action(key string) string {
	if key == "POST" || key == "post" {
		return "SET"
	} else if key == "DELETE" || key == "delete" {
		return "DEL"
	}
	return "GET"
}

func controller(action string, route_path string, c *gin.Context, s *Server) (int, map[string]interface{}) {
	if route_path == "todo" {
		c.Request.Header.Add("X-User-Id", fingerPrint(c))
		return s.db.Todo(action, c)
	} else if route_path == "health" {
		return s.db.Health(action, c)
	}
	return 500, map[string]interface{}{}
}
