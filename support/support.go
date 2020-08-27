package support

import (
	"encoding/json"
	"fmt"
	"os"
)

// Response ...
type Response struct {
	Msg     string `json:"msg"`
	Changed bool   `json:"changed"`
	Failed  bool   `json:"failed"`
}

// ExitJSON ...
func ExitJSON(responseBody Response) {
	returnResponse(responseBody)
}

// FailJSON ...
func FailJSON(responseBody Response) {
	responseBody.Failed = true
	returnResponse(responseBody)
}

// returnResponse ...
func returnResponse(responseBody Response) {
	var response []byte
	var err error
	response, err = json.Marshal(responseBody)
	if err != nil {
		response, _ = json.Marshal(Response{Msg: "Invalid response object"})
	}
	fmt.Println(string(response))
	if responseBody.Failed {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
