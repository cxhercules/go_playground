package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	_, err := http.Get("http://example.com/file.zip")
	if err != nil && hasTimeOut(err) {
		fmt.Println("A timout error occured")
		return
	}
}

func hasTimeOut(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	case net.Error:
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}
	}
	errTxt := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}
	return false
}
