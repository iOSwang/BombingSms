package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type SmsJobConfig struct {
	Url        string            `json:"Url"`
	Method     string            `json:"Method"`
	ParamsStr  string            `json:"ParamsBody"`
	ParamsType string            `json:"ParamsType"`
	ParamsForm url.Values        `json:"-"`
	ParamsJson map[string]string `json:"-"`
	Ignore     bool              `json:"Ignore"`
}

const Phone = "13051102520"

func main() {
	data, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var jobList []SmsJobConfig
	err = json.Unmarshal(data, &jobList)
	if err != nil {
		panic(err)
	}
	for index, job := range jobList {
		if job.Ignore {
			continue
		}
		if job.Method == "GET" {
			var resp *http.Response
			url := fmt.Sprintf("%s&phone=%s", job.Url, Phone)
			resp, err = http.DefaultClient.Get(url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			var body []byte
			body, err = io.ReadAll(resp.Body)
			fmt.Println(index, "-----json------", string(body))
		} else if job.Method == "POST" {
			var resp *http.Response
			url := fmt.Sprintf("%s&phone=%s", job.Url, Phone)
			b, _ := json.Marshal(job.ParamsJson)
			resp, err = http.DefaultClient.Post(url, "application/json", bytes.NewBuffer(b))
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			var body []byte
			body, err = io.ReadAll(resp.Body)
			fmt.Println(index, "-----json------", string(body))
		}
	}

}
