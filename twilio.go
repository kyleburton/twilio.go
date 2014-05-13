package twilio

import (
	"fmt"
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
