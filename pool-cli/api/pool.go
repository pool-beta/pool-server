package api

import (
	"io/ioutil"
	"fmt"
	"bytes"
	"strconv"
	"encoding/json"

	"github.com/pool-beta/pool-server/daemon/handlers/models"
)

func RunPoolGet(ctx Context, username string, password string, pool_id string) {
	fmt.Printf("Getting Pool List: %v\n", ctx.Url())

	var request models.RequestGetPool
	pid, _ := strconv.ParseUint(pool_id, 10, 64)

	request.UserAuth.UserName = username
	request.UserAuth.Password = password
	request.ID = pid

	route := ctx.Url() + "/pool/poolid"
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