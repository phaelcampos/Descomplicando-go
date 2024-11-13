package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GithubWebHookRequestBody1 struct {
	Action     string
	Number     int64
	Repository string
	Sender     string
}

type GithubWebHookHandler1 struct{}

func (h *GithubWebHookHandler1) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		githubWebHookRequest := new(GithubWebHookRequestBody1)
		err = json.Unmarshal(body, githubWebHookRequest)
		if err != nil {
			panic(err)
		}
		fmt.Println("Action", githubWebHookRequest.Action)
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("ok"))
		return
	}
	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte("Not OK"))
}

func main1() {
	handler := new(GithubWebHookHandler1)
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}
}

// curl -X POST -v  localhost:8080 -d {"acation": "opened","number": 5000,"repository": "teste","sender": "sender"}
