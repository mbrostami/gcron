package pages

import "github.com/gin-gonic/gin"

// Page interface for web pages
type Page interface {
	Handler(method string, c *gin.Context) Response

	GetRoute() string
	GetMethods() []string
}

// Response simple interface for page responses
type Response interface {
}
