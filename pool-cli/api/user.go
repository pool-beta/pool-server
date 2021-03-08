package api

import (
	"io/ioutil"
	"fmt"
	"bytes"
	"encoding/json"
)

func RunUserList(ctx Context) {
	fmt.Printf("Getting User List: %v\n", ctx.Url())

	route := ctx.Url() + "/users"
	data, _ := json.Marshal(map[string]string{})
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

func RunUserCreate(ctx Context, username string, password string) {
	fmt.Printf("Creating new user: %v\n", ctx.Url())

	route := ctx.Url() + "/users/create"
	data, _ := json.Marshal(map[string]string{
		"username": username,
		"password": password,
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