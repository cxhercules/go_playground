package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Error struct {
	HTTPCode int    `json: "-"`
	Code     int    `json: "code,omitempty"`
	Message  string `json: "message"`
}

func main() {
	res, _ := http.Get("http://goinpractice.com")

	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", b)
}
