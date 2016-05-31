package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	host       string
	username   string
	password   string
	port       string
	insecure   bool
	httpClient http.Client
}

func NewClient(host, username, password string, port string, insecure bool) Client {
	httpClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
		},
	}

	return Client{host, username, password, port, insecure, httpClient}
}

func (c Client) APICall(httpMethod, apiPath string, reqStruct interface{}, respStruct interface{}) error {
	b, err := json.Marshal(reqStruct)
	if err != nil {
		return fmt.Errorf("error marshalling param %s\n", err)
	}
	uri := fmt.Sprintf("https://%s:%s/univmax%s", c.host, c.port, apiPath)
	req, err := http.NewRequest(httpMethod, uri, bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("error creating http request %s", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.username, c.password)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making http call %s\n", err)
	}
	//no content, no error :)
	if resp.StatusCode == 204 {
		return nil
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body %s\n", err)
	}

	err = json.Unmarshal(respBody, respStruct)
	if err != nil {
		return fmt.Errorf("error reading json response %s. Response Body %s\n", err, respBody)
	}

	return nil
}
