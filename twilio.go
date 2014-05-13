package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var DefaultBaseUrl = "https://api.twilio.com/2010-04-01"

type Client struct {
	AccountSid  string
	AuthToken   string
	BaseUrl     string
	Verbose     bool
	AccountInfo *AccountResponse
}

type AccountResponse struct {
	Sid              string
	Friendly_name    string
	Type             string
	Status           string
	Date_created     string
	Date_updated     string
	Auth_token       string
	Uri              string
	Subresource_uris map[string]string
}

func (self *AccountResponse) ToJsonStr () (string, error) {
  data, err := json.Marshal(self)
  return string(data), err
}

func (self *Client) MakeUrl(path string, params *url.Values) string {
	var url string
	baseUrl := self.BaseUrl
	if "" == baseUrl {
		baseUrl = DefaultBaseUrl
	}

	if params == nil {
		url = fmt.Sprintf("%s/Accounts/%s%s", baseUrl, self.AccountSid, path)
	} else {
		url = fmt.Sprintf("%s/Accounts/%s%s?", baseUrl, self.AccountSid, path) + params.Encode()
	}
	return url
}

func (self *Client) LoadJsonConfig(path string) error {
	data, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}
	self.Unmarshal(data)
	return nil
}

func (self *Client) Unmarshal(data []byte) {
	var conf map[string]string
	err := json.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}
	if self.Verbose {
		fmt.Fprintf(os.Stderr, "Unmarshal: data=%s conf=%s\n", data, conf)
	}
	var val string
	var hasVal bool
	val, hasVal = conf["AccountSid"]
	if hasVal {
		self.AccountSid = val
	}

	val, hasVal = conf["AuthToken"]
	if hasVal {
		self.AuthToken = val
	}

	val, hasVal = conf["BaseUrl"]
	if hasVal {
		self.BaseUrl = val
	}

	val, hasVal = conf["Verbose"]
	if hasVal {
		if val == "true" {
			self.Verbose = true
		}
	}
}

func (self *Client) GetAccount() (*http.Response, *AccountResponse, error) {
	path := self.MakeUrl(".json", nil)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	if self.Verbose {
		fmt.Fprintf(os.Stderr, "GetAccount() req=%s\n", req)
	}
	req.SetBasicAuth(self.AccountSid, self.AuthToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != 200 {
		return resp, nil, errors.New(fmt.Sprintf("Status{%s} != OK{200}", resp.Status))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading body: %s\n", err)
		return resp, nil, err
	}

	if self.Verbose {
		fmt.Fprintf(os.Stderr, "GetAccount: body=`%s`\n", body)
	}

	result := &AccountResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return resp, nil, err
	}

	self.AccountInfo = result

	return resp, result, nil
}

