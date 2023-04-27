package handlers

import (
	"assignment4/internal/contacts"
	"assignment4/internal/db"
	"bytes"
	"encoding/json"
	"strings"

	"net/http"
	"net/http/httptest"
	"testing"
)

var contact1 = contacts.Contact{
	Id:    "test1",
	Name:  "Test1",
	Email: "test@addcontact",
	Phno:  "123456",
}
var contact2 = contacts.Contact{
	Id:    "test2",
	Name:  "Test2",
	Email: "test2@addcontact",
	Phno:  "1234567",
}
var invalidContact = contacts.Contact{
	Id:    "",
	Name:  "Test2",
	Email: "test2@addcontact",
	Phno:  "1234567",
}
var updatedcontact1 = contacts.Contact{
	Id:    "test1",
	Name:  "UpdatedTest1",
	Email: "Updatedtest1@addcontact",
	Phno:  "0123456",
}
var partialupdatedcontact2 = contacts.Contact{
	Id:    "test2",
	Name:  "",
	Email: "",
	Phno:  "",
}

func InsertContacts() {
	db.AddContact(contact1)

	db.AddContact(contact2)
}

func TestMain(m *testing.M) {
	m.Run()
}

func TestDeletecontact(t *testing.T) {

	InsertContacts()
	defer db.ClearDB()
	tests := []struct {
		name     string
		url      string
		respCode int
		respBody string
	}{
		{
			"Delete contact",
			"/contact?id=test1",
			200,
			"contact deleted successfully",
		},
		{
			"Delete contact without id",
			"/contact",
			400,
			"id not passed",
		},
		{
			"Delete not existing contact",
			"/contact?id=Invalidcontact",
			400,
			"contact with id not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			DeleteContact(rr, req)

			if status := rr.Code; status != tt.respCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.respCode)
			}

			resBody := strings.TrimSuffix(rr.Body.String(), "\n")
			if strings.Compare(resBody, tt.respBody) != 0 {
				t.Errorf("handler returned wrong resp body: got --%s-- want --%s--",
					resBody, tt.respBody)
			}
		})
	}
}

func TestGetcontacts(t *testing.T) {
	contactList := []contacts.Contact{contact1, contact2}

	allcontactjson, err := json.Marshal(contactList)
	if err != nil {
		t.Error("cannot marshal contact list")
	}

	contact1json, err := json.Marshal([]contacts.Contact{contact1})
	if err != nil {
		t.Error("cannot marshal contact 1")
	}

	InsertContacts()
	defer db.ClearDB()
	tests := []struct {
		name     string
		url      string
		respCode int
		respBody string
	}{
		{
			"Get All Contact",
			"/contact",
			200,
			string(allcontactjson),
		},
		{
			"Get Contact1",
			"/contact?id=test1",
			200,
			string(contact1json),
		},
		{
			"Get Invalid Contact id",
			"/contact?id=Invalidcontact",
			400,
			"Requested contact not present",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			GetContacts(rr, req)

			if status := rr.Code; status != tt.respCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.respCode)
			}

			resBody := strings.TrimSuffix(rr.Body.String(), "\n")
			if strings.Compare(resBody, tt.respBody) != 0 {
				t.Errorf("handler returned wrong resp body: got --%s-- want --%s--",
					resBody, tt.respBody)
			}
		})
	}
}

func TestSetcontact(t *testing.T) {

	defer db.ClearDB()
	contact1json, err := json.Marshal(contact1)
	if err != nil {
		t.Error("cannot marshal contact 1")
	}

	invalidjson, err := json.Marshal(invalidContact)
	if err != nil {
		t.Error("cannot marshal contact 1")
	}

	tests := []struct {
		name     string
		reqBody  []byte
		respCode int
		respBody string
	}{
		{
			"Add Contact1",
			contact1json,
			200,
			string(contact1json),
		},
		{
			"Add Invalid Contact id",
			invalidjson,
			400,
			"Invalid ID value",
		},
		{
			"Add existing Contact",
			contact1json,
			400,
			"Contact already present",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/contact", bytes.NewReader(tt.reqBody))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			SetContact(rr, req)

			if status := rr.Code; status != tt.respCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.respCode)
			}

			resBody := rr.Body.String()
			resBody = strings.TrimSuffix(resBody, "\n")
			//			t.Log(resBody)
			if strings.Compare(resBody, tt.respBody) != 0 {
				t.Errorf("handler returned wrong resp body: got --%s-- want --%s--",
					resBody, tt.respBody)
			}
		})
	}
}

func TestUpdatecontact(t *testing.T) {
	contact1json, err := json.Marshal(updatedcontact1)
	if err != nil {
		t.Error("cannot marshal contact 1")
	}

	contact2json, err := json.Marshal(contact2)
	if err != nil {
		t.Error("cannot marshal contact 1")
	}

	partialupdatedcontact2json, err := json.Marshal(partialupdatedcontact2)
	if err != nil {
		t.Error("cannot marshal contact 1")
	}

	InsertContacts()
	defer db.ClearDB()
	tests := []struct {
		name     string
		url      string
		reqBody  []byte
		respCode int
		respBody string
	}{
		{
			"update contact 1",
			"/contact?id=test1",
			contact1json,
			200,
			string(contact1json),
		},
		{
			"update partial contacts",
			"/contact?id=test2",
			partialupdatedcontact2json,
			200,
			string(contact2json),
		},
		{
			"update without id",
			"/contact",
			contact1json,
			400,
			"id not passed",
		},
		{
			"update with Invalid Contact id",
			"/contact?id=Invalidcontact",
			contact1json,
			400,
			"contact ID not present",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("PUT", tt.url, bytes.NewReader(tt.reqBody))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			UpdateContact(rr, req)

			if status := rr.Code; status != tt.respCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.respCode)
			}

			resBody := strings.TrimSuffix(rr.Body.String(), "\n")
			if strings.Compare(resBody, tt.respBody) != 0 {
				t.Errorf("handler returned wrong resp body: got --%s-- want --%s--",
					resBody, tt.respBody)
			}
		})
	}
}
