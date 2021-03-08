package api

import (
	"io/ioutil"
	"fmt"
	"bytes"
	"encoding/json"

	"github.com/pool-beta/pool-server/daemon/handlers/models"
)

func RunTankCreate(ctx Context, username string, password string, poolname string) {
	fmt.Printf("Creating tank: %v\n", ctx.Url())

	var request models.RequestCreatePool

	request.UserAuth.UserName = username
	request.UserAuth.Password = password

	route := ctx.Url() + "/tanks/create"
	data, _ := json.Marshal(request)
	
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