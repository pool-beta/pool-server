package api

import (
	"io/ioutil"
	"fmt"
	"bytes"
	"encoding/json"
)

func RunPoolGet(ctx Context, username string, password string, pool_id string) {
	fmt.Printf("Getting Pool List: %v\n", ctx.Url())

	route := ctx.Url() + "/pool/poolid"
	data, _ := json.Marshal(map[string]string{
		"user_auth": map[string]string{
			"username": username, 
			"password": password,
		},
		"pool_id": pool_id,
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