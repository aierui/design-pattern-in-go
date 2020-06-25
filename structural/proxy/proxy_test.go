package proxy

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {

	nginxServer := newNginxServer()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"
	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	/*
	   Url: /app/status
	   HttpCode: 200
	   Body: Ok

	   Url: /app/status
	   HttpCode: 200
	   Body: Ok

	   Url: /app/status
	   HttpCode: 403
	   Body: Not Allowed

	   Url: /app/status
	   HttpCode: 201
	   Body: User Created

	   Url: /app/status
	   HttpCode: 404
	   Body: Not Ok
	*/
}
