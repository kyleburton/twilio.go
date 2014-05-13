package twilio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

var DefaultBaseUrl = "https://api.twilio.com/2010-04-01"

type Client struct {
	AccountSid string
	AuthToken  string
  BaseUrl    string
}

func (self *Client) MakeUrl (path string, params *url.Values) string {
	if params == nil {
		params = &url.Values{}
	}
	params.Add("AccountSid", self.AccountSid)
  baseUrl := self.BaseUrl
  if "" == baseUrl {
    baseUrl = DefaultBaseUrl
  }
	url := fmt.Sprintf("%s/Accounts/%s/%s?", baseUrl, self.AccountSid, path) + params.Encode()
  return url
}

func (self *Client) LoadJsonConfig (path string) error {
	data, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}
  self.Unmarshal(data)
  return nil
}

func (self *Client) Unmarshal (data []byte) {
	json.Unmarshal(data, self)
}
