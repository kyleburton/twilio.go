package twilio

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var DefaultClient = &Client{}
var TestBaseUrl = "http://localhost:8000"
var TestServer *httptest.Server

func init() {
	// a 'live' test against the Twilio API
	if os.Getenv("TWILIO_TEST_CONFIG") != "" {
		if testing.Verbose() {
			fmt.Fprintf(os.Stderr, "TWILIO_TEST_CONFIG=%s\n", os.Getenv("TWILIO_TEST_CONFIG"))
		}
		data, e := ioutil.ReadFile(os.Getenv("TWILIO_TEST_CONFIG"))
		if e != nil {
			fmt.Fprintf(os.Stderr, "File error: %v\n", e)
		}
		if testing.Verbose() {
			fmt.Fprintf(os.Stderr, "twilio_test.go: init, data=%s\n", string(data))
		}
		DefaultClient.Unmarshal(data)
		DefaultClient.Verbose = testing.Verbose()
		if testing.Verbose() {
			fmt.Fprintf(os.Stderr, "twilio_test.go: init, DefaultClient=%s\n", DefaultClient)
		}
	} else {
		if testing.Verbose() {
			fmt.Fprintf(os.Stderr, "init: Creating default\n")
		}
		TestServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "%s", `{"sid": "account.test-sid"}`)
		}))
		// how to close this at the end of all the tests?
		// defer TestServer.Close()
		TestBaseUrl = TestServer.URL
		DefaultClient = MakeTestClient("banana", "kimchee")
	}
}

func MakeTestClient(sid, token string) *Client {
	return &Client{
		AccountSid: sid,
		AuthToken:  token,
		Verbose:    testing.Verbose(),
		BaseUrl:    TestBaseUrl,
	}
}

func TestMakeClient(t *testing.T) {
	sid := "banana"
	token := "kimchee"
	client := MakeTestClient(sid, token)

	if client.AccountSid != sid {
		t.Errorf("TestMakeClient(%s) :: %s != %s", client, sid, client.AccountSid)
	}

}

func TestMakeUrl(t *testing.T) {
	client := MakeTestClient("banana", "kimchee")

	expectedUrl := client.BaseUrl + "/Accounts/banana/foof"
	if client.MakeUrl("/foof", nil) != expectedUrl {
		t.Errorf("TestMakeUrl(%s) :: %s != %s", client, client.MakeUrl("/foof", nil), expectedUrl)
	}
}

func TestUnmarshal(t *testing.T) {
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

// TODO: need to start a simple mock http server to run tests vs the api
func TestGetAccount(t *testing.T) {
	DefaultClient.Verbose = testing.Verbose()
	_, body, err := DefaultClient.GetAccount()
	if err != nil {
		t.Errorf("TestGetAccount(%s) :: error: %s", DefaultClient, err)
	}

	if testing.Verbose() {
		fmt.Fprintf(os.Stderr, "TestGetAccount: body=%s\n", body)
	}

	if testing.Verbose() {
    s, err := DefaultClient.AccountInfo.ToJsonStr()
		fmt.Fprintf(os.Stderr, "TestGetAccount: client.AccountInfo=%s err=%s\n", s, err)
	}

	// make assertions about res
	if body.Sid == "" {
		t.Errorf("body.Sid[%s] should not be empty: body=%s", body.Sid, body)
	}
}
