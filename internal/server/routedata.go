package server

import "github.com/gin-gonic/gin"

func getTemplateVars(path string) gin.H {
	// ADD PATH SPECIFIC RESOURCES
	if path == "/" {
		return gin.H{
			"title":      "Home",
			"favicon":    "favicon.ico",
			"css_resets": "/static/css/resets.css",
		}
	}
	if path == "/todo" {
		return gin.H{
			"title":      "Todo",
			"favicon":    "/static/images/ico96.png",
			"css_resets": "/static/css/resets.css",
			"css":        "/static/pages/todo/css/style.css",
			"js":         "/static/pages/todo/js/main.js",
		}
	}
	return gin.H{
		"title":      "Not Found",
		"favicon":    "/static/images/ico96.png",
		"css_resets": "/static/css/resets.css",
	}
}
