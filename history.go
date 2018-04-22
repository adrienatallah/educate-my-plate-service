package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func historyHandler(code string) (*ProductInfo, error) {

	resp, err := http.Get("#" + code)

	if err != nil {
		fmt.Println("Error, ", err.Error(), ",with GET request with code", code)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response body", err.Error())
		return nil, err
	}

	p := new(ProductInfo)

	err = json.Unmarshal(body, &p)

	if err != nil {
		fmt.Println("Error unmarshalling json", err.Error())
		return nil, err
	}

	return p, nil
}
