package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const gateway = "http://gateway.openfaas:8080/function/"

type FlowModel struct {
	FuncName string `json:"funcName"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// Get data from environment variable
	flowData := os.Getenv("data")
	fmt.Println(flowData)

	// Get byte slice from string.
	bytes := []byte(flowData)

	// Unmarshal string into structs.
	var models []FlowModel
	marshalErr := json.Unmarshal(bytes, &models)
	if marshalErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(marshalErr.Error()))
	}

	// Run flow request
	var input []byte
	x := "Hello"
	fmt.Print(x)
	var reqFlow *http.Request
	var resFlow *http.Response
	for index, flow := range models {
		if index == 0 {
			reqFlow, _ = http.NewRequest(http.MethodPost, gateway+flow.FuncName, r.Body)

		} else {
			reqFlow, _ = http.NewRequest(http.MethodPost, gateway+flow.FuncName, resFlow.Body)

		}

		reqFlow.Header.Set("Content-Type", "application/octet-stream")
		resFlow, _ = http.DefaultClient.Do(reqFlow)
		if index == (len(models) - 1) {
			if resFlow.Body != nil {

				body, _ := ioutil.ReadAll(resFlow.Body)

				input = body
			}
		}
	}
	defer resFlow.Body.Close()

	w.WriteHeader(http.StatusOK)
	w.Write(input)

}
