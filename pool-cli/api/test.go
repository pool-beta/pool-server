package api

import (
	"io/ioutil"
	"fmt"
	"bytes"
	"encoding/json"
)

func RunTestSetup(ctx Context) {
	fmt.Printf("Setting up a simple network for testing: %v\n", ctx.Url())

	route := ctx.Url() + "/test/setup"
	data, _ := json.Marshal(map[string]string{
		"password": "admin",
	})
	body := bytes.NewBuffer(data)

	req, err := NewRequest("GET", route, body)
	if err != nil {
		fmt.Println(err.Error())
	}

	res, err := ctx.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	// Read the response body
	resp, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(resp))
}

func RunTestReset(ctx Context) {
	fmt.Printf("Reseting this simple network for testing: %v\n", ctx.Url())

	route := ctx.Url() + "/test/reset"
	data, _ := json.Marshal(map[string]string{
		"password": "admin",
	})
	body := bytes.NewBuffer(data)

	req, err := NewRequest("GET", route, body)
	if err != nil {
		fmt.Println(err.Error())
	}

	res, err := ctx.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	// Read the response body
	resp, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(resp))
}



