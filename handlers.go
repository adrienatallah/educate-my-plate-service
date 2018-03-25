package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func searchHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	code := r.URL.Query().Get("code")

	fmt.Println("searching for code", code)

	info, err := search(code)

	if err != nil {
		fmt.Println("Error searching for code", code)
		return
	}

	out, err := json.Marshal(info)

	if err != nil {
		fmt.Println("error marshalling json: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
