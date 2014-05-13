package twilio

import (
	//"strings"
	"testing"
)

func MakeTestClient (sid, token string) *Client {
  return &Client {
    AccountSid: sid,
    AuthToken:  token,
  }
}

func TestMakeClient (t *testing.T) {
  sid := "banana"
  token := "kimchee"
  client := MakeTestClient(sid, token)

  if client.AccountSid != sid {
		t.Errorf("TestMakeClient(%s) :: %s != %s", client, sid, client.AccountSid)
  }

}

func TestMakeUrl (t *testing.T) {
  client := MakeTestClient("banana", "kimchee")

  expectedUrl := DefaultBaseUrl + "/Accounts/banana/foof?AccountSid=banana"
  if client.MakeUrl("foof", nil) != expectedUrl {
		t.Errorf("TestMakeUrl(%s) :: %s != %s", client, client.MakeUrl("foof", nil), expectedUrl )
  }
}

func TestUnmarshal (t *testing.T) {
  data := []byte(`{"AccountSid": "banana", "AuthToken": "kimchee", "BaseUrl": "https://ha.ha/9999-99-99"}`)
  client := &Client{}
  client.Unmarshal(data)

  if client.AccountSid != "banana" {
		t.Errorf("TestMakeClient(%s) :: %s != %s", client, "banana", client.AccountSid)
  }

  if client.AuthToken != "kimchee" {
		t.Errorf("TestMakeClient(%s) :: %s != %s", client, "kimchee", client.AuthToken)
  }

  if client.BaseUrl != "https://ha.ha/9999-99-99" {
		t.Errorf("TestMakeClient(%s) :: %s != %s", client, "https://ha.ha/9999-99-99", client.BaseUrl)
  }
}
