package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type T struct {
	Authorization bool
}

func main() {
	authorization := T{}
	getJson("https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31", &authorization)
	println(authorization.Authorization)
}
