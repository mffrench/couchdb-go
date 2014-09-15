package couchdb

import (
	"net/http"
	"testing"
)

type couchWelcome struct {
	Couchdb string      `json:"couchdb"`
	Uuid    string      `json:"uuid"`
	Version string      `json:"version"`
	Vendor  interface{} `json:"vendor"`
}

var serverUrl = "http://maui-test:5984"
var couchReply couchWelcome

func TestConnection(t *testing.T) {
	client := &http.Client{}
	c := connection{serverUrl, client, "", ""}
	resp, err := c.request("GET", "/", nil)
	if err != nil {
		t.Fail()
	} else if resp == nil {
		t.Fail()
	} else {
		jsonError := parseBody(resp, &couchReply)
		if jsonError != nil {
			t.Fail()
		} else {
			if resp.StatusCode != 200 ||
				couchReply.Couchdb != "Welcome" {
				t.Fail()
			}
			t.Logf("STATUS: %v\n", resp.StatusCode)
			t.Logf("couchdb: %v", couchReply.Couchdb)
			t.Logf("uuid: %v", couchReply.Uuid)
			t.Logf("version: %v", couchReply.Version)
			t.Logf("vendor: %v", couchReply.Vendor)
		}
	}
}

func TestBasicAuth(t *testing.T) {
	client := &http.Client{}
	c := connection{serverUrl, client, "adminuser", "password"}
	resp, err := c.request("GET", "/", nil)
	if err != nil {
		t.Logf("Error: %v", err)
		t.Fail()
	} else if resp == nil {
		t.Logf("Response was nil")
		t.Fail()
	} else {
	}
}

func TestBadAuth(t *testing.T) {
	client := &http.Client{}
	c := connection{serverUrl, client, "notauser", "what?"}
	resp, err := c.request("GET", "/", nil)
	if err == nil {
		t.Fail()
	} else if resp.StatusCode != 401 {
		t.Logf("Wrong Status: %v", resp.StatusCode)
		t.Fail()
	}
}
