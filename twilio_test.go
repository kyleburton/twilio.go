package main

import (
	"strings"
	"testing"
	"github.com/kyleburton/twilio-go"
)


func TestMakeClient(t *testing.T) {
  sid := "banana"
  token := "kimchee"
  client = &twilio.Client {
    AccountSid: sid,
    AuthToken:  token,
  }

  if client.AccountSid != sid {
		t.Errorf("TestMakeClient(%s) :: %s != %s", client, sid, client.AccountSid)
  }

}
