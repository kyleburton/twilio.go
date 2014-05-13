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
