package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Note:
// This gateway can be simply replaced with a load balancer or proxy server
// But I wanted to do it manually
func main() {
	r := gin.Default()

	//targetURL, _ := fmt.Printf("http://%s:%s", os.Getenv("MYSQL_TCP_HOST"), os.Getenv("AUTH_MYSQL_TCP_PORT"))

	// TODO - Replace hardcoded with above commented.
	// Define routes and route handlers
	r.GET("/api/auth/*path", ProxyToService("http://host.docker.internal:3001"))
	r.GET("/api/offer/*path", ProxyToService("http://host.docker.internal:3002"))
	r.GET("/api/order/*path", ProxyToService("http://host.docker.internal:3003"))
	r.GET("/api/product/*path", ProxyToService("http://host.docker.internal:3004"))

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func ProxyToService(targetURL string) gin.HandlerFunc {
	target, err := url.Parse(targetURL)
	if err != nil {
		// Handle the error
		return func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid target URL"})
		}
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(target)

	return func(c *gin.Context) {
		// Extract the path and parameters from the URL
		queryParameters := c.Request.URL.RawQuery

		// Modify the request URL
		c.Request.URL.RawQuery = queryParameters

		// Update the request to pass through the reverse proxy
		reverseProxy.ServeHTTP(c.Writer, c.Request)
	}
}
