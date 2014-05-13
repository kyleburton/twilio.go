package twilio

import (
	//"strings"
	"testing"
)


func TestMakeClient(t *testing.T) {
  sid := "banana"
  token := "kimchee"
  client := &Client {
    AccountSid: sid,
    AuthToken:  token,
  }

  if client.AccountSid != sid {
		t.Errorf("TestMakeClient(%s) :: %s != %s", client, sid, client.AccountSid)
  }

}
